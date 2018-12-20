package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
)

const (
	//MySQLDriverName : mysql driver name
	MySQLDriverName = "mysql"

	//ErrFailedToConnectToSQL : error message
	ErrFailedToConnectToSQL = "Failed to connect to mysql %v"
)

var (
	engine = &xorm.Engine{}
	gormDB = &gorm.DB{}
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
func Initialize(p *MySQLParamsStruct) (gormDB *gorm.DB, err error) {
	return InitializeGorm(p)
}

//InitializeXorm intialize the xorm
func InitializeXorm(p *MySQLParamsStruct) (err error) {
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

//InitializeGorm initilize the gorm
func InitializeGorm(p *MySQLParamsStruct) (gormDB *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		p.Username, p.Password, p.Hostname, p.Schema)
	gormDB, err = gorm.Open(MySQLDriverName, dsn)
	if err != nil {
		err = fmt.Errorf(ErrFailedToConnectToSQL, err.Error())
		return
	}
	return
}

//GormDB get gorm db
func GormDB() *gorm.DB {
	return gormDB
}

//NewSQLSession : create new sql session
func NewSQLSession() *xorm.Session {
	return engine.NewSession()
}
