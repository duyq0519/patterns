package abstract_factory

import "fmt"

type Mongo interface {
	Find()
}
type Mysql interface {
	Query()
}
type DAOFactory interface {
	CreateMongoDAO() Mongo
	CreateMysqlDAO() Mysql
}

type MongoDriverA struct{}

func (m *MongoDriverA) Find() {
	fmt.Println("this is mongo A")
}

type MysqlDriverA struct{}

func (m *MysqlDriverA) Query() {
	fmt.Println("this is mysql A")
}

//实现工厂
type GDDAOFactory struct{}

func (*GDDAOFactory) CreateMongoDAO() Mongo {
	return &MongoDriverA{}
}
func (*GDDAOFactory) CreateMysqlDAO() Mysql {
	return &MysqlDriverA{}
}

/*-------------------------------*/
type MongoDriverB struct{}

func (m *MongoDriverB) Find() {
	fmt.Println("this is mongo B")
}

type MysqlDriverB struct{}

func (m *MysqlDriverB) Query() {
	fmt.Printf("this is mysql B")
}

type AliDAOFactory struct{}

func (*AliDAOFactory) CreateMongoDAO() Mongo {
	return &MongoDriverB{}
}
func (*AliDAOFactory) CreateMysqlDAO() Mysql {
	return &MysqlDriverB{}
}
