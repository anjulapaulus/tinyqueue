package tinyqueue

type Queue struct {
	length 		int
	data   		[]Item
}

type Item interface {
	Compare(b Item) bool
}

func (q *Queue) Len () int{
	return q.length
}

func (q *Queue) Peek() Item{
	return q.data[0]
}


func (q *Queue) All() []Item{
	return q.data
}

func (q *Queue) Push (item Item){
	if item == nil{
		return
	}
	q.data = append(q.data,item)
	q.length++
	q.up(q.length-1)
}

func (q *Queue) up (pos int){
	data := q.data
	item := data[pos]
	for pos > 0 {
		parent := (pos-1) >> 1
		current := data[parent]
		if !item.Compare(current){
			break
		}
		data[pos] = current
		pos = parent
	}
	data[pos] = item
}

func (q *Queue) Pop() Item {
	if q.length == 0 {
		return nil
	}
	top := q.data[0]
	q.length--
	if q.length > 0{
		q.data[0] = q.data[q.length]
		q.down(0)
	}
	q.data = q.data[:len(q.data)-1]
	return top
}

func (q *Queue) down (pos int){
	data := q.data
	halfLength := q.length >> 1
	item := q.data[pos]

	for pos < halfLength {
		left := (pos << 1) + 1
		best := q.data[left]
		right := left + 1

		if right < q.length && data[right].Compare(best) {
			left = right
			best = data[right]
		}

		if !best.Compare(item) {
			break
		}
		data[pos] = best
		pos = left
	}
	data[pos] = item
}
