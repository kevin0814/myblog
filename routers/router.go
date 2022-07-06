package routers

import (
	"github.com/astaxie/beego"
	//"github.com/taoshihan1991/imaptool/controller"
	"myblog/controllers"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
	//注册
	beego.Router("/register", &controllers.RegisterController{})
    //登录
    beego.Router("/login",&controllers.LoginController{})
    //首页
	beego.Router("/home",&controllers.HomeController{})
	//退出
	beego.Router("/exit", &controllers.ExitController{})
	//写文章
	beego.Router("/article/add", &controllers.AddArticleController{})
	//显示文章详情
	beego.Router("/article/:id", &controllers.ShowArticleController{})
	//更新文章
	beego.Router("/article/update", &controllers.UpdateArticleController{})
	// 删除文章
	beego.Router("/article/delete", &controllers.DeleteArticleController{})
	//标签
	beego.Router("/tags", &controllers.TagsController{})

	//相册
	beego.Router("/album", &controllers.AlbumController{})

	//文件上传
	beego.Router("/upload", &controllers.UploadController{})

	//关于我页面
	beego.Router("/aboutme",&controllers.AboutMeController{})
    ////测试刊登 亚马逊
	beego.Router("/test/publish",&controllers.TestController{}) //"get:publish;post:publish"
    //查询单条数据
	beego.Router("/prepublish/get_list",&controllers.PrePublishProductController{})
    ////添加数据
	//查询单条数据
	beego.Router("/prepublish/add_listing",&controllers.AddPublishProductController{})
	//beego.Router("/prepublish/add_listing",&controllers.PrePublishProductController{},"post:addListing")
	//beego.Post("/api", func(ctx *context.Context) {
	//	ctx.Output.Body([]byte("hello world"))
	//})
	//beego.Router("/api/list",&controllers.api.RestController{},"*:ListFood")

}
