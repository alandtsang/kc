package resources

import (
	"fmt"
	"log"

	"k8s.io/api/core/v1"
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
	if len(s.name) > 0 {
		s.getSecret()
	} else {
		s.getSecretList()
	}
}

func (s *Secrets) Delete() {
	err := s.clientSet.CoreV1().Secrets(s.namespace).Delete(s.name, &metav1.DeleteOptions{})
	if err != nil {
		log.Fatalf("[Error] Delete secret %s failed: %s\n", s.name, err.Error())
	} else {
		fmt.Printf("secret \"%s\" deleted\n", s.name)
	}
}

func (s *Secrets) getSecret() {
	secret, err := s.clientSet.CoreV1().Secrets(s.namespace).Get(s.name, metav1.GetOptions{})
	if err != nil {
		log.Fatalln("[Error] GetSecret()", err.Error())
	}
	s.printSecret(secret)
}

func (s *Secrets) getSecretList() {
	secrets, err := s.clientSet.CoreV1().Secrets(s.namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("[Error] GetSecretList()", err.Error())
	}
	s.printSecretList(secrets)
}

func (s *Secrets) printSecret(secret *v1.Secret) {
	fmt.Printf("NAME\t\t\t TYPE\t\t\t\t\t DATA\t AGE\n")
	fmt.Printf("%-24s %s\t %d\t %s\n", secret.Name, secret.Type, len(secret.Data), secret.CreationTimestamp)
}

func (s *Secrets) printSecretList(secrets *v1.SecretList) {
	if len(secrets.Items) > 0 {
		fmt.Printf("NAME\t\t\t TYPE\t\t\t\t\t DATA\t AGE\n")
		for _, secret := range secrets.Items {
			fmt.Printf("%-24s %s\t %d\t %s\n", secret.Name, secret.Type, len(secret.Data), secret.CreationTimestamp)
		}
	} else {
		fmt.Println(ERR_NO_RESOURCE)
	}
}
