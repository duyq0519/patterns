package factory

import (
	"fmt"
	"testing"
)

func TestOperatorFactory(t *testing.T) {
	plus := &PlusOperatorFactory{}
	p := plus.Create()
	p.SetA(4)
	p.SetB(6)
	fmt.Printf("4+6=%d", p.Result())

	fmt.Println()

	minus := &MinusOperatorFactory{}
	m := minus.Create()
	m.SetA(5)
	m.SetB(1)
	fmt.Printf("5-1=%d", m.Result())

	fmt.Println()

}
