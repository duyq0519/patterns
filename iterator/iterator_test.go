package iterator

import (
	"testing"
)

func TestIterator(t *testing.T) {
	list := List([]Data{Data(30), Data(45), Data(5), Data(546), Data(65), Data(65), Data(435)})
	iter := list.GetInterator()
	for iter.HasNext() {
		v := iter.Next()
		t.Log(v)
	}
}
