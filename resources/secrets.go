package resources

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Secrets struct {
}

func GetSecrets(clientset *kubernetes.Clientset) {
	namespace := "default"
	secrets, err := clientset.CoreV1().Secrets(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("NAME\t\t\t TYPE\t\t\t\t\t DATA\t AGE\n")
	for _, secret := range secrets.Items {
		fmt.Printf("%-24s %s\t %d\t %s\n", secret.Name, secret.Type, len(secret.Data), secret.CreationTimestamp)
	}
}
