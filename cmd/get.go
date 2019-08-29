/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"fmt"
	"log"

	"github.com/alandtsang/kc/manager"
	"github.com/alandtsang/kc/resources"
	"github.com/spf13/cobra"
)

var (
	getLong = `
Display one or many resources

Prints a table of the most important information about the specified resources.
You can filter the list using a label selector and the --selector flag. If the
desired resource type is namespaced you will only see results in your current
namespace unless you pass --all-namespaces.

Uninitialized objects are not shown unless --include-uninitialized is passed.

By specifying the output as 'template' and providing a Go template as the value
of the --template flag, you can filter the attributes of the fetched resources.`

	getExample = `
  # List all pods in ps output format.
  kubectl get pods

  # List all pods in ps output format with more information (such as node name).
  kubectl get pods -o wide

  # List a single replication controller with specified NAME in ps output format.
  kubectl get replicationcontroller web

  # List deployments in JSON output format, in the "v1" version of the "apps" API group:
  kubectl get deployments.v1.apps -o json

  # List a single pod in JSON output format.
  kubectl get -o json pod web-pod-13je7

  # List a pod identified by type and name specified in "pod.yaml" in JSON output format.
  kubectl get -f pod.yaml -o json

  # List resources from a directory with kustomization.yaml - e.g. dir/kustomization.yaml.
  kubectl get -k dir/

  # Return only the phase value of the specified pod.
  kubectl get -o template pod/web-pod-13je7 --template={{.status.phase}}

  # List resource information in custom columns.
  kubectl get pod test-pod -o custom-columns=CONTAINER:.spec.containers[0].name,IMAGE:.spec.containers[0].image

  # List all replication controllers and services together in ps output format.
  kubectl get rc,services

  # List one or more resources by their type and names.
  kubectl get rc/web service/frontend pods/web-pod-13je7`
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Display one or many resources",
	Long:    getLong,
	Example: getExample,
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("namespace")
		getArgs := cmd.Flags().Args()
		validate(getArgs)
		do(namespace, getArgs)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("namespace", "n", "default", "help for namespace")
	getCmd.Flags().StringP("output", "o", "", "help for output")
}

func validate(args []string) {
	argsLen := len(args)
	if argsLen == 0 {
		log.Fatalln("[Error] Empty resource names are not allowed")
	} else if argsLen > 2 {
		log.Fatalln("[Error] Too many parameters are not allowed")
	}

	if _, ok := resources.ResourcesLists[args[0]]; !ok {
		log.Fatalf("[Error] Resource name %s are not allowed\n", args[0])
	}
}

func do(namespace string, args []string) {
	var resource, name string
	resource = args[0]
	if len(args) == 2 {
		name = args[1]
	}
	fmt.Println(namespace, resource, name)
	clusterName := "k8s-test2"
	var kcmanager manager.KCManager
	kcmanager.Init(clusterName, namespace, resource, name)
	kcmanager.GetResource()
}
