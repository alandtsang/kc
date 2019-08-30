package manager

import (
	"log"

	"github.com/alandtsang/kc/clientset"
	"github.com/alandtsang/kc/resources"
)

type KCManager struct {
	client     *clientset.Client
	action     resources.Action
	namespace  string
	resManager *resources.ResourcesManager
}

func (kcm *KCManager) Init(action resources.Action, clusterName, namespace, resource, name string) {
	kcm.client = clientset.NewClient(clusterName)
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
