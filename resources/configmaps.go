package resources

import (
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigMaps struct {
	clientSet *kubernetes.Clientset
	namespace string
	name      string
}

func NewConfigMaps(clientSet *kubernetes.Clientset, namespace, name string) *ConfigMaps {
	return &ConfigMaps{clientSet: clientSet, namespace: namespace, name: name}
}

func (cm *ConfigMaps) Get() {
	if len(cm.name) > 0 {
		cm.GetConfigMap()
	} else {
		cm.GetConfigMapList()
	}
}

func (cm *ConfigMaps) GetConfigMap() {
	configMap, err := cm.clientSet.CoreV1().ConfigMaps(cm.namespace).Get(cm.name, metav1.GetOptions{})
	if err != nil {
		log.Fatalln("[Error] GetConfigMap()", err.Error())
	}
	cm.printConfigMap(configMap)
}

func (cm *ConfigMaps) GetConfigMapList() {
	configMaps, err := cm.clientSet.CoreV1().ConfigMaps(cm.namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("[Error] GetConfigMaps()", err.Error())
	}
	cm.printConfigMaps(configMaps)
}

func (cm *ConfigMaps) printConfigMap(configMap *v1.ConfigMap) {
	fmt.Printf("NAME\t\t\t DATA\t AGE\n")
	fmt.Printf("%-24s %d\t %s\n", configMap.Name, len(configMap.Data), configMap.CreationTimestamp)
}

func (cm *ConfigMaps) printConfigMaps(configMaps *v1.ConfigMapList) {
	if len(configMaps.Items) > 0 {
		fmt.Printf("NAME\t\t\t DATA\t AGE\n")
		for _, configMap := range configMaps.Items {
			fmt.Printf("%-24s %d\t %s\n", configMap.Name, len(configMap.Data), configMap.CreationTimestamp)
		}
	} else {
		fmt.Println(ERR_NO_RESOURCE)
	}
}
