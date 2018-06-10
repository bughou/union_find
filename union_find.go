package union_find

import (
	listPkg "container/list"
)

// UnionFind data structure.
type UnionFind map[interface{}]*listPkg.List

// New return a new UnionFind instance.
func New() UnionFind {
	return make(UnionFind)
}

// Union objects into the same set, including the existing set members of objects
func (uf UnionFind) Union(objects ...interface{}) {
	for i := 1; i < len(objects); i++ {
		uf.union(objects[i-1], objects[i])
	}
}

func (uf UnionFind) union(a, b interface{}) {
	if a == b {
		return
	}
	la, lb := uf[a], uf[b]
	switch {
	case la == nil && lb == nil:
		uf.init(a, b)
	case la != nil && lb == nil:
		la.PushBack(b)
		uf[b] = la
	case lb != nil && la == nil:
		lb.PushBack(a)
		uf[a] = lb
	case la != lb:
		uf.concat(la, lb)
	default:
		// la == lb, so a and b is already in the same sets.
	}
}

func (uf UnionFind) init(a, b interface{}) {
	list := listPkg.New()
	list.PushBack(a)
	list.PushBack(b)
	uf[a] = list
	uf[b] = list
}

func (uf UnionFind) concat(la, lb *listPkg.List) {
	if la.Len() < lb.Len() {
		la, lb = lb, la
	}
	for b := lb.Front(); b != nil; b = b.Next() {
		la.PushBack(b.Value)
		uf[b.Value] = la
	}
}

// Find all set members of objects, with no duplicates.
func (uf UnionFind) Find(objects ...interface{}) (result []interface{}) {
	if len(objects) > 1 {
		objects = uf.Distinct(objects...)
	}
	for i := 0; i < len(objects); i++ {
		result = append(result, uf.find(objects[i])...)
	}
	return
}

func (uf UnionFind) find(object interface{}) (result []interface{}) {
	list := uf[object]
	if list == nil {
		return nil
	}
	for e := list.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value)
	}
	return
}

// Distinct removes the duplicate objects in any same set.
func (uf UnionFind) Distinct(objects ...interface{}) []interface{} {
	result := []interface{}{objects[0]}
	for i := 1; i < len(objects); i++ {
		if !uf.InSameSetSlice(objects[i], result) {
			result = append(result, objects[i])
		}
	}
	return result
}

// InSameSetSlice check if object is in the same set with any member of slice.
func (uf UnionFind) InSameSetSlice(object interface{}, slice []interface{}) bool {
	for _, e := range slice {
		if uf.InSameSet(object, e) {
			return true
		}
	}
	return false
}

// InSameSet check if a and b are in the same set
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
