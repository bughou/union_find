package union_find

import "fmt"

type UnionFind map[interface{}]map[interface{}]struct{}

func New() UnionFind {
	return make(UnionFind)
}

func (uf UnionFind) Union(objects ...interface{}) {
}

func (uf UnionFind) Find(object interface{}) map[interface{}]struct{} {
	return uf[object]
}
