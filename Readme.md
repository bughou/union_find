# union\_find
a union find sets data structure based on map and list in golang.

[![Build Status](https://travis-ci.org/lovego/union_find.svg?branch=master)](https://travis-ci.org/lovego/union_find)
[![Go Report Card](https://goreportcard.com/badge/github.com/lovego/union_find)](https://goreportcard.com/report/github.com/lovego/union_find)
[![Coverage Status](https://coveralls.io/repos/github/lovego/union_find/badge.svg?branch=master)](https://coveralls.io/github/lovego/union_find?branch=master)
[![GoDoc](https://godoc.org/github.com/lovego/union_find?status.svg)](https://godoc.org/github.com/lovego/union_find)

## Install
`$ go get github.com/lovego/union_find`

## Usage
```go
uf := union_find.New()

uf.Union(1, 2, 3)
fmt.Println(uf.Find(1)) // Prints [1, 2, 3]
fmt.Println(uf.Find(2)) // Prints [1, 2, 3]
fmt.Println(uf.Find(3)) // Prints [1, 2, 3]
fmt.Println(uf.InSameSet(1, 2)) // Prints true
fmt.Println(uf.InSameSet(1, 3)) // Prints true
fmt.Println(uf.InSameSet(2, 3)) // Prints true

uf.Union(4, 5, 6)
fmt.Println(uf.Find(4)) // Prints [4, 5, 6]
fmt.Println(uf.Find(5)) // Prints [4, 5, 6]
fmt.Println(uf.Find(6)) // Prints [4, 5, 6]
fmt.Println(uf.InSameSet(4, 5)) // Prints true
fmt.Println(uf.InSameSet(4, 6)) // Prints true
fmt.Println(uf.InSameSet(5, 6)) // Prints true

fmt.Println(uf.InSameSet(1, 4)) // Prints false
fmt.Println(uf.InSameSet(1, 5)) // Prints false
fmt.Println(uf.InSameSet(2, 6)) // Prints false

uf.Union(2, 6)
fmt.Println(uf.Find(1)) // Prints [1, 2, 3, 4, 5, 6]
fmt.Println(uf.Find(2)) // Prints [1, 2, 3, 4, 5, 6]
fmt.Println(uf.Find(3)) // Prints [1, 2, 3, 4, 5, 6]
fmt.Println(uf.Find(4)) // Prints [1, 2, 3, 4, 5, 6]
fmt.Println(uf.Find(5)) // Prints [1, 2, 3, 4, 5, 6]
fmt.Println(uf.Find(6)) // Prints [1, 2, 3, 4, 5, 6]

fmt.Println(uf.InSameSet(1, 4)) // Prints true
fmt.Println(uf.InSameSet(1, 5)) // Prints true
fmt.Println(uf.InSameSet(2, 6)) // Prints true
```

## Documentation
[https://godoc.org/github.com/lovego/union_find](https://godoc.org/github.com/lovego/union_find)
