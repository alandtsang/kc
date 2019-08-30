package resources

import (
	"fmt"

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
	configMaps, err := cm.clientSet.CoreV1().ConfigMaps(cm.namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	if len(configMaps.Items) > 0 {
		fmt.Printf("NAME\t\t\t DATA\t AGE\n")
		for _, configMap := range configMaps.Items {
			fmt.Printf("%-24s %d\t %s\n", configMap.Name, len(configMap.Data), configMap.CreationTimestamp)
		}
	} else {
		fmt.Println(ERR_NO_RESOURCE)
	}
}
