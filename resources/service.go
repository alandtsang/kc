package resources

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Services struct {
	clientSet *kubernetes.Clientset
	namespace string
}

func NewServices(clientSet *kubernetes.Clientset, namespace string) *Services {
	return &Services{clientSet: clientSet, namespace: namespace}
}

func (s *Services) Get() {
	services, err := s.clientSet.CoreV1().Services(s.namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	if len(services.Items) > 0 {
		fmt.Printf("NAME\t\t TYPE\t\t CLUSTER-IP\t EXTERNAL-IP\t PORT(S)\t AGE\n")
		for _, service := range services.Items {
			fmt.Printf("%-16s %s\t %s\t %v\t\t %v\t %s\n",
				service.Name, service.Spec.Type, service.Spec.ClusterIP,
				service.Spec.ExternalIPs, service.Spec.Ports, service.CreationTimestamp)
		}
	} else {
		fmt.Println(ERR_NO_RESOURCE)
	}
}
