package resources

import (
	"log"

	"github.com/alandtsang/kc/clientset"
)

const ERR_NO_RESOURCE string = "No resources found."

type Action int

const (
	ActionCreate Action = iota
	ActionExpose
	ActionRun
	ActionSet
	ActionExplain
	ActionGet
	ActionEdit
	ActionDelete
	ActionRollout
	ActionScale
	ActionAutoscale
	ActionCertificate
	ActionClusterInfo
	ActionTop
	ActionCordon
	ActionUncordon
	ActionDrain
	ActionTaint
	ActionDescribe
	ActionLogs
	ActionAttach
	ActionExec
	ActionPortForward
	ActionProxy
	ActionCp
	ActionAuth
	ActionApply
	ActionPatch
	ActionReplace
	ActionWait
	ActionConvert
	ActionLabel
	ActionAnnotate
	ActionCompletion
	ActionAlpha
	ActionApiResources
	ActionApiVersions
	ActionConfig
	ActionPlugin
	ActionVersion
)

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
	Delete()
}

func NewResourcesManager(action Action, client *clientset.Client, namespace, resource, name string) *ResourcesManager {
	resourcer := NewResourcer(client, namespace, resource, name)
	return &ResourcesManager{resourcer: resourcer}
}

func NewResourcer(client *clientset.Client, namespace, resource, name string) Resourcer {
	var resourcer Resourcer
	switch resource {
	case "ns":
		resourcer = NewNamespaces(client.ClientSet, name)
	case "node":
		resourcer = NewNodes(client.ClientSet)
	case "pod":
		resourcer = NewPods(client.ClientSet, namespace, name)
	case "svc":
		resourcer = NewServices(client.ClientSet, namespace, name)
	case "sa":
		resourcer = NewServiceAccounts(client.ClientSet, namespace, name)
	case "sec":
		resourcer = NewSecrets(client.ClientSet, namespace, name)
	case "cm":
		resourcer = NewConfigMaps(client.ClientSet, namespace, name)
	case "ep":
		resourcer = NewEndPoints(client.ClientSet, namespace, name)
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