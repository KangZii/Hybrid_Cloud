/*
Copyright 2018 The Multicluster-Controller Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import ( //"flag"
	//"strings"
	//"context"
	//"admiralty.io/multicluster-controller/pkg/controller"
	//"admiralty.io/multicluster-controller/pkg/reconcile"
	//"admiralty.io/multicluster-service-account/pkg/config"
	//"k8s.io/api/core/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	//"k8s.io/client-go/rest"
	//genericclient "sigs.k8s.io/kubefed/pkg/client/generic"
	//fedv1b1 "sigs.k8s.io/kubefed/pkg/apis/core/v1beta1"
	//"sigs.k8s.io/kubefed/pkg/controller/util"
)

func main() {
	//var f = flag.String("contexts", "", "a comma-separated list of contexts to watch, e.g., cluster1,cluster2")
	//flag.Parse()
	//ctxs := strings.Split(*f, ",")
	/*
	   	ghosts := []cluster.Cluster{}
	   	ghostNamespaces := []string{}

	   	host_ctx := "openmcp"
	   	ghost_ctxs := []string{"cluster1, cluster2"}

	   	host_cfg, _, err := config.NamedConfigAndNamespace(host_ctx)

	   	if err != nil {
	                   log.Fatal(err)
	           }



	   	for _, ghost_ctx := range ghost_ctxs {
	   		ghost_cfg, _, err := config.NamedConfigAndNamespace(ghost_ctx)
	   		if err != nil {
	   			log.Fatal(err)
	   		}
	   		ghost := cluster.New(ghost_ctx, ghost_cfg, cluster.Options{})
	   		ghostNamespace := "default"

	   		ghosts = append(ghosts, *ghost)
	   		ghostNamespaces = append(ghostNamespaces, ghostNamespace)
	   	}
	   	co, _ := templateresource.NewController(live, ghosts, ghostNamespaces)

	   	m := manager.New()
	   	m.AddController(co)

	   	if err := m.Start(signals.SetupSignalHandler()); err != nil {
	   		log.Fatal(err)
	   	}
	*/

	// cm := clusterManager.NewClusterManager()

	// host_ctx := "master"
	// namespace := "hybrid"
	// host_cfg, _ := cobrautil.BuildConfigFromFlags("master", "/root/.kube/config")
	// // host_cfg := cm.Host_config
	// live := cluster.New(host_ctx, host_cfg, cluster.Options{CacheOptions: cluster.CacheOptions{Namespace: namespace}})
	// ghosts := []*cluster.Cluster{}

	// co, _ := templateresource.NewController(live, ghosts, namespace, cm)

	// m := manager.New()
	// m.AddController(co)
	// if err := m.Start(signals.SetupSignalHandler()); err != nil {
	// 	log.Fatal(err)
	// }

}
