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

	"github.com/alandtsang/kc/action"
	"github.com/alandtsang/kc/resources"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete resources by filenames, stdin, resources and names, or by resources and label selector",
	Long: `Delete resources by filenames, stdin, resources and names, or by resources and label selector.

JSON and YAML formats are accepted. Only one type of the arguments may be specified: filenames, resources and names, or
resources and label selector.

Some resources, such as pods, support graceful deletion. These resources define a default period before they are
forcibly terminated (the grace period) but you may override that value with the --grace-period flag, or pass --now to
set a grace-period of 1. Because these resources often represent entities in the cluster, deletion may not be
acknowledged immediately. If the node hosting a pod is down or cannot reach the API server, termination may take
significantly longer than the grace period. To force delete a resource, you must pass a grace period of 0 and specify
the --force flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("namespace")
		getArgs := cmd.Flags().Args()
		deleteValidate(getArgs)
		do(action.ActionDelete, namespace, getArgs)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("namespace", "n", "default", "help for namespace")
	deleteCmd.Flags().StringP("filename", "f", "", "help for filename")
}

func deleteValidate(args []string) {
	argsLen := len(args)
	if argsLen == 0 {
		log.Fatalln("[Error] Empty resource are not allowed")
	} else if argsLen < 2 {
		log.Fatalln("[Error] Empty resource name are not allowed")
	} else if argsLen > 2 {
		log.Fatalln("[Error] Too many parameters are not allowed")
	}

	if _, ok := resources.ResourcesLists[args[0]]; !ok {
		log.Fatalf("[Error] Resource name %s are not allowed\n", args[0])
	}
}
