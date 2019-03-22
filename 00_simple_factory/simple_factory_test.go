package simple_factory

import "testing"

func TestSimple(t *testing.T) {
	vege := NewFood("vege")
	t.Logf(vege.Desc("tomato"))
	meat := NewFood("meat")
	t.Logf(meat.Desc("fish"))
}
