package main

import (
	controller "Hybrid_Cloud/hcp-cluster-manager/src/controller"
	informers "Hybrid_Cloud/pkg/client/hcpcluster/v1alpha1/informers/externalversions"
	"Hybrid_Cloud/util/clusterManager"
	"flag"
	"time"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/klog/v2"
	"k8s.io/sample-controller/pkg/signals"
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	cm, err := clusterManager.NewClusterManager()
	if err != nil {
		klog.Errorln(err)
	}

	stopCh := signals.SetupSignalHandler()
	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(cm.Host_kubeClient, time.Second*30)
	hcpclusterInformerFactory := informers.NewSharedInformerFactory(cm.HCPCluster_Client, time.Second*30)
	//
	controller := controller.NewController(cm.Host_kubeClient, cm.HCPCluster_Client, hcpclusterInformerFactory.Hcp().V1alpha1().HCPClusters())
	kubeInformerFactory.Start(stopCh)
	hcpclusterInformerFactory.Start(stopCh)
	if err := controller.Run(2, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}

}