package resources

import "github.com/alandtsang/kc/clientset"

type ResourcesManager struct {
	rs     *Resources
	client *clientset.Client
}
type Resources struct {
	namespaces *Namespaces
	nodes      *Nodes
	pods       *Pods
	svcs       *Services
	sa         *ServiceAccounts
	secret     *Secrets
	configMaps *ConfigMaps
	ep         *EndPoints
}

func NewResourcesManager(client *clientset.Client) *ResourcesManager {
	return &ResourcesManager{rs: &Resources{}, client: client}
}

func (rsm *ResourcesManager) GetNamespaces() {
	rsm.rs.namespaces = NewNamespaces(rsm.client.ClientSet)
	rsm.rs.namespaces.GetNamespaces()
}
