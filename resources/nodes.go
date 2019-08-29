package resources

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Nodes struct {
}

func GetNodes(clientset *kubernetes.Clientset) {
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("NAME\t\t\t STATUS\t ROLES\t AGE\t VERSION\n")
	for _, node := range nodes.Items {
		fmt.Printf("%-24s %s\t %s\t %s\t %s\n", node.Name, "Ready", "None",
			node.CreationTimestamp, node.Status.NodeInfo.KubeletVersion)
	}
}
