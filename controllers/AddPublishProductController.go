package controllers

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"myblog/models"
)

type AddPublishProductController struct {
	BaseController
}
type AddPublishParams struct {
	Id     string `json:"id"`
	Remarks   string `json:"remarks"`
	SubmitStatus int `json:"submit_status"`
	Title   string `json:"title"`
}



func (this *AddPublishProductController) Get() {
	res :=this.Input()
	submit_status,_    := this.GetInt("submit_status",0)
	//id,_   :=  this.GetInt("id",0)
	valid := validation.Validation{}
	params := AddPublishParams{
		Id       :    res.Get("id"),
		Remarks  :    res.Get("remarks"),
		SubmitStatus :submit_status,
		Title  :    res.Get("title"),
	}
	if err := this.ParseForm(&params); err != nil {
		//handle error
		return
	}
	valid.Required(params.Id,"id").Message("id不能为空")
	valid.Required(params.Remarks,"remarks").Message("备注不能为空")
	valid.Required(params.SubmitStatus,"submit_status").Message("submit_status不能为空")
	valid.Required(params.Title,"title").Message("title不能为空")
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			fmt.Println(err.Key, err.Message)
			this.ajaxList("faild",100,err.Message)
		}
	}
    //加入数据
     err :=models.AddAmazonListing(params.Id,params.Remarks,params.SubmitStatus,params.Title)
     if err !=nil {
     	fmt.Println(err)
	 }
	this.ajaxList("success",200,"")
}

//func (this *AddPublishProductController) addListing(){
//	//加入数据
//	 models.AddAmazonListing(1,"标注",1,"标题内容文件")
//	 this.ajaxList("success",200,"")
//}

//list := make([]map[string]interface{}, len(result))
//	for k, v := range result {
//		row := make(map[string]interface{})
//		row["id"] = v.Id
//		row["api_name"] = v.ApiName
//		row["api_url"] = v.ApiUrl
//		row["status_text"] = AUDIT_STATUS[v.Status]
//		row["status"] = v.Status
//		row["method"] = REQUEST_METHOD[v.Method]
//		sourceInfo := getSourceInfo(sourceList, v.SourceId)
//		row["source_name"] = sourceInfo.SourceName
//		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
//		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
//		list[k] = row
//	}



