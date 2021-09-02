// Copyright © 2021 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	analyticEngine "Hybrid_Cluster/hcp-analytic-engine/analyticEngine"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// GKE terraform structure

type Cluster_info struct {
	Project_id    string `json:"project_id"`
	Cluster_name  string `json:"cluster_name"`
	Region        string `json:"region"`
	Gke_num_nodes uint64 `json:"gke_num_nodes"`
}

type TF struct {
	Resource *Resource `json:"resource"`
}

type TF_AKS struct {
	ResourceAksCluster *ResourceAksCluster `json:"resource"`
}

type ResourceAksCluster struct {
	AzurernKubernetesCluster *map[string]AksCluster `json:"azurerm_kubernetes_cluster"`
}

type Resource struct {
	Google_container_cluster *map[string]Cluster_type `json:"google_container_cluster"`
}

type SSHKey struct {
	KeyData string `json:"key_data"`
}
type LinuxProfile struct {
	AdminUsername string `json:"admin_username"`
	SSHKey        SSHKey `json:"ssh_key"`
}
type DefaultNodePool struct {
	Name      string `json:"name"`
	NodeCount int    `json:"node_count"`
	VMSize    string `json:"vm_size"`
}
type ServicePrincipal struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
type OmsAgent struct {
	Enabled                 string `json:"enabled"`
	LogAnalyticsWorkspaceID string `json:"log_analytics_workspace_id"`
}
type AddonProfile struct {
	OmsAgent OmsAgent `json:"oms_agent"`
}
type NetworkProfile struct {
	LoadBalancerSku string `json:"load_balancer_sku"`
	NetworkPlugin   string `json:"network_plugin"`
}
type Tags struct {
	Environment string `json:"Environment"`
}
type AksCluster struct {
	Name              string           `json:"name"`
	Location          string           `json:"location"`
	ResourceGroupName string           `json:"resource_group_name"`
	DNSPrefix         string           `json:"dns_prefix"`
	LinuxProfile      LinuxProfile     `json:"linux_profile"`
	DefaultNodePool   DefaultNodePool  `json:"default_node_pool"`
	ServicePrincipal  ServicePrincipal `json:"service_principal"`
	AddonProfile      AddonProfile     `json:"addon_profile"`
	NetworkProfile    NetworkProfile   `json:"network_profile"`
	Tags              Tags             `json:"tags"`
}
type AzurermKubernetesCluster struct {
	AksCluster AksCluster `json:"aks-cluster"`
}

type Cluster_type struct {
	Name                     string `json:"name"`
	Location                 string `json:"location"`
	Remove_default_node_pool string `json:"remove_default_node_pool"`
	Initial_node_count       int    `json:"initial_node_count"`
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("Create Cluster Called")
		cli := Cli{args[0], args[1]}

		if len(args) == 0 {
			//fmt.Println("Run 'hybridctl create --help' to view all commands") //나중에 다시 살릴것
			create_gke_data_path()

		} else if args[0] == "gke" {
			if args[1] == "" {

				create_gke(cli)
				fmt.Println("Run 'hybridctl create --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name : ", args[0])
				fmt.Printf("Cluster Name : %s\n", args[1])

				create_gke(cli)
			}
		} else if args[0] == "eks" {
			if args[1] == "" {
				fmt.Println("Run 'hybridctl create --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name : ", args[0])
				fmt.Printf("Cluster Name : %s\n", args[1])
			}

		} else if args[0] == "aks" {
			if args[1] == "" {
				fmt.Println("Run 'hybridctl create --help' to view all commands")
			} else {
				fmt.Println("kubernetes engine Name : ", args[0])
				fmt.Printf("Cluster Name : %s\n", args[1])

				create_aks(cli)
			}
		}
		// create_gke_data_path()
	},
}

// func create_eks(info Cli) {

// }

