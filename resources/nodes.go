package resources

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Nodes struct {
	clientSet *kubernetes.Clientset
}

func NewNodes(clientSet *kubernetes.Clientset) *Nodes {
	return &Nodes{clientSet: clientSet}
}

func (n *Nodes) Get() {
	nodes, err := n.clientSet.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	if len(nodes.Items) > 0 {
		fmt.Printf("NAME\t\t\t STATUS\t ROLES\t AGE\t VERSION\n")
		for _, node := range nodes.Items {
			fmt.Printf("%-24s %s\t %s\t %s\t %s\n", node.Name, "Ready", "None",
				node.CreationTimestamp, node.Status.NodeInfo.KubeletVersion)
		}
	} else {
		fmt.Println(ERR_NO_RESOURCE)
	}
}
