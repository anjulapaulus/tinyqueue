package tinyqueue

import (
	"fmt"
	"testing"
)

type floatValue float64

func (a floatValue) Less(b Item) bool {
	return a >= b.(floatValue)
}

func TestQueue_Push(t *testing.T) {
	var q Queue
	q.Push(0, floatValue(1))
	q.Push(1, floatValue(2))
	 if q.Len() != 2{
	 	t.Error("Push Function Failed")
	 }
}

func BenchmarkQueue_Push(b *testing.B) {
	var q Queue
			for i := 0; i < b.N; i++ {
				q.Push(0, floatValue(1))
			}
}

func TestQueue_Len(t *testing.T) {
	var q Queue
	q.Push(0, floatValue(1))
	q.Push(1, floatValue(2))
	if q.Len() != 2{
		t.Error("Push Function Failed")
	}
}

func TestQueue_Peek(t *testing.T) {
	var q Queue
	q.Push(0, floatValue(1))
	q.Push(1, floatValue(2))

	if q.Peek() != 1{
		t.Error("Peek Function Failed")
	}
}

func TestQueue_PeekValue(t *testing.T) {
	var q Queue
	q.Push(0, floatValue(1))
	q.Push(1, floatValue(2))

	val := q.PeekValue()
	fmt.Println(val)
}

func TestQueue_Pop(t *testing.T) {
	var q Queue
	q.Push(0, floatValue(1))
	q.Push(1, floatValue(2))

	if q.Pop() != 1{
		t.Error("Pop function failed")
	}
}