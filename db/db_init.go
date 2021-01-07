package db

import (
	"fmt"
	"github.com/alan-liu2020/xblog/com"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/astaxie/beego/orm"
	"log"
	"github.com/alan-liu2020/xblog/models"
)

var DB orm.Ormer
var PGDB orm.Ormer

func init() {
	initMYSQL()
	//initPGSQL()
}

func initMYSQL()  {
	cnf, err := com.AppConfig()
	runmode := cnf.String("runmode")
	dbHost := cnf.String(runmode+"::db_host")
	dbPort := cnf.String(runmode+"::db_port")
	dbUser := cnf.String(runmode+"::db_user")
	dbPassword := cnf.String(runmode+"::db_password")
	dbName := cnf.String(runmode+"::db_name")
	CheckErr(err, "ini config not exist!")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local",dbUser,dbPassword,dbHost,dbPort,dbName)
	orm.RegisterDataBase("default", "mysql",url , 30)

	// register model
	orm.RegisterModel(new(models.AdminUser), new(models.Article))

	// create table
	orm.RunSyncdb("default", false, true)
	// debug sql
	//orm.Debug = true

	DB = orm.NewOrm()
}

func initPGSQL()  {
	cnf, err := com.AppConfig()
	runmode := cnf.String("runmode")
	dbHost := cnf.String(runmode+"::pgdb_host")
	dbPort := cnf.String(runmode+"::pgdb_port")
	dbUser := cnf.String(runmode+"::pgdb_user")
	dbPassword := cnf.String(runmode+"::pgdb_password")
	dbName := cnf.String(runmode+"::pgdb_name")
	CheckErr(err, "ini config not exist!")
	if dbHost == "" {
		return
	}

	url := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",dbUser,dbPassword,dbHost,dbPort,dbName)
	orm.RegisterDataBase("pgsql", "postgres",url , 30)

	// register model
	//orm.RegisterModel()

	// create table
	//orm.RunSyncdb("pgsql", false, true)
	// debug sql
	//orm.Debug = true

	PGDB = orm.NewOrm()
	PGDB.Using("pgsql")
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
	}
}
