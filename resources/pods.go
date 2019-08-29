package resources

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Pods struct {
	clientSet *kubernetes.Clientset
	namespace string
}

func NewPods(clientSet *kubernetes.Clientset, namespace string) *Pods {
	return &Pods{clientSet: clientSet, namespace: namespace}
}

func (p *Pods) Get() {
	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PodsGetter 接口方法 Pods(namespace string)返回 PodInterface
	// PodInterface 接口拥有操作 Pod 资源的方法，例如 Create、Update、Get、List 等方法

	// 注意：Pods() 方法中 namespace 不指定则获取 Cluster 所有 namespace 的 Pod 列表
	// 若指定 namespace 则获取指定 Pod 列表信息
	pods, err := p.clientSet.CoreV1().Pods(p.namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	if len(pods.Items) > 0 {
		fmt.Printf("NAME\t\t\t\t READY\t\t STATUS\t\t RESTARTS\t AGE\n")
		for _, pod := range pods.Items {
			fmt.Printf("%-30s\t %s\t\t %s\t %-d\t\t %s\n", pod.ObjectMeta.Name, "1/1", getPodStatus(pod), 0, pod.ObjectMeta.CreationTimestamp)
		}
	} else {
		fmt.Println("No resources found.")
	}
}

func getPodStatus(pod v1.Pod) string {
	if pod.Status.Phase == "Runnging" {
		for _, cond := range pod.Status.Conditions {
			if cond.Status != "True" {
				return cond.Reason
			}
		}
	}
	return string(pod.Status.Phase)
}
