package abstract_factory

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	fa := GDDAOFactory{}
	mgoa := fa.CreateMongoDAO()
	mgoa.Find()
	sqla := fa.CreateMysqlDAO()
	sqla.Query()

	fb := AliDAOFactory{}
	mgob := fb.CreateMongoDAO()
	mgob.Find()
	sqlb := fb.CreateMysqlDAO()
	sqlb.Query()
	fmt.Print("\n")
}
