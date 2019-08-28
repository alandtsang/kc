package resources

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetServiceAccounts(clientset *kubernetes.Clientset) {
	namespace := "default"
	serviceAccounts, err := clientset.CoreV1().ServiceAccounts(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("NAME\t\t\t SECRETS\t AGE\n")
	for _, sa := range serviceAccounts.Items {
		fmt.Printf("%-24s %d\t %s\n", sa.Name, len(sa.Secrets), sa.CreationTimestamp)
	}
}
