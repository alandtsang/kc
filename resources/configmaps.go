package resources

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetConfigMaps(clientset *kubernetes.Clientset) {
	namespace := "logging"
	configMaps, err := clientset.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("NAME\t\t\t DATA\t AGE\n")
	for _, configMap := range configMaps.Items {
		fmt.Printf("%-24s %d\t %s\n", configMap.Name, len(configMap.Data), configMap.CreationTimestamp)
	}
}
