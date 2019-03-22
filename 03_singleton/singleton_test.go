package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	for i := 0; i < 10; i++ {
		leader := GetLeader()
		t.Logf("leader is %s", leader.Name)
		if i == 4 {
			leader.Name = "wdd"
		}
	}

}
