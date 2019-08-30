package resources

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type EndPoints struct {
	clientSet *kubernetes.Clientset
	namespace string
	name      string
}

func NewEndPoints(clientSet *kubernetes.Clientset, namespace, name string) *EndPoints {
	return &EndPoints{clientSet: clientSet, namespace: namespace, name: name}
}

func (ep *EndPoints) Get() {
	endpoints, err := ep.clientSet.CoreV1().Endpoints(ep.namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	if len(endpoints.Items) > 0 {
		fmt.Printf("NAME\t\t\t ENDPOINTS\t AGE\n")
		for _, endpoint := range endpoints.Items {
			fmt.Printf("%-24s %v\t %s\n", endpoint.Name, getEndPointsIPs(endpoint.Subsets), endpoint.CreationTimestamp)
		}
	} else {
		fmt.Println(ERR_NO_RESOURCE)
	}
}

func getEndPointsIPs(subsets []v1.EndpointSubset) string {
	var endpoints string
	for _, subset := range subsets {
		for _, address := range subset.Addresses {
			endpoints = endpoints + address.IP + ","
		}
	}
	return endpoints
}
