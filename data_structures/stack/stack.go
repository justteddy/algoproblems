package stack

type Stack []interface{}

func (s *Stack) IsEmpty() bool       { return len(*s) == 0 }
func (s *Stack) Len() int            { return len(*s) }
func (s *Stack) Push(el interface{}) { *s = append(*s, el) }
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	element := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return element
}
