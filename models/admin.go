package models

import "time"

type AdminUser struct {
	Id int64 `orm:"null" json:"id"`
	Username string `orm:"null;unique" json:"username"`
	Password string `orm:"null" json:"password"`
	Salt string `orm:"null;size(6)" json:"salt"`
	Token string `orm:"null;size(32)" json:"token"`
	CreatedAt time.Time `orm:"null" json:"created_at"`
	UpdatedAt time.Time `orm:"null" json:"updated_at"`
	LastLoginTime time.Time `orm:"null" json:"last_login_time"`
}

type Article struct {
	Id int64 `orm:"null" json:"id"`
	Title string `orm:"null;" json:"title"`
	Content string `orm:"null" json:"content"`
	Slug string `orm:"null" json:"slug"`
	CreatedAt time.Time `orm:"null" json:"created_at"`
	UpdatedAt time.Time `orm:"null" json:"updated_at"`
}
