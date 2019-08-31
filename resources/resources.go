package resources

import (
	"log"

	"github.com/alandtsang/kc/clientset"
)

const ERR_NO_RESOURCE string = "No resources found."

var ResourcesLists = map[string]bool{
	"cm":   true,
	"ep":   true,
	"ns":   true,
	"node": true,
	"pod":  true,
	"sa":   true,
	"sec":  true,
	"svc":  true,
}

type Resourcer interface {
	Get()
	Delete()
}

func NewResourcesManager(client *clientset.Client, namespace, resource, name string) *ResourcesManager {
	resourcer := NewResourcer(client, namespace, resource, name)
	return &ResourcesManager{resourcer: resourcer}
}

func NewResourcer(client *clientset.Client, namespace, resource, name string) Resourcer {
	var resourcer Resourcer
	switch resource {
	case "cm":
		resourcer = NewConfigMaps(client.ClientSet, namespace, name)
	case "ep":
		resourcer = NewEndPoints(client.ClientSet, namespace, name)
	case "ns":
		resourcer = NewNamespaces(client.ClientSet, name)
	case "node":
		resourcer = NewNodes(client.ClientSet)
	case "pod":
		resourcer = NewPods(client.ClientSet, namespace, name)
	case "sa":
		resourcer = NewServiceAccounts(client.ClientSet, namespace, name)
	case "sec":
		resourcer = NewSecrets(client.ClientSet, namespace, name)
	case "svc":
		resourcer = NewServices(client.ClientSet, namespace, name)
	default:
		log.Fatal("[Error] Unsupported resource type")
	}
	return resourcer
}

type ResourcesManager struct {
	resourcer Resourcer
}

func (rsm *ResourcesManager) Get() {
	rsm.resourcer.Get()
}

func (rsm *ResourcesManager) GetLogs() {
	rsm.resourcer.(*Pods).GetLogs()
}

func (rsm *ResourcesManager) Delete() {
	rsm.resourcer.Delete()
}
