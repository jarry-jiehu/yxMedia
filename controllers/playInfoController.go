package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"yxMedia/models"
)

type PlayInfoController struct {
	BaseController
}

func (this *PlayInfoController) Insert() {
	playinfo := this.ParsePlanInfo()

	if playinfo.Name == "" || playinfo.Title == "" {
		this.response(models.RESP_PLAY_ERROR, "添加媒资库失败", nil)
		return
	}

	err := playinfo.Insert(playinfo)
	if err != nil {
		beego.Info("err :", err)
		this.response(models.RESP_PLAY_ERROR, "添加媒资库失败", nil)
		return
	}

	this.response(models.RESP_OK, "success", nil)
}

func (this *PlayInfoController) PlayInfo() {
	name := this.GetString("name", "")

	if name == "" {
		this.response(models.RESP_PLAY_ERROR, "获取媒资信息参数失败", nil)
		return
	}

	var play models.PlayModel
	playmodel, err := play.GetByName(name)
	if err != nil {
		this.response(models.RESP_PLAY_ERROR, "获取媒资信息失败", nil)
		return
	}

	this.response(models.RESP_OK, "success", playmodel)
}

func (this *PlayInfoController) PlayList() {

}

func (this *PlayInfoController) ParsePlanInfo() *models.PlayModel {
	body := this.Ctx.Input.RequestBody

	beego.Info("ParseAppInfo bode:%s", body)
	if len(body) == 0 {
		this.response(models.RESP_PLAY_INSERT_FAILED, "param null", nil)
		return nil
	}

	var planInfo models.PlayModel
	err := json.Unmarshal(body, &planInfo)
	if err != nil {
		beego.Info("post body error, " + err.Error())
		this.response(models.RESP_PLAY_ERROR, " 参数解析出错! "+err.Error(), nil)
		return nil
	}

	return &planInfo
}
