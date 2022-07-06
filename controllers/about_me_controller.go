package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {

	c.Data["wechat"] = "微信：正义之手"
	c.Data["qq"] = "QQ：洞拐洞拐不动就乖"
	c.Data["tel"] = "Tel：正义之手雄起"

	c.TplName = "aboutme.html"
}