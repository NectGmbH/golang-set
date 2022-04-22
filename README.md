![example workflow](https://github.com/deckarep/golang-set/actions/workflows/ci.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/deckarep/golang-set)](https://goreportcard.com/report/github.com/deckarep/golang-set)
[![GoDoc](https://godoc.org/github.com/deckarep/golang-set?status.svg)](http://godoc.org/github.com/deckarep/golang-set)

# golang-set

This is a fork of the excellent [deckarep/golang-set](https://github.com/deckarep/golang-set) package with a relaxed type constraint.

The missing `generic` set collection for the Go language.  Until Go has sets built-in...use this.

## Update 4/22/2022
* Packaged version: `3.0.0` release for struct generics support with breaking changes.
* supports `new generic` syntax
* Go `1.18.0` or higher

## Update 3/26/2022
* Packaged version: `2.0.0` release for generics support with breaking changes.
* supports `new generic` syntax
* Go `1.18.0` or higher

![With Generics](new_improved.jpeg)

## Features

* *NEW* [Generics](https://go.dev/doc/tutorial/generics) based implementation (requires [Go 1.18](https://go.dev/blog/go1.18beta1) or higher)
* One common *interface* to both implementations
  * a **non threadsafe** implementation favoring *performance*
  * a **threadsafe** implementation favoring *concurrent* use
* Feature complete set implementation modeled after [Python's set implementation](https://docs.python.org/3/library/stdtypes.html#set).
* Exhaustive unit-test and benchmark suite

## Usage

The code below demonstrates how a Set collection can better manage data and actually minimize boilerplate and needless loops in code. This package now fully supports *generic* syntax so you are now able to instantiate a collection for any Keyable type object (implementing this package's EqualKeyer interface).

The Key returned on an object's Key method must be stable (not change) and unique.
A good example are uuids for structs.
As a caveat to the extension to EqualKeyer interface from the comparable type constraint
is that on changes to a value in the set the key returned by the value's Key() function
must also change to have a true set implementation.

Using this library is as simple as creating either a threadsafe or non-threadsafe set and providing a `EqualKeyer` type for instantiation of the collection.

```go

import (
    "github.com/google/uuid"
)

type StructT struct {
    id uuid.UUID
}

func (i StructT) Equal(jAny any) bool {
	j, ok := jAny.(StructT)
	if !ok {
		return false
	}

	return i == j
}

func (t StructT) Key() string {
    return t.id.String()
}

// Syntax example, doesn't compile.
mySet := mapset.NewSet[T]() // where T is some concrete comparable type.

// Therefore this code creates an StructT set
mySet := mapset.NewSet[StructT]()

// Or perhaps you want a string set
// Wrap string in a EqualKeyer type

type String string

func (i String) Equal(jAny any) bool {
	j, ok := jAny.(String)
	if !ok {
		return false
	}

	return i == j
}

func (s String) Key() string {
    return string(s)
}

mySet := mapset.NewSet[String]()
```
