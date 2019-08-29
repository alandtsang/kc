package resources

import (
	"log"

	"github.com/alandtsang/kc/clientset"
)

const ERR_NO_RESOURCE string = "No resources found."

var ResourcesLists = map[string]bool{
	"ns":   true,
	"node": true,
	"pod":  true,
	"svc":  true,
	"sec":  true,
	"sa":   true,
	"cm":   true,
	"ep":   true,
}

type Resourcer interface {
	Get()
}

func NewResourcesManager(client *clientset.Client, namespace, resource, name string) *ResourcesManager {
	rss := NewResource(client, namespace, resource)
	return &ResourcesManager{rss: rss}
}

func NewResource(client *clientset.Client, namespace, resource string) Resourcer {
	var resourcer Resourcer
	switch resource {
	case "ns":
		resourcer = NewNamespaces(client.ClientSet)
	case "node":
		resourcer = NewNodes(client.ClientSet)
	case "pod":
		resourcer = NewPods(client.ClientSet, namespace)
	case "svc":
		resourcer = NewPods(client.ClientSet, namespace)
	case "sa":
		resourcer = NewPods(client.ClientSet, namespace)
	case "sec":
		resourcer = NewPods(client.ClientSet, namespace)
	case "cm":
		resourcer = NewPods(client.ClientSet, namespace)
	case "ep":
		resourcer = NewPods(client.ClientSet, namespace)
	default:
		log.Fatal("[Error] NewResource")
	}
	return resourcer
}

type ResourcesManager struct {
	rss Resourcer
}

func (rsm *ResourcesManager) Get() {
	rsm.rss.Get()
}
