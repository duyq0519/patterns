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
func TestSingleNil(t *testing.T) {
	leader1 := GetLeader()
	var ll leader
	(*leader1) = ll
	t.Log(leader1) //nil
	leader2 := GetLeader()
	t.Log(leader2) //非空
}
