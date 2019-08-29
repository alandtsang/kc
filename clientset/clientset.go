package clientset

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Client struct {
	ClientSet *kubernetes.Clientset
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

func NewClient(host, token string) *Client {
	config := newConfig(host, token)
	// 根据指定的 config 创建一个新的 clientset
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return &Client{ClientSet: clientSet}
}
