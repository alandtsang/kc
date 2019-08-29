package resources

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Services struct {
}

func GetService(clientset *kubernetes.Clientset) {
	namespace := "default"
	services, err := clientset.CoreV1().Services(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("NAME\t\t TYPE\t\t CLUSTER-IP\t EXTERNAL-IP\t PORT(S)\t AGE\n")
	for _, service := range services.Items {
		fmt.Printf("%-16s %s\t %s\t %v\t\t %v\t %s\n",
			service.Name, service.Spec.Type, service.Spec.ClusterIP,
			service.Spec.ExternalIPs, service.Spec.Ports, service.CreationTimestamp)
	}
}
