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

	"github.com/alandtsang/kc/resources"
	"github.com/spf13/cobra"
)

const (
	log_example = `# Return snapshot logs from pod nginx with only one container
$ kubectl logs nginx

# Return snapshot of previous terminated ruby container logs from pod web-1
$ kubectl logs -p -c ruby web-1

# Begin streaming the logs of the ruby container in pod web-1
$ kubectl logs -f -c ruby web-1

# Display only the most recent 20 lines of output in pod nginx
$ kubectl logs --tail=20 nginx

# Show all logs from pod nginx written in the last hour
$ kubectl logs --since=1h nginx`
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "A brief description of your command",
	Long: `Print the logs for a container in a pod or specified resource. If the pod has only one container, the container name is
optional.`,
	Example: log_example,
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("namespace")
		logsArgs := cmd.Flags().Args()
		logsValidate(logsArgs)
		logsDo(namespace, logsArgs)
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
	logsCmd.Flags().StringP("namespace", "n", "default", "help for namespace")
	logsCmd.Flags().BoolP("follow", "f", false, "Specify if the logs should be streamed.")
	logsCmd.Flags().Bool("timestamps", false, "Include timestamps on each line in the log output")
	logsCmd.Flags().Bool("interactive", true, "If true, prompt the user for input when required. Default true.")
	logsCmd.Flags().BoolP("previous", "p", false, "If true, print the logs for the previous instance of the container in a pod if it exists.")
	logsCmd.Flags().Int("limit-bytes", 0, "Maximum bytes of logs to return. Defaults to no limit.")
	logsCmd.Flags().Int("tail", -1, "Lines of recent log file to display. Defaults to -1, showing all log lines.")
	logsCmd.Flags().String("since-time", "", "Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used.")
	logsCmd.Flags().Duration("since", 0, "Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one of since-time / since may be used.")
	//logsCmd.Flags().StringVarP(&params.containerName, "container", "c", "", "Container name")
}

func logsValidate(args []string) {
	argsLen := len(args)
	if argsLen == 0 {
		log.Fatalln("[Error] Empty pod name are not allowed")
	} else if argsLen > 1 {
		log.Fatalln("[Error] Too many parameters are not allowed")
	}
}

func logsDo(namespace string, args []string) {
	podName := args[0]
	NewKCManager(resources.ActionLogs, namespace, "", podName)
}
