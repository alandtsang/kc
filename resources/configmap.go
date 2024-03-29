package resources

import (
	"fmt"
	"log"

	"k8s.io/api/core/v1"
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
		cm.getConfigMap()
	} else {
		cm.getConfigMapList()
	}
}

func (cm *ConfigMaps) Delete() {
	err := cm.clientSet.CoreV1().ConfigMaps(cm.namespace).Delete(cm.name, &metav1.DeleteOptions{})
	if err != nil {
		log.Fatalf("[Error] Delete configmaps %s failed: %s\n", cm.name, err.Error())
	} else {
		fmt.Printf("configmaps \"%s\" deleted\n", cm.name)
	}
}

func (cm *ConfigMaps) getConfigMap() {
	configMap, err := cm.clientSet.CoreV1().ConfigMaps(cm.namespace).Get(cm.name, metav1.GetOptions{})
	if err != nil {
		log.Fatalln("[Error] GetConfigMap()", err.Error())
	}
	cm.printConfigMap(configMap)
}

func (cm *ConfigMaps) getConfigMapList() {
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
