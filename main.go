package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql" //导入数据库驱动
	_ "go-admin-test-system/routers"
)

func init() {
	// need to register models in init
	//orm.RegisterModel(new(models.User))

	// need to register db driver
	err:=orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Println("注册数据库驱动失败：",err)
	}

	// need to register default database
	err_:=orm.RegisterDataBase("default", "mysql", "test-system:test-system@tcp(127.0.0.1:3306)/test-system?charset=utf8&loc=Asia%2FShanghai")
	if err_ != nil {
		fmt.Println("RegisterDataBase失败：",err_)
	}
}

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type","X-Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.Run()
}

