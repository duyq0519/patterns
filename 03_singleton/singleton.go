package singleton

import (
	"fmt"
	"sync"
)

var l *leader

var once sync.Once

type leader struct {
	Name string
}

func (l *leader) Look() {
	fmt.Printf("hello %s", l.Name)
}

func GetLeader() *leader {
	once.Do(func() {
		l = &leader{Name: "duyq"}
	})
	return l
}
