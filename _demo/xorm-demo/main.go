package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

const (
	DriverName           = "mysql"
	MasterDataSourceName = "root:123456@tcp(127.0.0.1:3306)/go-superstar?charset=utf8"
)

var (
	engine *xorm.Engine
	err    error
)

type UserInfo struct {
	Id         int `xorm:"not null pk autoincr"`
	Name       string
	SysCreated int
	SysUpdated int
}

func newEngine() *xorm.Engine {
	engine, err := xorm.NewEngine(DriverName, MasterDataSourceName)
	if err != nil {
		log.Fatal(engine, err)
		return nil
	}
	engine.ShowSQL(true)
	return engine
}
func main() {
	engine = newEngine()
	ormFindRows()
}

func ormFindRows() {
	list := make([]UserInfo, 0)
	log.Printf("engine: %v", engine)
	err := engine.Cols("id", "name").Where("id>?", 0).
		Limit(10).Asc("id", "sys_created").Find(&list)
	if err != nil {
		log.Fatal("ormFindRows error", err)
	}
	fmt.Printf("%v\n", list)
}
