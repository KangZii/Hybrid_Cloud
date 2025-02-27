package controller

import (
	v1alpha1hcppolicy "Hybrid_Cloud/pkg/client/hcppolicy/v1alpha1/clientset/versioned"
	"fmt"
	"time"

	hcppolicyscheme "Hybrid_Cloud/pkg/client/hcppolicy/v1alpha1/clientset/versioned/scheme"
	informers "Hybrid_Cloud/pkg/client/hcppolicy/v1alpha1/informers/externalversions/hcppolicy/v1alpha1"
	lister "Hybrid_Cloud/pkg/client/hcppolicy/v1alpha1/listers/hcppolicy/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

const controllerAgentName = "hcppolicy-controller"

const (
	// SuccessSynced is used as part of the Event 'reason' when a Foo is synced
	SuccessSynced = "Synced"
	// ErrResourceExists is used as part of the Event 'reason' when a Foo fails
	// to sync due to a Deployment of the same name already existing.
	ErrResourceExists = "ErrResourceExists"

	// MessageResourceExists is the message used for Events when a resource
	// fails to sync due to a Deployment already existing
	MessageResourceExists = "Resource %q already exists and is not managed by Foo"
	// MessageResourceSynced is the message used for an Event fired when a Foo
	// is synced successfully
	MessageResourceSynced = "Foo synced successfully"
)

type Controller struct {
	// Kubernetes Core Resource 접근시 사용하는 ClientSet
	kubeclientset kubernetes.Interface
	// Custom Resource 접근시  사용하는 ClientSet
	hcppolicyclientset v1alpha1hcppolicy.Interface
	// deploymentsLister  appslisters.DeploymentLister
	// deploymentsSynced  cache.InformerSynced
	hcppolicyLister lister.HCPPolicyLister
	hcppolicySynced cache.InformerSynced
	workqueue       workqueue.RateLimitingInterface
	recorder        record.EventRecorder
}

func NewController(
	kubeclientset kubernetes.Interface,
	hcppolicyclientset v1alpha1hcppolicy.Interface,
	// deploymentInformer appsinformers.DeploymentInformer,
	hcppolicyInformer informers.HCPPolicyInformer) *Controller {
	utilruntime.Must(hcppolicyscheme.AddToScheme(scheme.Scheme))

	klog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(0)
	// kubernetes client가 클러스터 API를 이용해 내부에 이벤트 전송
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("hcp")})
	// 이벤트 생성
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &Controller{
		kubeclientset:      kubeclientset,
		hcppolicyclientset: hcppolicyclientset,
		// deploymentsLister:  deploymentInformer.Lister(),
		// deploymentsSynced:  deploymentInformer.Informer().HasSynced,
		hcppolicyLister: hcppolicyInformer.Lister(),
		hcppolicySynced: hcppolicyInformer.Informer().HasSynced,
		workqueue:       workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "hcppolicy"),
		recorder:        recorder,
	}

	klog.Info("Setting up event handlers")
	// Set up an event handler for when resources change
	hcppolicyInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueHCPPolicy,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueueHCPPolicy(new)
		},
	})

	// Set up an event handler for when Deployment resources change. This
	// handler will lookup the owner of the given Deployment, and if it is
	// owned by a Foo resource then the handler will enqueue that Foo resource for
	// processing. This way, we don't need to implement custom logic for
	// handling Deployment resources. More info on this pattern:
	// https://github.com/kubernetes/community/blob/8cafef897a22026d42f5e5bb3f104febe7e29830/contributors/devel/controllers.md
	// deploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
	// 	AddFunc: controller.handleObject,
	// 	UpdateFunc: func(old, new interface{}) {
	// 		newDepl := new.(*appsv1.Deployment)
	// 		oldDepl := old.(*appsv1.Deployment)
	// 		if newDepl.ResourceVersion == oldDepl.ResourceVersion {
	// 			// Periodic resync will send update events for all known Deployments.
	// 			// Two different versions of the same Deployment will always have different RVs.
	// 			return
	// 		}
	// 		controller.handleObject(new)
	// 	},
	// 	DeleteFunc: controller.handleObject,
	// })
	return controller
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *Controller) Run(workers int, stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()
	defer c.workqueue.ShutDown()

	// Start the informer factories to begin populating the informer caches
	klog.Info("Starting HCPPolicy controller")

	// Wait for the caches to be synced before starting workers
	klog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.hcppolicySynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	klog.Info("Starting workers")
	// Launch two workers to process Foo resources
	for i := 0; i < workers; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	klog.Info("Started workers")
	<-stopCh
	klog.Info("Shutting down workers")

	return nil
}

// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the
// workqueue.
func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
	}
}

// processNextWorkItem will read a single work item off the workqueue and
// attempt to process it, by calling the syncHandler.
func (c *Controller) processNextWorkItem() bool {
	obj, shutdown := c.workqueue.Get()

	if shutdown {
		return false
	}
	// We wrap this block in a func so we can defer c.workqueue.Done.
	err := func(obj interface{}) error {
		// We call Done here so the workqueue knows we have finished
		// processing this item. We also must remember to call Forget if we
		// do not want this work item being re-queued. For example, we do
		// not call Forget if a transient error occurs, instead the item is
		// put back on the workqueue and attempted again after a back-off
		// period.
		defer c.workqueue.Done(obj)
		var key string
		var ok bool
		// We expect strings to come off the workqueue. These are of the
		// form namespace/name. We do this as the delayed nature of the
		// workqueue means the items in the informer cache may actually be
		// more up to date that when the item was initially put onto the
		// workqueue.
		if key, ok = obj.(string); !ok {
			// As the item in the workqueue is actually invalid, we call
			// Forget here else we'd go into a loop of attempting to
			// process a work item that is invalid.
			c.workqueue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		// Run the syncHandler, passing it the namespace/name string of the
		// Foo resource to be synced.
		if err := c.syncHandler(key); err != nil {
			// Put the item back on the workqueue to handle any transient errors.
			c.workqueue.AddRateLimited(key)
			return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
		}
		// Finally, if no error occurs we Forget this item so it does not
		// get queued again until another change happens.
		c.workqueue.Forget(obj)
		klog.Infof("Successfully synced '%s'", key)
		return nil
	}(obj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}

func (c *Controller) syncHandler(key string) error {
	// Convert the namespace/name string into a distinct namespace and name
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	hcppolicy, err := c.hcppolicyLister.HCPPolicies(namespace).Get(name)
	if err != nil {
		// The Foo resource may no longer exist, in which case we stop
		// processing.
		if errors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("hcppolicy '%s' in work queue no longer exists", key))
			return nil
		}

		return err
	}

	if hcppolicy.Spec.PolicyStatus == "Disabled" {
		klog.Info("Policy Disabled")
	} else if hcppolicy.Spec.PolicyStatus == "Enabled" {
		klog.Info("Policy Enabled")
		if hcppolicy.Spec.RangeOfApplication == "FromNow" {
			klog.Info("Policy Enabled - FromNow")
		} else if hcppolicy.Spec.RangeOfApplication == "All" {
			object := hcppolicy.Spec.Template.Spec.TargetController.Kind
			if object == "" {
				klog.Info("No TargetController")
			}
		}
	}
	return nil

}

// enqueueFoo takes a HCPPolicy resource and converts it into a namespace/name
// string which is then put onto the work queue. This method should *not* be
// passed resources of any type other than Foo.
func (c *Controller) enqueueHCPPolicy(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.workqueue.Add(key)
}
