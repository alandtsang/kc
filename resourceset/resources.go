// Resource Set

package resourceset

import "fmt"

type Resource struct {
	Name  string
	Addr  string
	Token string
}

// Fill in your cluster name, address and token here
var resources = []Resource{
	{"k8s-test1", "10.20.30.40:6443", "xxxxxx"},
	{"k8s-test2", "10.20.30.41:6443", "xxxxxx"},
}

type ResourceSet struct {
	r map[string]*Resource
}

func (rs *ResourceSet) Init() {
	rs.r = make(map[string]*Resource)
	rs.Update()
}

func (rs *ResourceSet) Update() {
	for i, r := range resources {
		rs.r[r.Name] = &resources[i]
	}
}

func (rs *ResourceSet) Print() {
	for k, v := range rs.r {
		fmt.Println(k, v)
	}
}

func (rs *ResourceSet) GetResource(name string) *Resource {
	if resource, ok := rs.r[name]; ok {
		return resource
	}
	return nil
}
