package union_find

import (
	listPkg "container/list"
)

type UnionFind map[interface{}]*listPkg.List

func New() UnionFind {
	return make(UnionFind)
}

// union objects into the same set, including the existing set members of objects
func (uf UnionFind) Union(objects ...interface{}) {
	for i := 1; i < len(objects); i++ {
		uf.union(objects[i-1], objects[i])
	}
}

func (uf UnionFind) union(a, b interface{}) {
	if a == b {
		return
	}
	la := uf[a]
	lb := uf[b]
	switch {
	case la == nil && lb == nil:
		list := listPkg.New()
		list.PushBack(a)
		list.PushBack(b)
		uf[a] = list
		uf[b] = list
	case la != nil && lb == nil:
		la.PushBack(b)
		uf[b] = la
	case lb != nil && la == nil:
		lb.PushBack(a)
		uf[a] = lb
	case la != lb:
	default:
		// la == lb, so a and b is aready in the same sets.
	}
}

// check if a and b are in the same set
func (uf UnionFind) InSameSet(a, b interface{}) bool {
	la := uf[a]
	if la == nil {
		return false
	}
	lb := uf[b]
	if lb == nil {
		return false
	}
	return la == lb
}

// find all set members of object
func (uf UnionFind) Find(object interface{}) (result []interface{}) {
	list := uf[object]
	if list == nil {
		return nil
	}
	for e := list.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value)
	}
	return
}
