package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin   bool
	Loginuser interface{}
}

//判断是否登录
/*
	这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，
    用户可以重写这个函数实现类似用户验证之类。
*/
func (this *BaseController) Prepare() {
	loginuserId := this.GetSession("loginuserId")
	//fmt.Println("loginuserId---->", loginuserId)
	if loginuserId != nil {
		this.IsLogin = true
		this.Loginuser = loginuserId
	} else {
		this.IsLogin = false
	}
	this.Data["IsLogin"] = this.IsLogin
}

//ajax返回 列表
func (self *BaseController) ajaxList(msg interface{}, msgno int, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

