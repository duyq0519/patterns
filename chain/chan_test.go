package chain

import "testing"

func TestChain(t *testing.T) {
	group := NewHandlers()

	group.HandleReq(Group, "组内小事")
	group.HandleReq(DepartMent, "部门事务")
	group.HandleReq(Company, "公司大事")
	group.HandleReq(GroupCompany, "集团公司大事")
}
