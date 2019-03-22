package singleton

import "sync"

var leader *Leader

//此包只能有一个once，才能保证只执行一次
var once sync.Once

type Leader struct {
	Name string
}

func newSingleton() *Leader {
	return &Leader{
		Name: "duyq",
	}
}

//获取单例模式的对象，不用进行非空判断也是正确的
func GetLeader() *Leader {
	once.Do(func() {
		leader = newSingleton()
	})
	return leader
}
