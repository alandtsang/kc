package manager

import (
	"log"

	"github.com/alandtsang/kc/action"
	"github.com/alandtsang/kc/clientset"
	"github.com/alandtsang/kc/resources"
)

type KCManager struct {
	client     *clientset.Client
	action     action.Action
	namespace  string
	resManager *resources.ResourcesManager
}

func (kcm *KCManager) Init(act action.Action, clusterName, namespace, resource, name string) {
	kcm.client = clientset.NewClient(clusterName)
	kcm.namespace = namespace
	kcm.action = act
	if kcm.action == action.ActionLogs {
		resource = "pod"
	}
	kcm.resManager = resources.NewResourcesManager(kcm.client, namespace, resource, name)
}

func (kcm *KCManager) Do() {
	switch kcm.action {
	case action.ActionGet:
		kcm.GetResource()
	case action.ActionLogs:
		kcm.GetLogs()
	case action.ActionDelete:
		kcm.DeleteResource()
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

func (kcm *KCManager) DeleteResource() {
	kcm.resManager.Delete()
}
