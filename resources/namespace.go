package resources

import (
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Namespaces struct {
	clientSet *kubernetes.Clientset
	name      string
}

func NewNamespaces(clientSet *kubernetes.Clientset, name string) Resourcer {
	return &Namespaces{clientSet: clientSet, name: name}
}

func (n *Namespaces) Get() {
	namespaces, err := n.clientSet.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("[Error] GetNamespaces() failed: %s\n", err.Error())
	}

	fmt.Printf("NAME\t\t\t\t STATUS\t\t AGE\n")
	for _, namespace := range namespaces.Items {
		fmt.Printf("%-30s\t %s\t\t %s\n", namespace.Name, namespace.Status.Phase, namespace.CreationTimestamp)
	}
}

func (n *Namespaces) Delete() {
	err := n.clientSet.CoreV1().Namespaces().Delete(n.name, &metav1.DeleteOptions{})
	if err != nil {
		log.Fatalf("[Error] Delete namespace %s failed: %s\n", n.name, err.Error())
	} else {
		fmt.Printf("namespace \"%s\" deleted\n", n.name)
	}
}
