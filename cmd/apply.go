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
	"log"

	"github.com/alandtsang/kc/clientset"
	"github.com/spf13/cobra"
)

var body = []byte(`{
"apiVersion": "v1"
"kind": "ConfigMap"
"metadata":
    "name": "test-cm"
    "namespace": "default"
}`)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a configuration to a resource by filename or stdin.",
	Long: `Apply a configuration to a resource by filename or stdin. The resource name must be specified. This resource will be
created if it doesn't exist yet. To use 'apply', always create the resource initially with either 'apply' or 'create
--save-config'.

JSON and YAML formats are accepted.`,
	Run: func(cmd *cobra.Command, args []string) {
		clusterName := "k8s-test2"
		client := clientset.NewClient(clusterName).Get()
		//_, err := client.CoreV1().RESTClient().Patch(types.ApplyPatchType).
		_, err := client.CoreV1().RESTClient().Post().
			Namespace("default").
			Resource("configmaps").
			SetHeader("Accept", "application/json").
			SetHeader("Content-Type", "application/json").
			Name("test-cm").
			//Param("fieldManager", "apply_test").
			Body(body).
			Do().
			Get()
		if err != nil {
			log.Fatalf("[Error] Failed to create object using Apply patch: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
	applyCmd.Flags().StringP("yaml or json file", "f", "", "help for apply")
}
