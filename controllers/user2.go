package controllers

import (
	"github.com/astaxie/beego"
)

type User2Controller struct {
	beego.Controller
}

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

// @router /user2 [get]
func (c *User2Controller) Index() {
	a:=[]User{
		{1,"第一"},
		{2,"第二"},
		{3,"第三"},
	}

	c.Data["json"]=a
	c.ServeJSON()
}
