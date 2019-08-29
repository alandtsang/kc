package main

import (
	"github.com/alandtsang/kc/cmd"
)

func main() {
	cmd.Execute()

	/*var kcmanager manager.KCManager
	//clusterName := "k8s-test2"
	clusterName := os.Args[1]
	namespace := ""
	kcmanager.Init(clusterName, namespace)

	//kcmanager.GetNamespaces()
	//kcmanager.GetPods()
	//kcmanager.GetConfigMaps()
	//kcmanager.GetEndPoints()
	//kcmanager.GetNodes()
	//kcmanager.GetServiceAccounts()
	//kcmanager.GetSecrets()
	kcmanager.GetServices()*/
}
