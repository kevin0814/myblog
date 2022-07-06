package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"io"
	"myblog/models"
)

type PrePublishProductController struct {
	BaseController
}
type PrePublishParams struct {
	spu            string
	sku_name       string
	is_bulky       string
	cat_id_one     int
	shop_id        int
	is_publish     int
	is_amazon      int
	distri_status  int
	page_num       int
	page_size      int
	dept_id        []int
}



func (this *PrePublishProductController) Get() {
	res :=this.Input()
	cat_id_one,_   :=  this.GetInt("cat_id_one",0)
	shop_id,_      :=  this.GetInt("shop_id",0)
	is_amazon,_    :=  this.GetInt("is_amazon",0)
	page_size,_    :=  this.GetInt("page_size",1)
	page_num,_     :=   this.GetInt("page_num",10)
	distri_status,_    := this.GetInt("distri_status",0)
	is_publish,_    := this.GetInt("is_publish",0)
	id,_   :=  this.GetInt("id",0)
	valid := validation.Validation{}
	params := PrePublishParams{
		spu       :    res.Get("spu"),
		sku_name  :    res.Get("sku_name"),
		is_bulky  :    res.Get("is_bulky"),
		cat_id_one:    cat_id_one,
		shop_id   :shop_id,
		is_publish:is_publish,
		is_amazon:is_amazon,
		distri_status:distri_status,
		page_num:page_num,
		page_size:page_size,
		dept_id:[]int{1,3},
	}
	if err := this.ParseForm(&params); err != nil {
		//handle error
		return
	}
	valid.Required(params.spu,"spu").Message("spu不能为空")
	valid.Required(params.sku_name,"sku_name").Message("sku_name不能为空")
	valid.Required(params.is_bulky,"is_bulky").Message("is_bulky不能为空")
	valid.Required(params.is_bulky,"cat_id_one").Message("cat_id_one不能为空")
	valid.Required(params.is_bulky,"shop_id").Message("shop_id不能为空")
	valid.Required(params.is_bulky,"is_publish").Message("is_publish不能为空")
	valid.Required(params.is_bulky,"is_amazon").Message("is_amazon不能为空")
	valid.Required(params.is_bulky,"distri_status").Message("distri_status不能为空")
	valid.Required(params.is_bulky,"dept_id").Message("地址不能为空")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			fmt.Println(err.Key, err.Message)
			//message := err.Key + err.Message
			//this.Ctx.WriteString(message)
			this.ajaxList("faild",100,err.Message)
		}
	}
	param := models.CreateAmazonListingRequest{Id:id}
    //加入数据
    amazon, err:= models.GetAmazonListing(param)
    if err !=nil {
		this.ajaxList("faild",100,err)
	}
	this.ajaxList("success",200,amazon)
	this.Data["json"] = map[string]interface{}{"code": 1, "message": amazon}
   //data : map[string]interface{}{
   //
   //}


	o := orm.NewOrm()
	o.Using("default") //使用测试环境
	var w io.Writer //io写入
	orm.DebugLog = orm.NewLog(w) //获取
	//idStr := this.Ctx.Input.Param(":id")
	//
	//id, _ := strconv.Atoi(idStr)
	//fmt.Println("id:", id)
	//o.Raw("insert into usertest(id,Name)values(?,?)",1,"keuin")
    //userTest := UserTest{}
	//beego.Info(userTest)
    //_,err :=o.Insert(&userTest)
    //if err !=nil {
    //	fmt.Println(err)
	//}
	//fmt.Println(id)
	//AmanzonListing :=models.AmanzonListing{id:322}
	//err =o.Read(&AmanzonListing)
	//user := new(models.AmanzonListing)
	//user.spu ="WJWHP0001"
	//fmt.Println(user)
	//user.shop_name = "11"
	//fmt.Println(user)
	 data :=map[string]interface{}{
	 	"ok":111,"ssss":"222222",
	 }

	 this.ajaxList("success",200,data)
}

//func (this *PrePublishProductController) addListing(){
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



