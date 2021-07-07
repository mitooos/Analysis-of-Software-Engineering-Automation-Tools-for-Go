package benchmark

import (
	"testing"
	"time"
)

var result error

func action(i int) {
	time.Sleep(10)
}

func BenchmarkTraverseLinkedList(b *testing.B) {
	linkedL := &linkedList{}
	for i := 0; i < 100; i++ {
		node := &node{data: i}
		linkedL.push(node)
	}

	var err error
	for n := 0; n < b.N; n++ {
		err = linkedL.traverse(action)
	}

	result = err
}

func BenchmarkTraverseSliceReference(b *testing.B) {
	sl := make(slice, 100)
	for i := 0; i < 100; i++ {
		sl[i] = i
	}

	var err error
	for n := 0; n < b.N; n++ {
		err = sl.traverseReference(action)
	}

	result = err
}

func BenchmarkTraverseSliceValue(b *testing.B) {
	sl := make(slice, 100)
	for i := 0; i < 100; i++ {
		sl[i] = i
	}

	var err error
	for n := 0; n < b.N; n++ {
		err = sl.traverseValue(action)
	}

	result = err
}
