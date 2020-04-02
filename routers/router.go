package routers

import (
	"github.com/astaxie/beego"
	"yxMedia/controllers"
)

func init() {
	ss := beego.NewNamespace("/v1",
		//认证，注册接口
		beego.NSNamespace("/auth",
			beego.NSRouter("login", &controllers.UserController{}, "post:LogIn"),
			beego.NSRouter("logon", &controllers.UserController{}, "post:LogOn"),
		),
		//播放列表获取
		beego.NSNamespace("/play",
			beego.NSRouter("playinfo", &controllers.PlayInfoController{}, "post:PlayInfo"),
		),
		//app应用路由表
		beego.NSNamespace("/play",
			beego.NSRouter("create", &controllers.PlayInfoController{}, "post:PlayInfo"),
		),
	)

	beego.AddNamespace(ss)
}
