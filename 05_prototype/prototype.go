package prototype

import "fmt"

type CPU interface {
	Compute(string) string
}

type GaoTongCPU struct {
	CoreType string
}

func newGaoTongCPU(cstr string) *GaoTongCPU {
	return &GaoTongCPU{
		CoreType: cstr,
	}
}

func (c *GaoTongCPU) Compute(cstr string) string {
	return fmt.Sprintf("GaoTong %s", cstr)
}

type CPUManager struct {
	CoreType string
	pool     []CPU
}

func NewGaoTongCPUManager(cstr string) *CPUManager {
	cm := CPUManager{
		pool: []CPU{},
	}
	mc := newGaoTongCPU(cstr)
	cm.pool = append(cm.pool, mc)
	return &cm
}

func (c *CPUManager) Get() CPU {
	if len(c.pool) == 0 {
		mc := newGaoTongCPU(c.CoreType)
		c.pool = append(c.pool, mc)
	}
	conn := c.pool[0]
	return conn
}
