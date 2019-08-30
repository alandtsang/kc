package clientset

import (
	"log"

	"github.com/alandtsang/kc/clusters"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Client struct {
	ClientSet *kubernetes.Clientset
}

func (c *Client) Get() *kubernetes.Clientset {
	return c.ClientSet
}

func NewClient(clusterName string) *Client {
	cs := clusters.NewClusterSet()
	cluster := cs.GetCluster(clusterName)
	if cluster == nil {
		log.Fatalf("[Error] Cluster name %s not found, please enter the correct name\n", clusterName)
	}
	//fmt.Printf("%+v\n", *cluster)
	return newClient(cluster.Addr, cluster.Token)
}

func newClient(host, token string) *Client {
	config := newConfig(host, token)
	// 根据指定的 config 创建一个新的 clientset
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return &Client{ClientSet: clientSet}
}

func newConfig(host, token string) *rest.Config {
	// uses the current context in kubeconfig
	config := &rest.Config{
		Host:            host,
		BearerToken:     token,
		TLSClientConfig: rest.TLSClientConfig{Insecure: true},
	}
	return config
}
