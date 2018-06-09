package union_find

import (
	"reflect"
	"testing"
)

func TestInts1(t *testing.T) {
	uf := New()
	testUnionFind(t, uf, 1, 2, 3, 4)
	testUnionFind(t, uf, 5, 6, 7)
	testUnionFind(t, uf, 8, 9, 10)
	uf.Union(3, 6, 9)
	checkUnionFind(t, uf, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

func TestInts2(t *testing.T) {
	uf := New()
	testUnionFind(t, uf, 5, 6, 7)
	testUnionFind(t, uf, 8, 9, 10)
	testUnionFind(t, uf, 1, 2, 3, 4)
	uf.Union(3, 6, 9)
	checkUnionFind(t, uf, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

func TestInts3(t *testing.T) {
	uf := New()
	testUnionFind(t, uf, 5, 6, 7)
	testUnionFind(t, uf, 8, 9, 10)
	testUnionFind(t, uf, 1, 2, 3, 4)
	uf.Union(6, 9, 3)
	checkUnionFind(t, uf, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4)
}

func TestIntsDuplicae(t *testing.T) {
	uf := New()
	uf.Union(5, 6, 7, 7, 7)
	uf.Union(8, 8, 8, 9, 10)
	uf.Union(1, 2, 1, 3, 3, 4, 3)
	uf.Union(6, 6, 3, 9)
	checkUnionFind(t, uf, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

func testUnionFind(t *testing.T, uf UnionFind, values ...interface{}) {
	uf.Union(values...)
	checkUnionFind(t, uf, values...)
}

func checkUnionFind(t *testing.T, uf UnionFind, values ...interface{}) {
	for _, v := range values {
		got := uf.Find(v)
		if !reflect.DeepEqual(got, values) {
			t.Fatalf("unexpectd Find(%v): %v", v, got)
		}
	}
	for i := 1; i < len(values); i++ {
		if !uf.InSameSet(values[i-1], values[i]) {
			t.Fatalf("unexpectd InSameSet(%v, %v): false", values[i-1], values[i])
		}
	}
}
