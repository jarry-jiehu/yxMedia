package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"yxMedia/models"
	"yxMedia/utils"
)

type UserController struct {
	BaseController
}

func (this *UserController) LogIn() {
	name := this.GetString("name", "")
	pwd := this.GetString("pwd", "")
	phone := this.GetString("phone", "")
	if name == "" || pwd == "" || phone == "" {
		beego.Error("输入的用户名或者密码错误")
		this.response(models.RESP_LOGIN_FAILED, "输入的用户名或者密码错误", nil)
	}

	user := &models.UserModel{
		Name:     name,
		Password: pwd,
		Phone:    phone,
	}

	isOk := user.GetByName(name)
	if isOk == false {
		beego.Error("登录失败，用户名或密码错误")
		this.response(models.RESP_LOGIN_FAILED, "登录失败，用户名或密码错误", nil)
	}

	//需要增加token授权，to do 。。。。。
	this.response(models.RESP_OK, "登录成功", nil)
}

func (this *UserController) LogOn() {
	name := this.GetString("name", "")
	pwd := this.GetString("pwd", "")
	phone := this.GetString("phone", "")

	if name == "" || pwd == "" {
		beego.Error("输入的用户名或者密码错误")
		this.response(models.RESP_LOGON_FAILED, "输入的用户名或者密码错误", nil)
	}

	user := &models.UserModel{
		Name:       name,
		Password:   pwd,
		Phone:      phone,
		CreateTime: utils.UnixToString(time.Now().Unix()),
	}

	err := user.Insert(user)
	if err != nil {
		beego.Error("注册失败，请重新注册，用户名可能重复")
		this.response(models.RESP_LOGON_FAILED, "注册失败，请重新注册，用户名可能重复", nil)
	}

	this.response(models.RESP_OK, "注册成功", nil)
}
