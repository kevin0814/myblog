package main

import (
	"github.com/astaxie/beego"
	_ "myblog/routers"
	"myblog/utils"
)

func main() {
	//初始化mysql
	utils.InitMysql()
	beego.SetStaticPath("/down2", "download2") //配置静态资源 访问路径
	beego.BConfig.Listen.Graceful=true //热启动 默认 FALSE 未启动  TRUE 启动
	beego.BConfig.Log.Outputs["console"] = ""//日志输出配置 此参数不支持配置文件配置。
	beego.Run()
}

