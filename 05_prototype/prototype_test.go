package prototype

import "testing"

func TestPrototype(t *testing.T) {
	cpool := NewGaoTongCPUManager("localhost:27017")
	conn := cpool.Get()
	result := conn.Compute("855")
	t.Logf("your cpu is %s\n", result)
}
