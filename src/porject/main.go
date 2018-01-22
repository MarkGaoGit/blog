package main

import (
	_ "porject/routers"

	//"btest/controllers" //func undefined

	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/tptest?charset=utf8")
	orm.DefaultTimeLoc = time.UTC
	orm.Debug = true // print sql

}

func main() {
	beego.SetStaticPath("/static", "static") //setting the static file path
	//beego.ErrorController(&controllers.ErrorController{}) // Registration error processing function error: undefined controllers

	beego.SetLevel(beego.LevelInformational) //setting the error level
	beego.SetLogFuncCall(true)               //output filename and line number

	beego.Run()
}
