package main

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id   int
	Name string
	Prov string
	Time int
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
	//	orm.RegisterModelWithPrefix("tp_", new(User))
}
