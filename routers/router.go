package routers

import (
	"github.com/alan-liu2020/xblog/controllers"
	"github.com/astaxie/beego"
	"github.com/alan-liu2020/xblog/controllers/admin"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"strings"
	"github.com/alan-liu2020/xblog/models"
	"github.com/alan-liu2020/xblog/db"
	"encoding/json"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		//AllowOrigins:      []string{"https://192.168.0.102"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Router("/", &controllers.MainController{})
	adminNS := beego.NewNamespace("admin",
		beego.NSBefore(auth),
		beego.NSAutoRouter(&admin.IndexController{}),
		beego.NSAutoRouter(&admin.UserController{}),
		beego.NSAutoRouter(&admin.ArticleController{}),
		beego.NSGet("/notallowed", func(ctx *context.Context) {ctx.Output.Body([]byte("notAllowed"))}),
			)
	beego.AddNamespace(adminNS)
}

// 校验登录
func auth(c *context.Context) {
	whiteList := []string{
		"/admin/user/dologin",
		"/admin/index",
	}
	for _, value := range whiteList {
		if strings.Contains(c.Request.URL.Path, value) {
			return
		}
	}
	auth := c.Input.Header("Authorization")
	slices := strings.Split(auth, " ")
	if len(slices) > 1 {
		token := strings.Trim(slices[1], " ")
		var adminUsers []*models.AdminUser
		db.DB.QueryTable(new(models.AdminUser)).Filter("token", token).All(&adminUsers)
		if len(adminUsers) > 0{
			//登陆成功
			c.Input.SetData("adminUser", adminUsers[0])
		} else {
			result := admin.GetResult()
			result.Code = 1
			result.Msg = "Token无效，禁止登陆!"
			bytes, _ := json.Marshal(&result)
			c.Output.Body(bytes)
		}
	} else {
		result := admin.GetResult()
		result.Code = 1
		result.Msg = "Forbidden! Not Login!"
		bytes, _ := json.Marshal(&result)
		c.Output.Body(bytes)
	}
}
