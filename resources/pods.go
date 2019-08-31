package resources

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Pods struct {
	clientSet *kubernetes.Clientset
	namespace string
	name      string
}

func NewPods(clientSet *kubernetes.Clientset, namespace, name string) *Pods {
	return &Pods{clientSet: clientSet, namespace: namespace, name: name}
}

func (p *Pods) Get() {
	if len(p.name) > 0 {
		p.getPod()
	} else {
		p.getPodList()
	}
}

func (p *Pods) Delete() {

}

func (p *Pods) GetLogs() {
	req := p.clientSet.CoreV1().Pods(p.namespace).GetLogs(p.name, &v1.PodLogOptions{})
	podLogs, err := req.Stream()
	if err != nil {
		log.Fatalln("[Error] Logs()", err.Error())
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		log.Fatalln("[Error] Logs()", err.Error())
	}
	str := buf.String()
	fmt.Println(str)
}

func (p *Pods) getPod() {
	pod, err := p.clientSet.CoreV1().Pods(p.namespace).Get(p.name, metav1.GetOptions{})
	if err != nil {
		log.Fatalln("[Error] GetPod()", err.Error())
	}
	p.printPod(pod)
}

func (p *Pods) getPodList() {
	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PodsGetter 接口方法 Pods(namespace string)返回 PodInterface
	// PodInterface 接口拥有操作 Pod 资源的方法，例如 Create、Update、Get、List 等方法

	// 注意：Pods() 方法中 namespace 不指定则获取 Cluster 所有 namespace 的 Pod 列表
	// 若指定 namespace 则获取指定 Pod 列表信息
	pods, err := p.clientSet.CoreV1().Pods(p.namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("[Error] GetPodList()", err.Error())
	}
	p.printPodList(pods)
}

func (p *Pods) printPod(pod *v1.Pod) {
	fmt.Printf("NAME\t\t\t\t READY\t\t STATUS\t\t RESTARTS\t AGE\n")
	fmt.Printf("%-30s\t %s\t\t %s\t %-d\t\t %s\n", pod.ObjectMeta.Name, "1/1", getPodStatus(pod), 0, pod.ObjectMeta.CreationTimestamp)

}

func (p *Pods) printPodList(pods *v1.PodList) {
	if len(pods.Items) > 0 {
		fmt.Printf("NAME\t\t\t\t READY\t\t STATUS\t\t RESTARTS\t AGE\n")
		for _, pod := range pods.Items {
			fmt.Printf("%-30s\t %s\t\t %s\t %-d\t\t %s\n", pod.ObjectMeta.Name, "1/1", getPodStatus(&pod), 0, pod.ObjectMeta.CreationTimestamp)
		}
	} else {
		fmt.Println(ERR_NO_RESOURCE)
	}
}

func getPodStatus(pod *v1.Pod) string {
	if pod.Status.Phase == "Runnging" {
		for _, cond := range pod.Status.Conditions {
			if cond.Status != "True" {
				return cond.Reason
			}
		}
	}
	return string(pod.Status.Phase)
}
