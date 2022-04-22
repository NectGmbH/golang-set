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
	"math/rand"
	"testing"
)

func nrand(n Int) []Int {
	i := make([]Int, n)
	for ind := range i {
		i[ind] = Int(rand.Int())
	}
	return i
}

func benchAdd(b *testing.B, n Int, newSet func(...Int) Set[Int]) {
	nums := nrand(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := newSet()
		for _, v := range nums {
			s.Add(Int(v))
		}
	}
}

func BenchmarkAddSafe(b *testing.B) {
	benchAdd(b, 1000, NewSet[Int])
}

func BenchmarkAddUnsafe(b *testing.B) {
	benchAdd(b, 1000, NewThreadUnsafeSet[Int])
}

func benchRemove(b *testing.B, s Set[Int]) {
	nums := nrand(Int(b.N))
	for _, v := range nums {
		s.Add(Int(v))
	}

	b.ResetTimer()
	for _, v := range nums {
		s.Remove(Int(v))
	}
}

func BenchmarkRemoveSafe(b *testing.B) {
	benchRemove(b, NewSet[Int]())
}

func BenchmarkRemoveUnsafe(b *testing.B) {
	benchRemove(b, NewThreadUnsafeSet[Int]())
}

func benchCardinality(b *testing.B, s Set[Int]) {
	for i := 0; i < b.N; i++ {
		s.Cardinality()
	}
}

func BenchmarkCardinalitySafe(b *testing.B) {
	benchCardinality(b, NewSet[Int]())
}

func BenchmarkCardinalityUnsafe(b *testing.B) {
	benchCardinality(b, NewThreadUnsafeSet[Int]())
}

func benchClear(b *testing.B, s Set[Int]) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Clear()
	}
}

func BenchmarkClearSafe(b *testing.B) {
	benchClear(b, NewSet[Int]())
}

func BenchmarkClearUnsafe(b *testing.B) {
	benchClear(b, NewThreadUnsafeSet[Int]())
}

func benchClone(b *testing.B, n Int, s Set[Int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(Int(v))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Clone()
	}
}

func BenchmarkClone1Safe(b *testing.B) {
	benchClone(b, 1, NewSet[Int]())
}

func BenchmarkClone1Unsafe(b *testing.B) {
	benchClone(b, 1, NewThreadUnsafeSet[Int]())
}

func BenchmarkClone10Safe(b *testing.B) {
	benchClone(b, 10, NewSet[Int]())
}

func BenchmarkClone10Unsafe(b *testing.B) {
	benchClone(b, 10, NewThreadUnsafeSet[Int]())
}

func BenchmarkClone100Safe(b *testing.B) {
	benchClone(b, 100, NewSet[Int]())
}

func BenchmarkClone100Unsafe(b *testing.B) {
	benchClone(b, 100, NewThreadUnsafeSet[Int]())
}

func benchContains(b *testing.B, n Int, s Set[Int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(Int(v))
	}

	nums[n-1] = -1 // Definitely not in s

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Contains(nums...)
	}
}

func BenchmarkContains1Safe(b *testing.B) {
	benchContains(b, 1, NewSet[Int]())
}

func BenchmarkContains1Unsafe(b *testing.B) {
	benchContains(b, 1, NewThreadUnsafeSet[Int]())
}

func BenchmarkContains10Safe(b *testing.B) {
	benchContains(b, 10, NewSet[Int]())
}

func BenchmarkContains10Unsafe(b *testing.B) {
	benchContains(b, 10, NewThreadUnsafeSet[Int]())
}

func BenchmarkContains100Safe(b *testing.B) {
	benchContains(b, 100, NewSet[Int]())
}

func BenchmarkContains100Unsafe(b *testing.B) {
	benchContains(b, 100, NewThreadUnsafeSet[Int]())
}

func benchEqual(b *testing.B, n Int, s, t Set[Int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(Int(v))
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Equal(t)
	}
}

func BenchmarkEqual1Safe(b *testing.B) {
	benchEqual(b, 1, NewSet[Int](), NewSet[Int]())
}

