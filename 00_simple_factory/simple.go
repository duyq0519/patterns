package simple_factory

import "fmt"

/*
优点：简单，适合对象比较少的情况，可以创建任何对象
缺点：	 不适用很多的对象，
		不符合开闭原则，对于新增对象无法处理
		实现了所有类的创建逻辑，违反了高内聚的责任分配原则
*/

type Food interface {
	Desc(string) string
}

type Vege struct {
}

func (v *Vege) Desc(name string) string {
	return fmt.Sprintf("vege %s", name)
}

type Meat struct {
}

func (v *Meat) Desc(name string) string {
	return fmt.Sprintf("meat %s", name)
}

func NewFood(t string) Food {
	switch t {
	case "vege":
		return &Vege{}
	case "meat":
		return &Meat{}
	}
	return nil
}
