package resources

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Secrets struct {
	clientSet *kubernetes.Clientset
	namespace string
	name      string
}

func NewSecrets(clientSet *kubernetes.Clientset, namespace, name string) *Secrets {
	return &Secrets{clientSet: clientSet, namespace: namespace, name: name}
}

func (s *Secrets) Get() {
	secrets, err := s.clientSet.CoreV1().Secrets(s.namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	if len(secrets.Items) > 0 {
		fmt.Printf("NAME\t\t\t TYPE\t\t\t\t\t DATA\t AGE\n")
		for _, secret := range secrets.Items {
			fmt.Printf("%-24s %s\t %d\t %s\n", secret.Name, secret.Type, len(secret.Data), secret.CreationTimestamp)
		}
	} else {
		fmt.Println(ERR_NO_RESOURCE)
	}
}
