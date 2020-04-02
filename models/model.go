package models

import (
	log "code.google.com/p/log4go"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	RESP_OK                 = 10000
	RESP_LOGON_FAILED       = 10001
	RESP_LOGIN_FAILED       = 10002
	RESP_PLAY_ERROR         = 10003
	RESP_PLAY_INSERT_FAILED = 10004
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	var host string = beego.AppConfig.String("db.mysql.host")
	var username string = beego.AppConfig.String("db.mysql.username")
	var password string = beego.AppConfig.String("db.mysql.password")
	var dbName string = beego.AppConfig.String("db.mysql.dbname")
	dbDNS := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, dbName)
	log.Info("Connect mysqll , dbDNS is  %s", dbDNS)
	beego.Info("db host : ", host, ", username: ", username)
	err := orm.RegisterDataBase("default", "mysql", dbDNS)
	if err != nil {
		log.Info("Connect mysql failed %s", err.Error())
	} else {
		log.Info("connect mysql success!")
	}
	orm.RegisterModel(new(UserModel))
}
