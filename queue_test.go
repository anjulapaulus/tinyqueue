package tinyqueue

import (
	"testing"
)

type IntVal int64

func (a IntVal) Compare (b Item) bool{
	return a <= b.(IntVal)
}

func TestQueue_Push(t *testing.T) {

	q := Queue{}
	q.Push(IntVal(7))
	q.Push(IntVal(10))
	q.Push(IntVal(5))

	if q.Len() != 3{
		t.Fatal("Push Test : Failed")
	}
}
func TestQueue_PushNil(t *testing.T) {

	q := Queue{}
	q.Push(nil)

	if q.Len() != 0{
		t.Fatal("Push Test : Failed")
	}
}

func TestQueue_Pop(t *testing.T) {
	q := Queue{}
	q.Push(IntVal(7))
	q.Push(IntVal(5))
	q.Push(IntVal(10))

	if q.Pop() != IntVal(5) {
		t.Fatal("Pop Test : Failed")
	}
}
func TestQueue_PopNil(t *testing.T) {
	q := Queue{}
	if q.Pop() != nil {
		t.Fatal("Pop Test : Failed")
	}
}


func TestQueue_Len(t *testing.T) {
	q := Queue{}
	q.Push(IntVal(7))
	q.Push(IntVal(5))
	q.Push(IntVal(10))

	if q.Len() != 3{
		t.Fatal("Len Test: Failed")
	}
}

func TestQueue_All(t *testing.T) {
	q := Queue{}
	q.Push(IntVal(7))
	q.Push(IntVal(5))
	q.Push(IntVal(10))

	array := q.All()
	if array[0] != IntVal(5) && array[1] != IntVal(7) && array[2] != IntVal(10){
		t.Fatal("All Test: Failed")
	}
}

func TestQueue_Peek(t *testing.T) {
	q := Queue{}
	q.Push(IntVal(10))
	q.Push(IntVal(7))
	q.Push(IntVal(5))


	if q.Peek() != IntVal(5) {
		t.Fatal("PeekMin Test: Failed")
	}
}


