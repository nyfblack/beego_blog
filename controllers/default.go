package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type User struct {
	Id   int
	Name string
	Age  int
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	c.Data["Truecond"] = true
	c.Data["Falsecond"] = false

	c.Data["User"] = User{1, "ok", 12}

	c.Data["Num"] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

}
