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
	engine     = &xorm.Engine{}
	gormEngine = &gorm.DB{}
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
func InitializeGorm(p *MySQLParamsStruct) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		p.Username, p.Password, p.Hostname, p.Schema)
	gormEngine, err = gorm.Open(MySQLDriverName, dsn)
	if err != nil {
		err = fmt.Errorf(ErrFailedToConnectToSQL, err.Error())
		return
	}
	defer gormEngine.Close()
	return
}

//NewSQLSession : create new sql session
func NewSQLSession() *xorm.Session {
	return engine.NewSession()
}

//GormEngine return new gorm engine
func GormEngine() *gorm.DB {
	return gormEngine
}
