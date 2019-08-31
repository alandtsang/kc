package resources

import (
	"fmt"
	"log"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Services struct {
	clientSet *kubernetes.Clientset
	namespace string
	name      string
}

func NewServices(clientSet *kubernetes.Clientset, namespace, name string) *Services {
	return &Services{clientSet: clientSet, namespace: namespace, name: name}
}

func (s *Services) Get() {
	if len(s.name) > 0 {
		s.getService()
	} else {
		s.getServiceList()
	}
}

func (s *Services) Delete() {
	err := s.clientSet.CoreV1().Services(s.namespace).Delete(s.name, &metav1.DeleteOptions{})
	if err != nil {
		log.Fatalf("[Error] Delete service %s failed: %s\n", s.name, err.Error())
	} else {
		fmt.Printf("service \"%s\" deleted\n", s.name)
	}
}

func (s *Services) getService() {
	service, err := s.clientSet.CoreV1().Services(s.namespace).Get(s.name, metav1.GetOptions{})
	if err != nil {
		log.Fatalln("[Error] getService()", err.Error())
	}
	s.printService(service)
}

func (s *Services) getServiceList() {
	services, err := s.clientSet.CoreV1().Services(s.namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("[Error] getServiceList()", err.Error())
	}
	s.printServiceList(services)
}

func (s *Services) printService(service *v1.Service) {
	fmt.Printf("NAME\t\t TYPE\t\t CLUSTER-IP\t EXTERNAL-IP\t PORT(S)\t AGE\n")
	fmt.Printf("%-16s %s\t %s\t %v\t\t %v\t %s\n",
		service.Name, service.Spec.Type, service.Spec.ClusterIP,
		service.Spec.ExternalIPs, service.Spec.Ports, service.CreationTimestamp)
}

func (s *Services) printServiceList(services *v1.ServiceList) {
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
