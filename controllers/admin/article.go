package admin

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/alan-liu2020/xblog/models"
	"encoding/json"
	"time"
	"github.com/alan-liu2020/xblog/db"
	"errors"
	"github.com/alan-liu2020/xblog/com"
)

type ArticleController struct {
	beego.Controller
	Params com.XMap
	Result *Result
	AdminUser *models.AdminUser
}

func (c *ArticleController) Prepare() {
	c.Result = GetResult()
	json.Unmarshal(c.Ctx.Input.RequestBody, &c.Params)
	fmt.Println("传入参数为:", string(c.Ctx.Input.RequestBody))
	c.AdminUser = c.Ctx.Input.GetData("adminUser").(*models.AdminUser)
}

func (c *ArticleController) Finish() {
	c.Data["json"] = c.Result
	c.ServeJSON(true)
}

func (c *ArticleController) Create() {

	var article models.Article
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &article); err != nil {
		CheckErr(err, c.Result)
		return
	}

	if article.Title == "" {
		CheckErr(errors.New("标题不能为空!"), c.Result)
		return
	}

	if article.Slug == "" {
		CheckErr(errors.New("Slug不能为空!"), c.Result)
		return
	}

	if article.Content == "" {
		CheckErr(errors.New("内容不能为空!"), c.Result)
		return
	}

	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()
	if article.Id > 0 {
		_, err := db.DB.Update(&article)
		if err != nil {
			CheckErr(err, c.Result)
			return
		}
		return
	}
	_, err := db.DB.Insert(&article)
	if err != nil {
		CheckErr(err, c.Result)
		return
	}

	c.Result.Msg = "创建文章成功!"
}

func (c *ArticleController) List() {
	result := GetResult()
	defer func() {
		c.Data["json"] = result
		c.ServeJSON(true)
	}()

	var list []*models.Article
	db.DB.QueryTable(new(models.Article)).RelatedSel().All(&list)
	result.Data = list
}

func (c *ArticleController) One() {
	result := GetResult()
	defer func() {
		c.Data["json"] = result
		c.ServeJSON(true)
	}()

	var article models.Article
	article.Id = c.Params.GetInt64("id", 0)
	err := db.DB.Read(&article)
	if err != nil {
		CheckErr(err, result)
		return
	}

	result.Data = article
}