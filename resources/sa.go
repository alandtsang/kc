package resources

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ServiceAccounts struct {
	clientSet *kubernetes.Clientset
	namespace string
	name      string
}

func NewServiceAccounts(clientSet *kubernetes.Clientset, namespace, name string) *ServiceAccounts {
	return &ServiceAccounts{clientSet: clientSet, namespace: namespace, name: name}
}

func (sa *ServiceAccounts) Get() {
	serviceAccounts, err := sa.clientSet.CoreV1().ServiceAccounts(sa.namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	if len(serviceAccounts.Items) > 0 {
		fmt.Printf("NAME\t\t\t SECRETS\t AGE\n")
		for _, sa := range serviceAccounts.Items {
			fmt.Printf("%-24s %d\t %s\n", sa.Name, len(sa.Secrets), sa.CreationTimestamp)
		}
	} else {
		fmt.Println(ERR_NO_RESOURCE)
	}
}
