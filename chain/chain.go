package chain

import "fmt"

const (
	Group = iota
	DepartMent
	Company
	GroupCompany
)

type Organization struct {
	level   int
	NextOrg *Organization
}

func NewOrganization(level int) *Organization {
	return &Organization{level: level}
}

func (o *Organization) WriteMsg(msg string) {
	part := ""
	switch o.level {
	case Group:
		part = "Group"
	case DepartMent:
		part = "Department"
	case Company:
		part = "GroupCompany"
	default:
		part = "Department"
	}
	fmt.Printf("level %s,handler %s\n", part, msg)
}

func (o *Organization) HandleReq(level int, msg string) {
	if o.level >= level {
		o.WriteMsg(msg)
	} else if o.NextOrg != nil {
		o.NextOrg.HandleReq(level, msg)
	}
}
func NewHandlers() *Organization {
	group := NewOrganization(Group)
	department := NewOrganization(DepartMent)
	company := NewOrganization(Company)
	groupCompany := NewOrganization(GroupCompany)

	group.NextOrg = department
	department.NextOrg = company
	company.NextOrg = groupCompany
	return group
}
