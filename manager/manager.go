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

func (kcm *KCManager) Init(clusterName, namespace, resource, name string) {
	kcm.cs = clusters.NewClusterSet()
	cluster := kcm.cs.GetCluster(clusterName)
	if cluster == nil {
		log.Fatalf("[Error] Cluster name %s not found, please enter the correct name\n", clusterName)
	}
	//fmt.Printf("%+v\n", *cluster)
	kcm.client = clientset.NewClient(cluster.Addr, cluster.Token)
	kcm.resManager = resources.NewResourcesManager(kcm.client, namespace, resource, name)
	kcm.namespace = namespace
}

func (kcm *KCManager) GetResource() {
	kcm.resManager.Get()
}
