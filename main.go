package main

import (
	"fmt"
	_ "genesis/docs"
	_ "genesis/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func init() {

	dbName := beego.AppConfig.String("db.name")
	dbType := beego.AppConfig.String("db.type")

	userName := beego.AppConfig.String("db.user.name")
	password := beego.AppConfig.String("db.user.password")
	ip := beego.AppConfig.String("db.address.ip")
	port := beego.AppConfig.String("db.address.port")

	dbConn := userName + ":" + password + "@tcp(" + ip + ":" + port + ")/" + dbName
	//orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	fmt.Println("dbType:" + dbType)
	fmt.Println("dbName:" + dbName)
	fmt.Println("dbConn:" + dbConn)
	orm.RegisterDataBase("default", dbType, dbConn)
	orm.Debug = true
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
