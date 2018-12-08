package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

const (
	//MySQLDriverName : mysql driver name
	MySQLDriverName = "mysql"

	//ErrFailedToConnectToSQL : error message
	ErrFailedToConnectToSQL = "Failed to connect to mysql %v"
)

var (
	engine = &xorm.Engine{}
)

//MySQLParamsStruct : mysql params struct
type MySQLParamsStruct struct {
	Username     string
	Password     string
	Hostname     string
	Schema       string
	MaxOpenConns int
}

//Initialize : intialize db
func Initialize(p *MySQLParamsStruct) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		p.Username, p.Password, p.Hostname, p.Schema)
	engine, err = xorm.NewEngine(MySQLDriverName, dsn)
	if err != nil {
		err = fmt.Errorf(ErrFailedToConnectToSQL, err.Error())
		return
	}
	engine.SetMaxOpenConns(p.MaxOpenConns)
	return
}

//NewSQLSession : create new sql session
func NewSQLSession() *xorm.Session {
	return engine.NewSession()
}
