/*
Open Source Initiative OSI - The MIT License (MIT):Licensing

The MIT License (MIT)
Copyright (c) 2013 - 2022 Ralph Caraveo (deckarep@gmail.com)

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package mapset

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type yourType struct {
	name string
}

func (i yourType) Equal(jAny any) bool {
	j, ok := jAny.(yourType)
	if !ok {
		return false
	}

	return i == j
}

func (i yourType) Key() string {
	return i.name
}

func Test_ExampleIterator(t *testing.T) {
	r := require.New(t)
	s := NewSet(
		[]*yourType{
			{name: "Alise"},
			{name: "Bob"},
			{name: "John"},
			{name: "Nick"},
		}...,
	)

	var found *yourType
	it := s.Iterator()

	for elem := range it.C {
		if elem.name == "John" {
			found = elem
			it.Stop()
		}
	}

	r.NotNil(found)
	r.Equal("John", found.name, "expected iterator to have found `John` record but got nil or something else")
}
