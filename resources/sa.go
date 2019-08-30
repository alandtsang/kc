package resources

import (
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
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
	if len(sa.name) > 0 {
		sa.GetServiceAccount()
	} else {
		sa.GetServiceAccountList()
	}
}

func (sa *ServiceAccounts) GetServiceAccount() {
	serviceAccount, err := sa.clientSet.CoreV1().ServiceAccounts(sa.namespace).Get(sa.name, metav1.GetOptions{})
	if err != nil {
		log.Fatalln("[Error] GetServiceAccount()", err.Error())
	}
	sa.printServiceAccount(serviceAccount)
}

func (sa *ServiceAccounts) GetServiceAccountList() {
	serviceAccounts, err := sa.clientSet.CoreV1().ServiceAccounts(sa.namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("[Error] GetServiceAccountList()", err.Error())
	}
	sa.printServiceAccountList(serviceAccounts)
}

func (sa *ServiceAccounts) printServiceAccount(s *v1.ServiceAccount) {
	fmt.Printf("NAME\t\t\t SECRETS\t AGE\n")
	fmt.Printf("%-24s %d\t %s\n", s.Name, len(s.Secrets), s.CreationTimestamp)
}

func (sa *ServiceAccounts) printServiceAccountList(serviceAccounts *v1.ServiceAccountList) {
	if len(serviceAccounts.Items) > 0 {
		fmt.Printf("NAME\t\t\t SECRETS\t AGE\n")
		for _, s := range serviceAccounts.Items {
			fmt.Printf("%-24s %d\t %s\n", s.Name, len(s.Secrets), s.CreationTimestamp)
		}
	} else {
		fmt.Println(ERR_NO_RESOURCE)
	}
}
