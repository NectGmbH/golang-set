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
	"encoding/json"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

const N = 1000

func Test_AddConcurrent(t *testing.T) {
	r := require.New(t)
	runtime.GOMAXPROCS(2)

	s := NewSet[Int]()
	ints := rand.Perm(N)

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for i := 0; i < len(ints); i++ {
		go func(i Int) {
			s.Add(i)
			wg.Done()
		}(Int(i))
	}

	wg.Wait()
	for _, i := range ints {
		r.Truef(s.Contains(Int(i)), "Set is missing element: %v", i)
	}
}

func Test_CardinalityConcurrent(t *testing.T) {
	r := require.New(t)
	runtime.GOMAXPROCS(2)

	s := NewSet[Int]()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		elems := s.Cardinality()
		for i := 0; i < N; i++ {
			r.GreaterOrEqual(s.Cardinality(), elems, "cardinality shrunk")
		}
		wg.Done()
	}()

	for i := 0; i < N; i++ {
		s.Add(Int(rand.Int()))
	}
	wg.Wait()
}

func Test_ClearConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewSet[Int]()
	ints := rand.Perm(N)

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for i := 0; i < len(ints); i++ {
		go func() {
			s.Clear()
			wg.Done()
		}()
		go func(i Int) {
			s.Add(i)
		}(Int(i))
	}

	wg.Wait()
}

func Test_CloneConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewSet[Int]()
	ints := rand.Perm(N)

	for _, v := range ints {
		s.Add(Int(v))
	}

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for i := range ints {
		go func(i Int) {
			s.Remove(i)
			wg.Done()
		}(Int(i))
	}

	s.Clone()
}

