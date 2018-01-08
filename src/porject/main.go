package main

import (
	_ "porject/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static", "static") //设置静态文件路径
	beego.Run()
}
