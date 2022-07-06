package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblog/models"
)

/**
 * 相册控制器
 */
type AlbumController struct {
	BaseController
}

func (this *AlbumController) Get() {
	albums, err := models.FindAllAlbums()
	if err != nil {
		beego.BeeLogger.Error(err.Error())
	}
	fmt.Println(albums)
	this.Data["Album"] = albums
	this.TplName = "album.html"
}


