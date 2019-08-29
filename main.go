package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/alandtsang/kc/clusters"
	"github.com/alandtsang/kc/resources"
)

// 实际中应该用更好的变量名
var (
	h bool

	//v, V bool
	//t, T bool
	//q    *bool

	//s string
	//p string
	//c string
	//g string

	n string
	g bool
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&n, "n", "", "namespace")
	flag.BoolVar(&g, "get", false, "Display one or many resources")

	//flag.BoolVar(&v, "v", false, "show version and exit")
	//flag.BoolVar(&V, "V", false, "show version and configure options then exit")

	//flag.BoolVar(&t, "t", false, "test configuration and exit")
	//flag.BoolVar(&T, "T", false, "test configuration, dump it and exit")

	// 另一种绑定方式
	//q = flag.Bool("q", false, "suppress non-error messages during configuration testing")

	// 注意 `signal`。默认是 -s string，有了 `signal` 之后，变为 -s signal
	//flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
	//flag.StringVar(&p, "p", "/usr/local/nginx/", "set `prefix` path")
	//flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")
	//flag.StringVar(&g, "g", "conf/nginx.conf", "set global `directives` out of configuration file")

	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `kc controls the Kubernetes cluster manager.

Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
	}

	if h {
		flag.Usage()
	}

	var cs clusters.ClusterSet
	cs.Init()
	//cs.Print()
	clusterName := "k8s-test2"
	cluster := cs.GetCluster(clusterName)
	if cluster == nil {
		log.Fatalf("[Error] Cluster name %s not found, please enter the correct name\n", clusterName)
	}
	//fmt.Printf("%+v\n", *cluster)
	clientset := GetClient(newConfig(cluster.Addr, cluster.Token))
	//resources.GetNamespaces(clientset)
	//resources.GetPods(clientset)
	//resources.GetService(clientset)
	//resources.GetConfigMaps(clientset)
	//resources.GetNodes(clientset)
	//resources.GetEndPoints(clientset)
	//resources.GetSecrets(clientset)
	resources.GetServiceAccounts(clientset)
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

func GetClient(config *rest.Config) *kubernetes.Clientset {
	// 根据指定的 config 创建一个新的 clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}
