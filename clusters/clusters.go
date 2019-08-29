package clusters

import "fmt"

type Cluster struct {
	Name  string
	Addr  string
	Token string
}

// Fill in your cluster name, address and token here
var clusters = []Cluster{
	{"k8s-test1", "10.20.30.40:6443", "xxxxxx"},
	{"k8s-test2", "10.20.30.41:6443", "xxxxxx"},
}

type ClusterSet struct {
	c map[string]*Cluster
}

func NewClusterSet() *ClusterSet {
	var cs ClusterSet
	cs.c = make(map[string]*Cluster)
	cs.Update()
	return &cs
}

func (cs *ClusterSet) Update() {
	for i, r := range clusters {
		cs.c[r.Name] = &clusters[i]
	}
}

func (cs *ClusterSet) Print() {
	for k, v := range cs.c {
		fmt.Println(k, v)
	}
}

func (cs *ClusterSet) GetCluster(name string) *Cluster {
	if resource, ok := cs.c[name]; ok {
		return resource
	}
	return nil
}