func Test_ContainsConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewSet[Int]()
	ints := rand.Perm(N)
	integers := make([]Int, 0)
	for _, v := range ints {
		s.Add(Int(v))
		integers = append(integers, Int(v))
	}

	var wg sync.WaitGroup
	for range ints {
		wg.Add(1)
		go func() {
			s.Contains(integers...)
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_DifferenceConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s, ss := NewSet[Int](), NewSet[Int]()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(Int(v))
		ss.Add(Int(v))
	}

	var wg sync.WaitGroup
	for range ints {
		wg.Add(1)
		go func() {
			s.Difference(ss)
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_EqualConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s, ss := NewSet[Int](), NewSet[Int]()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(Int(v))
		ss.Add(Int(v))
	}

	var wg sync.WaitGroup
	for range ints {
		wg.Add(1)
		go func() {
			s.Equal(ss)
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_IntersectConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s, ss := NewSet[Int](), NewSet[Int]()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(Int(v))
		ss.Add(Int(v))
	}

	var wg sync.WaitGroup
	for range ints {
		wg.Add(1)
		go func() {
			s.Intersect(ss)
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_IsSubsetConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s, ss := NewSet[Int](), NewSet[Int]()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(Int(v))
		ss.Add(Int(v))
	}

	var wg sync.WaitGroup
	for range ints {
		wg.Add(1)
		go func() {
			s.IsSubset(ss)
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_IsProperSubsetConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s, ss := NewSet[Int](), NewSet[Int]()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(Int(v))
		ss.Add(Int(v))
	}

	var wg sync.WaitGroup
	for range ints {
		wg.Add(1)
		go func() {
			s.IsProperSubset(ss)
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_IsSupersetConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s, ss := NewSet[Int](), NewSet[Int]()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(Int(v))
		ss.Add(Int(v))
	}

	var wg sync.WaitGroup
	for range ints {
		wg.Add(1)
		go func() {
			s.IsSuperset(ss)
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_IsProperSupersetConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s, ss := NewSet[Int](), NewSet[Int]()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(Int(v))
		ss.Add(Int(v))
	}

	var wg sync.WaitGroup
	for range ints {
		wg.Add(1)
		go func() {
			s.IsProperSuperset(ss)
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_EachConcurrent(t *testing.T) {
	r := require.New(t)

	runtime.GOMAXPROCS(2)
	concurrent := 10

	s := NewSet[Int]()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(Int(v))
	}

	var count int64
	wg := new(sync.WaitGroup)
	wg.Add(concurrent)
	for n := 0; n < concurrent; n++ {
		go func() {
			defer wg.Done()
			s.Each(func(elem Int) bool {
				atomic.AddInt64(&count, 1)
				return false
			})
		}()
	}
	wg.Wait()

	r.Equal(count, int64(N*concurrent), "count mismatch")
}

func Test_IterConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewSet[Int]()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(Int(v))
	}

	cs := make([]<-chan Int, 0)
	for range ints {
		cs = append(cs, s.Iter())
	}

	c := make(chan interface{})
	go func() {
		for n := 0; n < len(ints)*N; {
			for _, d := range cs {
				select {
				case <-d:
					n++
					c <- nil
				default:
				}
			}
		}
		close(c)
	}()

	for range c {
	}
}

func Test_RemoveConcurrent(t *testing.T) {
	r := require.New(t)
	runtime.GOMAXPROCS(2)

	s := NewSet[Int]()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(Int(v))
	}

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for _, v := range ints {
		go func(i Int) {
			s.Remove(i)
			wg.Done()
		}(Int(v))
	}
	wg.Wait()

	r.Zero(s.Cardinality(), "cardinality not zero after removing all elems")
}

func Test_StringConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := NewSet[Int]()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(Int(v))
	}

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for range ints {
		go func() {
			_ = s.String()
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_SymmetricDifferenceConcurrent(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s, ss := NewSet[Int](), NewSet[Int]()
	ints := rand.Perm(N)
	for _, v := range ints {
		s.Add(Int(v))
		ss.Add(Int(v))
	}

	var wg sync.WaitGroup
	for range ints {
		wg.Add(1)
		go func() {
			s.SymmetricDifference(ss)
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_ToSlice(t *testing.T) {
	r := require.New(t)
	runtime.GOMAXPROCS(2)

	s := NewSet[Int]()
	ints := rand.Perm(N)

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for i := 0; i < len(ints); i++ {
		go func(i Int) {
			s.Add(i)
			wg.Done()
		}(Int(i))
	}

	wg.Wait()
	setAsSlice := s.ToSlice()
	r.Equal(len(setAsSlice), s.Cardinality(), "set length is incorrect")

	for _, i := range setAsSlice {
		r.Truef(s.Contains(i), "set is missing element: %+v", i)
	}
}

// Test_ToSliceDeadlock - fixes issue: https://github.com/deckarep/golang-set/issues/36
// This code reveals the deadlock however it doesn't happen consistently.
func Test_ToSliceDeadlock(t *testing.T) {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	set := NewSet[Int]()
	workers := 10
	wg.Add(workers)
	for i := 1; i <= workers; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				set.Add(1)
				set.ToSlice()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_UnmarshalJSON(t *testing.T) {
	r := require.New(t)
	s := []byte(`["test", "1", "2", "3"]`) //,["4,5,6"]]`)
	expected := NewSet(
		[]String{
			String(json.Number("1")),
			String(json.Number("2")),
			String(json.Number("3")),
			"test",
		}...,
	)

	actual := NewSet[String]()
	err := json.Unmarshal(s, actual)
	r.NoError(err)

	r.Truef(expected.Equal(actual), "Expected no difference, got: %v", expected.Difference(actual))
}

func Test_MarshalJSON(t *testing.T) {
	r := require.New(t)

	expected := NewSet(
		[]String{
			String(json.Number("1")),
			"test",
		}...,
	)

	b, err := json.Marshal(
		NewSet(
			[]String{
				"1",
				"test",
			}...,
		),
	)
	r.NoError(err, "marshaling to json")

	actual := NewSet[String]()
	err = json.Unmarshal(b, actual)
	r.NoError(err, "unmarshaling from json")

	r.Truef(expected.Equal(actual), "Expected no difference, got: %v", expected.Difference(actual))
}
