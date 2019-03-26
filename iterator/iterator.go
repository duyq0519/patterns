package iterator

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() Data
}

type Container interface {
	GetInterator() Iterator
}

type Data int

func (d Data) String() string {
	return fmt.Sprintf("d:%d", d)
}

type List []Data

func (l List) GetInterator() *ListIterator {
	return &ListIterator{
		List: l,
		next: 0,
	}
}

type ListIterator struct {
	List
	next int
}

func (l *ListIterator) HasNext() bool {
	return l.next < len(l.List) && l.next >= 0
}
func (l *ListIterator) Next() interface{} {
	if l.HasNext() {
		next := l.List[l.next]
		l.next++
		return next
	}
	return nil
}
