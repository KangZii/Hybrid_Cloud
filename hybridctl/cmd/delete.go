package cmd

import (
	"Hybrid_Cloud/hybridctl/util"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var file_name string

// DeleteCmd represents the Delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `hybridctl delete deployment <name> -n <namespace> 
	hybridctl delete -f <filename> `,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		file_name = util.Option_file
		if file_name == "" {
			if len(args) < 2 {
				fmt.Println(cmd.Help())
			} else {
				LINK := "/resources"
				namespace, _ := cmd.Flags().GetString("namespace")
				if namespace == "" {
					namespace = "default"
				}
				LINK += "/namespaces/" + namespace

				util.Option_Resource = args[0]
				util.Option_Name = args[1]
				LINK += "/deployments/" + util.Option_Name

				_, err := util.GetResponseBody("DELETE", LINK, nil)
				if err != nil {
					fmt.Println(err)
				}
			}
		} else {
			DeleteResource()
		}
	},
}

func DeleteResource() {

	yaml, err := ReadFile()
	if err != nil {
		println(err)
		return
	}

	obj, gvk, err := GetObject(yaml)
	if err != nil {
		println(err)
		return
	}

	RequestDeleteResource(obj, gvk)
}

func RequestDeleteResource(obj runtime.Object, gvk *schema.GroupVersionKind) ([]byte, error) {

	LINK := "/resources"
	// check context flag
	//	flag_context := util.Option_context
	// var target_cluster string
	// var resource Resource

	// config, _ := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
	// 	&clientcmd.ClientConfigLoadingRules{ExplicitPath: "/root/.kube/config"},
	// 	&clientcmd.ConfigOverrides{
	// 		CurrentContext: "",
	// 	}).RawConfig()

	// if flag_context == "" {
	// 	target_cluster = ""
	// } else {
	// 	target_cluster = flag_context
	// }

	// match obj kind
	switch gvk.Kind {
	case "Deployment":
		real_resource := obj.(*appsv1.Deployment)
		namespace := real_resource.Namespace
		if namespace == "" {
			namespace = "default"
		}
		LINK += "/namespaces/" + namespace + "/deployments/" + real_resource.Name
	}

	fmt.Println(LINK)
	bytes, err := util.GetResponseBody("DELETE", LINK, nil)
	if err != nil {
		fmt.Println(err)
	}

	return bytes, err
}

// delete common CLI

var deleteNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("delete called")
		platform_name, err := cmd.Flags().GetString("platform")
		if err != nil {
			panic(err)
		}
		cluster_name, err := cmd.Flags().GetString("cluster-name")
		if err != nil {
			panic(err)
		}
		NodeName, err := cmd.Flags().GetString("nodepool-name")
		if err != nil {
			panic(err)
		}
		var cli = Cli{}
		cli.ClusterName = cluster_name
		cli.NodeName = NodeName
		cli.PlatformName = platform_name
		if platform_name == "aks" {
			deleteNodepool_aks(cli)
			fmt.Println("call delete_aks_nodepool func")
			fmt.Println(cli)
		} else if platform_name == "eks" {
			fmt.Println("call delete_eks_nodepool func")
			fmt.Println(cli)
			deleteNodepool_eks(cli)
		} else if platform_name == "gke" {
			fmt.Println("call delete_gke_nodepool func")
			fmt.Println(cli)
			deleteNodepool_gke(cli)
		} else {
			fmt.Println("Error: please enter the correct --platform arguments")
		}
	},
}
var deleteClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		platform_name, err := cmd.Flags().GetString("platform")
		if err != nil {
			panic(err)
		}
		cluster_name, err := cmd.Flags().GetString("cluster-name")
		if err != nil {
			panic(err)
		}

		var cli = Cli{}
		cli.ClusterName = cluster_name

		cli.PlatformName = platform_name

		if platform_name == "aks" {
			fmt.Println("call delete_aks func")
			delete_aks(cli)
		} else if platform_name == "eks" {
			fmt.Println("call delete_eks func")
			delete_eks(cli)
		} else if platform_name == "gke" {
			fmt.Println("call delete_gke func")
			delete_gke(cli)
		} else {
			fmt.Println("Error: please enter the correct --platform arguments")
		}

	},
}

func deleteNodepool_gke(info Cli) {

	cluster := "cluster"

	cmd_rm_nodepool := exec.Command("rm", info.ClusterName+"_"+info.NodeName+".tf.json")
	cmd_rm_nodepool.Dir = "../terraform/gke/" + cluster

	output, err := cmd_rm_nodepool.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	//cmd := exec.Command("terraform", "destroy", "-auto-approve")
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = "../terraform/gke/" + cluster

	output, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

}

func deleteNodepool_eks(info Cli) {

	cmd_rm_nodepool := exec.Command("rm", info.ClusterName+"nodepool"+".tf.json")
	cmd_rm_nodepool.Dir = "../terraform/eks"

	output, err := cmd_rm_nodepool.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	//cmd := exec.Command("terraform", "destroy", "-auto-approve")
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = "../terraform/eks"

	output, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

}

func deleteNodepool_aks(info Cli) {

	cmd_rm_nodepool := exec.Command("rm", info.ClusterName+"_"+info.NodeName+".tf.json")
	cmd_rm_nodepool.Dir = "../terraform/aks"

	output, err := cmd_rm_nodepool.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	//cmd := exec.Command("terraform", "destroy", "-auto-approve")
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = "../terraform/aks"

	output, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

}

func delete_aks(info Cli) {
	// num := 1
	// data := make([]Cluster_info, 1)
	// cluster := "cluster"

	// fmt.Println("!", info.ClusterName, "!")

	// data[0].Project_id = "keti-container"
	// data[0].Cluster_name = info.ClusterName
	// data[0].Region = "us-central1-a"
	// data[0].Gke_num_nodes = uint64(num)

	// doc, _ := json.Marshal(data)

	// fmt.Println(strings.Trim(string(doc), "[]"))

	// err := ioutil.WriteFile("/root/go/src/Hybrid_Cloud/terraform/gke/create/", []byte(strings.Trim(string(doc), "[]")), os.FileMode(0644))

	// if err != nil {
	// 	panic(err)
	// }

	cmd_rm_cluster := exec.Command("rm", info.ClusterName+".tf.json")
	cmd_rm_cluster.Dir = "../terraform/aks"

	output, err := cmd_rm_cluster.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	// cmd := exec.Command("terraform", "destroy", "-auto-approve")

	// cmd1 := exec.Command("terraform", "plan", "-lock=false")
	// cmd1.Dir = "../terraform/aks"
	// output, err = cmd1.Output()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(string(output))
	// }
	cmd2 := exec.Command("terraform", "apply", "-auto-approve")

	cmd2.Dir = "../terraform/aks"

	output, err = cmd2.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

}

func delete_gke(info Cli) {
	num := 1
	data := make([]Cluster_info, 1)
	cluster := "cluster"

	fmt.Println("!", info.ClusterName, "!")

	data[0].Project_id = "keti-container"
	data[0].Cluster_name = info.ClusterName
	data[0].Region = "us-central1-a"
	data[0].Gke_num_nodes = uint64(num)

	doc, _ := json.Marshal(data)

	fmt.Println(strings.Trim(string(doc), "[]"))

	// err := ioutil.WriteFile("/root/go/src/Hybrid_Cloud/terraform/gke/create/", []byte(strings.Trim(string(doc), "[]")), os.FileMode(0644))

	// if err != nil {
	// 	panic(err)
	// }

	cmd_rm_cluster := exec.Command("rm", info.ClusterName+".tf.json")
	cmd_rm_cluster.Dir = "../terraform/gke/" + cluster

	output, err := cmd_rm_cluster.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	cmd_rm_nodepool := exec.Command("rm", info.ClusterName+"nodePool"+".tf.json")
	cmd_rm_nodepool.Dir = "../terraform/gke/" + cluster

	output, err = cmd_rm_nodepool.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	//cmd := exec.Command("terraform", "destroy", "-auto-approve")
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = "../terraform/gke/" + cluster

	output, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

}

func delete_eks(info Cli) {
	cmd_rm_cluster := exec.Command("rm", info.ClusterName+".tf.json")
	cmd_rm_cluster.Dir = "../terraform/eks"

	output, err := cmd_rm_cluster.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	cmd_rm_nodepool := exec.Command("rm", info.ClusterName+"nodepool.tf.json")
	cmd_rm_nodepool.Dir = "../terraform/eks"

	output, err = cmd_rm_nodepool.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	//cmd := exec.Command("terraform", "destroy", "-auto-approve")
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = "../terraform/eks"

	output, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}

func init() {
	RootCmd.AddCommand(DeleteCmd)
	DeleteCmd.AddCommand(deleteClusterCmd)
	DeleteCmd.AddCommand(deleteNodeCmd)
	DeleteCmd.Flags().StringVarP(&util.Option_file, "file", "f", "", "FILENAME")
	DeleteCmd.MarkFlagRequired("file")
	DeleteCmd.Flags().StringVarP(&util.Option_context, "context", "", "", "CLUSTERNAME")
	DeleteCmd.Flags().StringP("namespace", "n", "default", "enter the namespace")

	deleteClusterCmd.Flags().String("platform", "", "input your platform name")
	deleteClusterCmd.Flags().String("cluster-name", "", "input your cluster name")

	// deleteNodeCmd.Flags().String("nodepool-name", "", "input your nodepool name")
	deleteNodeCmd.Flags().String("nodepool-name", "", "input your nodepool name")
	deleteNodeCmd.Flags().String("platform", "", "input your platform name")
	deleteNodeCmd.Flags().String("cluster-name", "", "input your cluster name")
}
