package stack

type Stack []interface{}

func (s *Stack) IsEmpty() bool       { return len(*s) == 0 }
func (s *Stack) Len() int            { return len(*s) }
func (s *Stack) Peek() interface{}   { return (*s)[len(*s)-1] }
func (s *Stack) Push(el interface{}) { *s = append(*s, el) }
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	element := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return element
}

// GenericStack generic stack
type GenericStack[T any] []T

func (s *GenericStack[T]) IsEmpty() bool { return len(*s) == 0 }
func (s *GenericStack[T]) Len() int      { return len(*s) }
func (s *GenericStack[T]) Peek() T       { return (*s)[len(*s)-1] }
func (s *GenericStack[T]) Push(el T)     { *s = append(*s, el) }
func (s *GenericStack[T]) Pop() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}
	element := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return element
}
