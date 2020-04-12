package tinyqueue

type Queue struct {
	length 	int
	ids 	[]interface{}
	values 	[]Item
}

type Item interface {
	Less (Item) bool
}

func (q *Queue) Push (id interface{}, value Item){
	q.ids = append(q.ids,id)
	q.values = append(q.values,value)
	q.length ++
	pos := q.length -1

	for pos > 0 {
		parent := (pos -1) >> 1
		parentValue := q.values[parent]
		if !q.values[pos].Less(parentValue) {
			break
		}
		q.ids[pos] = q.ids[parent]
		q.values[pos] = parentValue
		pos = parent
	}
	q.ids[pos] = id
	q.values[pos] = value
}
func (q *Queue) Pop () interface{}{
	if q.length == 0{
		return nil
	}

	top := q.ids[0]
	q.length--

	if q.length > 0 {
		q.ids[0] = q.ids[q.length]
		q.values[0] = q.values[q.length]

		id := q.ids[0]
		value := q.values[0]
		halfLength := q.length >> 1
		pos := 0

		for pos < halfLength {
			left := (pos << 1) + 1
			right := left + 1
			bestIndex := q.ids[left]
			bestValue := q.values[left]
			rightValue := q.values[right]

			if right < q.length && rightValue.Less(bestValue) {
				left = right
				bestIndex = q.ids[right]
				bestValue = rightValue
			}
			if !bestValue.Less(value) {
				break
			}

			q.ids[pos] = bestIndex
			q.values[pos] = bestValue
			pos = left
		}
		q.ids[pos] = id
		q.values[pos] = value
	}
	return top
}

func (q *Queue) Peek() interface{} {
	if q.length == 0 {
		return nil
	}
	return q.ids[0]
}

func (q *Queue) PeekValue() Item{
	if q.length == 0 {
		return nil
	}
	return q.values[0]
}

func (q *Queue) Len() int{
	return q.length
}