package admin

import (
	"github.com/astaxie/beego"
	"github.com/alan-liu2020/xblog/models"
	"time"
	"github.com/alan-liu2020/xblog/com"
	"github.com/alan-liu2020/xblog/db"
	"errors"
	"encoding/json"
	"fmt"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Prepare() {

}

func (c *UserController) Login() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "admin/login.html"
}

func (c *UserController) DoLogin() {
	result := GetResult()
	defer func() {
		c.Data["json"] = result
		c.ServeJSON(true)
	}()

	var adminUser models.AdminUser
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &adminUser); err != nil {
		CheckErr(err, result)
		return
	}

	inputPwd := adminUser.Password

	err := db.DB.Read(&adminUser, "Username")
	if err != nil {
		CheckErr(errors.New("用户名密码不存在!"), result)
		return
	}

	if adminUser.Password == com.MD5(inputPwd+adminUser.Salt) {
		//登陆成功
		adminUser.LastLoginTime = time.Now()
		adminUser.Token = com.MD5(adminUser.Password+fmt.Sprintf("%v", adminUser.LastLoginTime))
		db.DB.Update(&adminUser, "token", "LastLoginTime")
		result.Data = adminUser.Token
	} else {
		result.Code = 1
		result.Msg = "输入密码不正确!"
	}

}

func (c *UserController) Regist() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "admin/regist.html"
}

func (c *UserController) DoRegist() {
	result := GetResult()
	defer func() {
		c.Data["json"] = result
		c.ServeJSON(true)
	}()

	username := c.Input().Get("username")
	password := c.Input().Get("password")
	salt := com.RandStringRunes(6)
	adminUser := models.AdminUser{
		Username:username,
		Password:com.MD5(password+salt),
		Salt:salt,
		CreatedAt:time.Now(),
	}
	_, err := db.DB.Insert(&adminUser)
	if err != nil {
		CheckErr(err, result)
		return
	}

	return
}