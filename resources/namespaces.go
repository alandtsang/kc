package resources

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetNamespaces(clientset *kubernetes.Clientset) {
	namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("NAME\t\t\t\t STATUS\t\t AGE\n")
	for _, namespace := range namespaces.Items {
		fmt.Printf("%-30s\t %s\t\t %s\n", namespace.Name, namespace.Status.Phase, namespace.CreationTimestamp)
	}
}
