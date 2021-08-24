package queue

type Queue []interface{}

func (q *Queue) IsEmpty() bool          { return len(*q) == 0 }
func (q *Queue) Len() int               { return len(*q) }
func (q *Queue) Enqueue(el interface{}) { *q = append(*q, el) }
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element
}
