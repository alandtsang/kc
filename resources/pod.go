package resources

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetPod(clientset *kubernetes.Clientset) {
	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PodsGetter 接口方法 Pods(namespace string)返回 PodInterface
	// PodInterface 接口拥有操作 Pod 资源的方法，例如 Create、Update、Get、List 等方法
	// 注意：Pods() 方法中 namespace 不指定则获取 Cluster 所有 Pod 列表
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	//fmt.Printf("There are %d pods in the k8s cluster\n", len(pods.Items))

	// 获取指定 namespace 中的 Pod 列表信息
	namespace := "default"
	pods, err = clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("NAME\t\t\t\t READY\t\t STATUS\t\t RESTARTS\t AGE\n")
	for _, pod := range pods.Items {
		fmt.Printf("%-30s\t %s\t\t %s\t %-d\t\t %s\n", pod.ObjectMeta.Name, "1/1", getPodStatus(pod), 0, pod.ObjectMeta.CreationTimestamp)
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
