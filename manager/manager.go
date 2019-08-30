package manager

import (
	"log"

	"github.com/alandtsang/kc/clientset"
	"github.com/alandtsang/kc/clusters"
	"github.com/alandtsang/kc/resources"
)

type KCManager struct {
	cs         *clusters.ClusterSet
	client     *clientset.Client
	resManager *resources.ResourcesManager
	namespace  string
	action     resources.Action
}

func (kcm *KCManager) Init(action resources.Action, clusterName, namespace, resource, name string) {
	kcm.cs = clusters.NewClusterSet()
	cluster := kcm.cs.GetCluster(clusterName)
	if cluster == nil {
		log.Fatalf("[Error] Cluster name %s not found, please enter the correct name\n", clusterName)
	}
	//fmt.Printf("%+v\n", *cluster)
	kcm.client = clientset.NewClient(cluster.Addr, cluster.Token)
	kcm.namespace = namespace
	kcm.action = action
	if kcm.action == resources.ActionLogs {
		resource = "pod"
	}
	kcm.resManager = resources.NewResourcesManager(action, kcm.client, namespace, resource, name)

}

func (kcm *KCManager) Do() {
	switch kcm.action {
	case resources.ActionGet:
		kcm.GetResource()
	case resources.ActionLogs:
		kcm.GetLogs()
	default:
		log.Fatalln("Wrong Action")
	}
}

func (kcm *KCManager) GetResource() {
	kcm.resManager.Get()
}

func (kcm *KCManager) GetLogs() {
	kcm.resManager.GetLogs()
}
