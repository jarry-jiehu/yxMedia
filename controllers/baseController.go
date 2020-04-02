package controllers

import (
	"github.com/astaxie/beego"
	"yxMedia/models"
)

type BaseController struct {
	beego.Controller
}

type Response struct {
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	BusinessObj interface{} `json:"businessObj"`
}

func (this *BaseController) Prepare() {

}

func (this *BaseController) response(code int, message string, data interface{}) {

	resp := &Response{
		Code:        code,
		Message:     message,
		BusinessObj: data,
	}

	this.Data["json"] = &resp
	this.ServeJSON()
}

func (this *BaseController) Options() {
	this.response(models.RESP_OK, "Success", nil)
}