func create_aks(info Cli) {

	send_js_cluster := TF_AKS{
		ResourceAksCluster: &ResourceAksCluster{
			AzurernKubernetesCluster: &map[string]AksCluster{
				info.ClusterName: {
					Name:              info.ClusterName,
					Location:          "${azurerm_resource_group.k8s.location}",
					ResourceGroupName: "${azurerm_resource_group.k8s.name}",
					DNSPrefix:         "${var.dns_prefix}",
					LinuxProfile: LinuxProfile{
						AdminUsername: "ubuntu",
						SSHKey: SSHKey{
							KeyData: "${file(var.ssh_public_key)}",
						},
					},
					DefaultNodePool: DefaultNodePool{
						Name:      "agentpool",
						NodeCount: 1,
						VMSize:    "Standard_D2_v2",
					},
					ServicePrincipal: ServicePrincipal{
						ClientID:     "${var.client_id}",
						ClientSecret: "${var.client_secret}",
					},
					AddonProfile: AddonProfile{
						OmsAgent: OmsAgent{
							Enabled:                 "true",
							LogAnalyticsWorkspaceID: "${azurerm_log_analytics_workspace.test.id}",
						},
					},
					NetworkProfile: NetworkProfile{
						LoadBalancerSku: "Standard",
						NetworkPlugin:   "kubenet",
					},
					Tags: Tags{
						Environment: "Development",
					},
				},
			},
		},
	}
	send, err := json.MarshalIndent(send_js_cluster, "", " ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("/root/go/src/Hybrid_Cluster/terraform/aks/"+info.ClusterName+".tf.json", []byte(string(send)), os.FileMode(0644))
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("terraform", "apply", "-auto-approve")
	// cmd := exec.Command("terraform", "plan")
	cmd.Dir = "../terraform/aks/"

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

}

func create_gke_data_path() {

	cluster := "cluster"
	num := 1

	clusterName := "cluster-2"

	// data[0].Project_id = "keti-container"
	// data[0].Cluster_name = clusterName
	// data[0].Region = "us-central1-a"
	// data[0].Gke_num_nodes = uint64(num)

	fmt.Println("Policy Engine Checking")
	fmt.Println("Send Result to User Requirement Checking Module")
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("[Option2] Policy is nonexistent")
	fmt.Println("Call Analytic Engine")
	analyticEngine.HybridctlAnalyticEngine()

	// doc, _ := json.Marshal(data)

	// fmt.Println(strings.Trim(string(doc), "[]"))

	// err := ioutil.WriteFile("/root/go/src/Hybrid_Cluster/terraform/gke/"+cluster+"/"+info.ClusterName+".tfvars.json", []byte(strings.Trim(string(doc), "[]")), os.FileMode(0644))

	// if err != nil {
	// 	panic(err)
	// }

	send_js_cluster := TF{
		Resource: &Resource{
			Google_container_cluster: &map[string]Cluster_type{
				clusterName: {
					Name:                     clusterName,
					Location:                 "us-central1-a",
					Remove_default_node_pool: "true",
					Initial_node_count:       num,
				},
			},
		},
	}

	send, err := json.MarshalIndent(send_js_cluster, "", " ")
	if err != nil {
		panic(err)
	}

	// src, _ := json.Marshal([]byte(string(resource)))

	err = ioutil.WriteFile("/root/Hybrid_Cluster/terraform/gke/"+cluster+"/"+clusterName+".tf.json", []byte(string(send)), os.FileMode(0644))
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("terraform", "apply", "-auto-approve")
	//cmd := exec.Command("terraform", "plan")
	cmd.Dir = "../terraform/gke/cluster"

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}

func create_gke(info Cli) {

	cluster := "cluster"
	num := 1
	data := make([]Cluster_info, 1)

	data[0].Project_id = "keti-container"
	data[0].Cluster_name = info.ClusterName
	data[0].Region = "us-central1-a"
	data[0].Gke_num_nodes = uint64(num)

	fmt.Println("Policy Engine Checking")
	fmt.Println("Send Result to User Requirement Checking Module")
	fmt.Println("[Option1] Policy exist")
	fmt.Println("---Create Cluster Start---")

	// doc, _ := json.Marshal(data)

	// fmt.Println(strings.Trim(string(doc), "[]"))

	// err := ioutil.WriteFile("/root/go/src/Hybrid_Cluster/terraform/gke/"+cluster+"/"+info.ClusterName+".tfvars.json", []byte(strings.Trim(string(doc), "[]")), os.FileMode(0644))

	// if err != nil {
	// 	panic(err)
	// }

	send_js_cluster := TF{
		Resource: &Resource{
			Google_container_cluster: &map[string]Cluster_type{
				info.ClusterName: {
					Name:                     info.ClusterName,
					Location:                 "us-central1-a",
					Remove_default_node_pool: "true",
					Initial_node_count:       num,
				},
			},
		},
	}

	send, err := json.MarshalIndent(send_js_cluster, "", " ")
	if err != nil {
		panic(err)
	}

	// src, _ := json.Marshal([]byte(string(resource)))

	err = ioutil.WriteFile("/root/Hybrid_Cluster/terraform/gke/"+cluster+"/"+info.ClusterName+".tf.json", []byte(string(send)), os.FileMode(0644))
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("terraform", "apply", "-auto-approve")
	//cmd := exec.Command("terraform", "plan")
	cmd.Dir = "../terraform/gke/cluster"

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}

func init() {
	RootCmd.AddCommand(createCmd)

	// flag.IntVar(&flagvar, "node", 1, "set node count")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
