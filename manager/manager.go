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
}

func (kcm *KCManager) Init(clusterName, namespace string) {
	kcm.cs = clusters.NewClusterSet()
	cluster := kcm.cs.GetCluster(clusterName)
	if cluster == nil {
		log.Fatalf("[Error] Cluster name %s not found, please enter the correct name\n", clusterName)
	}
	//fmt.Printf("%+v\n", *cluster)
	kcm.client = clientset.NewClient(cluster.Addr, cluster.Token)
	kcm.resManager = resources.NewResourcesManager(kcm.client)
	if len(namespace) == 0 {
		kcm.namespace = "default"
	} else {
		kcm.namespace = namespace
	}

	//resources.GetNamespaces(clientset)
	//resources.GetPods(clientset)
	//resources.GetService(clientset)
	//resources.GetConfigMaps(clientset)
	//resources.GetNodes(clientset)
	//resources.GetEndPoints(clientset)
	//resources.GetSecrets(clientset)
	//resources.GetServiceAccounts(clientset)
}

func (kcm *KCManager) GetNamespaces() {
	kcm.resManager.GetNamespaces()
}

func (kcm *KCManager) GetPods() {
	kcm.resManager.GetPods(kcm.namespace)
}

func (kcm *KCManager) GetConfigMaps() {
	kcm.resManager.GetConfigMaps(kcm.namespace)
}

func (kcm *KCManager) GetEndPoints() {
	kcm.resManager.GetEndPoints(kcm.namespace)
}

func (kcm *KCManager) GetNodes() {
	kcm.resManager.GetNodes()
}
