package admin

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "admin/index.html"
}