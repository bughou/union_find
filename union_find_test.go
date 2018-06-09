package union_find

import (
	"reflect"
	"testing"

	"github.com/lovego/errs"
)

func TestInts1(t *testing.T) {
	uf := New()

	testUnionFind(t, uf, 1, 2, 3, 4)

	checkInvalid(t, uf, 5, 6, 7)
	testUnionFind(t, uf, 5, 6, 7)
	if uf.InSameSet(1, 5) {
		t.Fatalf("unexpectd InSameSet(1, 5): true\n%s", errs.Stack(2))
	}

	checkInvalid(t, uf, 8, 9, 10)
	testUnionFind(t, uf, 8, 9, 10)
	if uf.InSameSet(7, 9) {
		t.Fatalf("unexpectd InSameSet(7, 9): true\n%s", errs.Stack(2))
	}

	checkValid(t, uf, 1, 2, 3, 4)
	checkValid(t, uf, 5, 6, 7)
	checkValid(t, uf, 8, 9, 10)

	uf.Union(3, 6, 9)
	checkValid(t, uf, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	uf.Union(100, 1)
	checkValid(t, uf, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100)

	if uf.InSameSet(100, 101) {
		t.Fatalf("unexpectd InSameSet(%v, %v): true\n%s", errs.Stack(2))
	}
}

func TestInts2(t *testing.T) {
	uf := New()
	testUnionFind(t, uf, 5, 6, 7)
	testUnionFind(t, uf, 8, 9, 10)
	testUnionFind(t, uf, 1, 2, 3, 4)
	uf.Union(3, 6, 9)
	checkValid(t, uf, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

func TestInts3(t *testing.T) {
	uf := New()
	testUnionFind(t, uf, 5, 6, 7)
	testUnionFind(t, uf, 8, 9, 10)
	testUnionFind(t, uf, 1, 2, 3, 4)
	uf.Union(6, 9, 3)
	checkValid(t, uf, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4)
}

func TestIntsDuplicae(t *testing.T) {
	uf := New()
	uf.Union(5, 6, 7, 7, 7)
	uf.Union(8, 8, 8, 9, 10)
	uf.Union(1, 2, 1, 3, 3, 4, 3)
	uf.Union(6, 6, 3, 9)
	checkValid(t, uf, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

func testUnionFind(t *testing.T, uf UnionFind, values ...interface{}) {
	uf.Union(values...)
	checkValid(t, uf, values...)
}

func checkValid(t *testing.T, uf UnionFind, values ...interface{}) {
	for _, v := range values {
		got := uf.Find(v)
		if !reflect.DeepEqual(got, values) {
			t.Fatalf("unexpectd Find(%v): %v\n%s", v, got, errs.Stack(2))
		}
	}
	for i := 1; i < len(values); i++ {
		if !uf.InSameSet(values[i-1], values[i]) {
			t.Fatalf("unexpectd InSameSet(%v, %v): false\n%s", values[i-1], values[i],
				errs.Stack(2))
		}
	}
}

func checkInvalid(t *testing.T, uf UnionFind, values ...interface{}) {
	for _, v := range values {
		got := uf.Find(v)
		if got != nil {
			t.Fatalf("unexpectd Find(%v): %v\n%s", v, got, errs.Stack(2))
		}
	}
	for i := 1; i < len(values); i++ {
		if uf.InSameSet(values[i-1], values[i]) {
			t.Fatalf("unexpectd InSameSet(%v, %v): true\n%s", values[i-1], values[i],
				errs.Stack(2))
		}
	}
}
