package resources

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetEndPoints(clientset *kubernetes.Clientset) {
	namespace := "default"
	endpoints, err := clientset.CoreV1().Endpoints(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("NAME\t\t\t ENDPOINTS\t AGE\n")
	for _, endpoint := range endpoints.Items {
		fmt.Printf("%-24s %v\t %s\n", endpoint.Name, getEndPointsIPs(endpoint.Subsets), endpoint.CreationTimestamp)
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