func BenchmarkEqual1Unsafe(b *testing.B) {
	benchEqual(b, 1, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkEqual10Safe(b *testing.B) {
	benchEqual(b, 10, NewSet[Int](), NewSet[Int]())
}

func BenchmarkEqual10Unsafe(b *testing.B) {
	benchEqual(b, 10, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkEqual100Safe(b *testing.B) {
	benchEqual(b, 100, NewSet[Int](), NewSet[Int]())
}

func BenchmarkEqual100Unsafe(b *testing.B) {
	benchEqual(b, 100, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func benchDifference(b *testing.B, n Int, s, t Set[Int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(Int(v))
	}
	for _, v := range nums[:n/2] {
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Difference(t)
	}
}

func benchIsSubset(b *testing.B, n Int, s, t Set[Int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(Int(v))
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.IsSubset(t)
	}
}

func BenchmarkIsSubset1Safe(b *testing.B) {
	benchIsSubset(b, 1, NewSet[Int](), NewSet[Int]())
}

func BenchmarkIsSubset1Unsafe(b *testing.B) {
	benchIsSubset(b, 1, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkIsSubset10Safe(b *testing.B) {
	benchIsSubset(b, 10, NewSet[Int](), NewSet[Int]())
}

func BenchmarkIsSubset10Unsafe(b *testing.B) {
	benchIsSubset(b, 10, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkIsSubset100Safe(b *testing.B) {
	benchIsSubset(b, 100, NewSet[Int](), NewSet[Int]())
}

func BenchmarkIsSubset100Unsafe(b *testing.B) {
	benchIsSubset(b, 100, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func benchIsSuperset(b *testing.B, n Int, s, t Set[Int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(Int(v))
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.IsSuperset(t)
	}
}

func BenchmarkIsSuperset1Safe(b *testing.B) {
	benchIsSuperset(b, 1, NewSet[Int](), NewSet[Int]())
}

func BenchmarkIsSuperset1Unsafe(b *testing.B) {
	benchIsSuperset(b, 1, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkIsSuperset10Safe(b *testing.B) {
	benchIsSuperset(b, 10, NewSet[Int](), NewSet[Int]())
}

func BenchmarkIsSuperset10Unsafe(b *testing.B) {
	benchIsSuperset(b, 10, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkIsSuperset100Safe(b *testing.B) {
	benchIsSuperset(b, 100, NewSet[Int](), NewSet[Int]())
}

func BenchmarkIsSuperset100Unsafe(b *testing.B) {
	benchIsSuperset(b, 100, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func benchIsProperSubset(b *testing.B, n Int, s, t Set[Int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(Int(v))
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.IsProperSubset(t)
	}
}

func BenchmarkIsProperSubset1Safe(b *testing.B) {
	benchIsProperSubset(b, 1, NewSet[Int](), NewSet[Int]())
}

func BenchmarkIsProperSubset1Unsafe(b *testing.B) {
	benchIsProperSubset(b, 1, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkIsProperSubset10Safe(b *testing.B) {
	benchIsProperSubset(b, 10, NewSet[Int](), NewSet[Int]())
}

func BenchmarkIsProperSubset10Unsafe(b *testing.B) {
	benchIsProperSubset(b, 10, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkIsProperSubset100Safe(b *testing.B) {
	benchIsProperSubset(b, 100, NewSet[Int](), NewSet[Int]())
}

func BenchmarkIsProperSubset100Unsafe(b *testing.B) {
	benchIsProperSubset(b, 100, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func benchIsProperSuperset(b *testing.B, n Int, s, t Set[Int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(Int(v))
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.IsProperSuperset(t)
	}
}

func BenchmarkIsProperSuperset1Safe(b *testing.B) {
	benchIsProperSuperset(b, 1, NewSet[Int](), NewSet[Int]())
}

func BenchmarkIsProperSuperset1Unsafe(b *testing.B) {
	benchIsProperSuperset(b, 1, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkIsProperSuperset10Safe(b *testing.B) {
	benchIsProperSuperset(b, 10, NewSet[Int](), NewSet[Int]())
}

func BenchmarkIsProperSuperset10Unsafe(b *testing.B) {
	benchIsProperSuperset(b, 10, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkIsProperSuperset100Safe(b *testing.B) {
	benchIsProperSuperset(b, 100, NewSet[Int](), NewSet[Int]())
}

func BenchmarkIsProperSuperset100Unsafe(b *testing.B) {
	benchIsProperSuperset(b, 100, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkDifference1Safe(b *testing.B) {
	benchDifference(b, 1, NewSet[Int](), NewSet[Int]())
}

func BenchmarkDifference1Unsafe(b *testing.B) {
	benchDifference(b, 1, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkDifference10Safe(b *testing.B) {
	benchDifference(b, 10, NewSet[Int](), NewSet[Int]())
}

func BenchmarkDifference10Unsafe(b *testing.B) {
	benchDifference(b, 10, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkDifference100Safe(b *testing.B) {
	benchDifference(b, 100, NewSet[Int](), NewSet[Int]())
}

func BenchmarkDifference100Unsafe(b *testing.B) {
	benchDifference(b, 100, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func benchIntersect(b *testing.B, n Int, s, t Set[Int]) {
	nums := nrand(Int(float64(n) * float64(1.5)))
	for _, v := range nums[:n] {
		s.Add(Int(v))
	}
	for _, v := range nums[n/2:] {
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Intersect(t)
	}
}

func BenchmarkIntersect1Safe(b *testing.B) {
	benchIntersect(b, 1, NewSet[Int](), NewSet[Int]())
}

func BenchmarkIntersect1Unsafe(b *testing.B) {
	benchIntersect(b, 1, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkIntersect10Safe(b *testing.B) {
	benchIntersect(b, 10, NewSet[Int](), NewSet[Int]())
}

func BenchmarkIntersect10Unsafe(b *testing.B) {
	benchIntersect(b, 10, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkIntersect100Safe(b *testing.B) {
	benchIntersect(b, 100, NewSet[Int](), NewSet[Int]())
}

func BenchmarkIntersect100Unsafe(b *testing.B) {
	benchIntersect(b, 100, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func benchSymmetricDifference(b *testing.B, n Int, s, t Set[Int]) {
	nums := nrand(Int(float64(n) * float64(1.5)))
	for _, v := range nums[:n] {
		s.Add(Int(v))
	}
	for _, v := range nums[n/2:] {
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.SymmetricDifference(t)
	}
}

func BenchmarkSymmetricDifference1Safe(b *testing.B) {
	benchSymmetricDifference(b, 1, NewSet[Int](), NewSet[Int]())
}

func BenchmarkSymmetricDifference1Unsafe(b *testing.B) {
	benchSymmetricDifference(b, 1, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkSymmetricDifference10Safe(b *testing.B) {
	benchSymmetricDifference(b, 10, NewSet[Int](), NewSet[Int]())
}

func BenchmarkSymmetricDifference10Unsafe(b *testing.B) {
	benchSymmetricDifference(b, 10, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkSymmetricDifference100Safe(b *testing.B) {
	benchSymmetricDifference(b, 100, NewSet[Int](), NewSet[Int]())
}

func BenchmarkSymmetricDifference100Unsafe(b *testing.B) {
	benchSymmetricDifference(b, 100, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func benchUnion(b *testing.B, n Int, s, t Set[Int]) {
	nums := nrand(n)
	for _, v := range nums[:n/2] {
		s.Add(Int(v))
	}
	for _, v := range nums[n/2:] {
		t.Add(v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Union(t)
	}
}

func BenchmarkUnion1Safe(b *testing.B) {
	benchUnion(b, 1, NewSet[Int](), NewSet[Int]())
}

func BenchmarkUnion1Unsafe(b *testing.B) {
	benchUnion(b, 1, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkUnion10Safe(b *testing.B) {
	benchUnion(b, 10, NewSet[Int](), NewSet[Int]())
}

func BenchmarkUnion10Unsafe(b *testing.B) {
	benchUnion(b, 10, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func BenchmarkUnion100Safe(b *testing.B) {
	benchUnion(b, 100, NewSet[Int](), NewSet[Int]())
}

func BenchmarkUnion100Unsafe(b *testing.B) {
	benchUnion(b, 100, NewThreadUnsafeSet[Int](), NewThreadUnsafeSet[Int]())
}

func benchEach(b *testing.B, n Int, s Set[Int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(Int(v))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Each(func(elem Int) bool {
			return false
		})
	}
}

func BenchmarkEach1Safe(b *testing.B) {
	benchEach(b, 1, NewSet[Int]())
}

func BenchmarkEach1Unsafe(b *testing.B) {
	benchEach(b, 1, NewThreadUnsafeSet[Int]())
}

func BenchmarkEach10Safe(b *testing.B) {
	benchEach(b, 10, NewSet[Int]())
}

func BenchmarkEach10Unsafe(b *testing.B) {
	benchEach(b, 10, NewThreadUnsafeSet[Int]())
}

func BenchmarkEach100Safe(b *testing.B) {
	benchEach(b, 100, NewSet[Int]())
}

func BenchmarkEach100Unsafe(b *testing.B) {
	benchEach(b, 100, NewThreadUnsafeSet[Int]())
}

func benchIter(b *testing.B, n Int, s Set[Int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(Int(v))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := s.Iter()
		for range c {

		}
	}
}

func BenchmarkIter1Safe(b *testing.B) {
	benchIter(b, 1, NewSet[Int]())
}

func BenchmarkIter1Unsafe(b *testing.B) {
	benchIter(b, 1, NewThreadUnsafeSet[Int]())
}

func BenchmarkIter10Safe(b *testing.B) {
	benchIter(b, 10, NewSet[Int]())
}

func BenchmarkIter10Unsafe(b *testing.B) {
	benchIter(b, 10, NewThreadUnsafeSet[Int]())
}

func BenchmarkIter100Safe(b *testing.B) {
	benchIter(b, 100, NewSet[Int]())
}

func BenchmarkIter100Unsafe(b *testing.B) {
	benchIter(b, 100, NewThreadUnsafeSet[Int]())
}

func benchIterator(b *testing.B, n Int, s Set[Int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(Int(v))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := s.Iterator().C
		for range c {

		}
	}
}

func BenchmarkIterator1Safe(b *testing.B) {
	benchIterator(b, 1, NewSet[Int]())
}

func BenchmarkIterator1Unsafe(b *testing.B) {
	benchIterator(b, 1, NewThreadUnsafeSet[Int]())
}

func BenchmarkIterator10Safe(b *testing.B) {
	benchIterator(b, 10, NewSet[Int]())
}

func BenchmarkIterator10Unsafe(b *testing.B) {
	benchIterator(b, 10, NewThreadUnsafeSet[Int]())
}

func BenchmarkIterator100Safe(b *testing.B) {
	benchIterator(b, 100, NewSet[Int]())
}

func BenchmarkIterator100Unsafe(b *testing.B) {
	benchIterator(b, 100, NewThreadUnsafeSet[Int]())
}

func benchString(b *testing.B, n Int, s Set[Int]) {
	nums := nrand(n)
	for _, v := range nums {
		s.Add(Int(v))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.String()
	}
}

func BenchmarkString1Safe(b *testing.B) {
	benchString(b, 1, NewSet[Int]())
}

func BenchmarkString1Unsafe(b *testing.B) {
	benchString(b, 1, NewThreadUnsafeSet[Int]())
}

func BenchmarkString10Safe(b *testing.B) {
	benchString(b, 10, NewSet[Int]())
}

func BenchmarkString10Unsafe(b *testing.B) {
	benchString(b, 10, NewThreadUnsafeSet[Int]())
}

func BenchmarkString100Safe(b *testing.B) {
	benchString(b, 100, NewSet[Int]())
}

func BenchmarkString100Unsafe(b *testing.B) {
	benchString(b, 100, NewThreadUnsafeSet[Int]())
}

func benchToSlice(b *testing.B, s Set[Int]) {
	nums := nrand(Int(b.N))
	for _, v := range nums {
		s.Add(Int(v))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.ToSlice()
	}
}

func BenchmarkToSliceSafe(b *testing.B) {
	benchToSlice(b, NewSet[Int]())
}

func BenchmarkToSliceUnsafe(b *testing.B) {
	benchToSlice(b, NewThreadUnsafeSet[Int]())
}
