package controllers

import (
	"fmt"
)

type TestController struct {
	BaseController
}
//type Data struct {
//	ID    int `json:"id"`
//	IP    string `json:"IP"`
//	DESC  string `json:"desc"`
//	OWNER string `json:"owner"`
//}
//
//type CloumnsData struct {
//	NAME  string `json:"name"`
//	ALIAS string `json:"alias"`
//}
//
//type Employee struct {
//	CODE    int `json:"code"`
//	ISADMIN bool `json:"isadmin"`
//	DATA    []Data `json:"data"`
//	COLUMNS []CloumnsData `json:"columns"`
//	MESSGAE string `json:"messgae"`
//}

//type Response struct {
//	code    int                `json:"code"`
//	message string                 `json:"message"`
//	data    map[string]interface{} `json:"data"`
//}


func (this *TestController) Get() {
	res := make(map[string]interface{})
	res["code"] = 200
	res["msg"]  = "success"
	res["data"] = map[string]interface{}{
		"username" : "Tom",
		"age"      : 30,
		"hobby"    : []string{"读书","爬山"},
	}

	fmt.Println("map data :", res)
	//re :=Response{code:200,message:"成功",data:map[string]interface{}{"aaa":"bbb","ccc":"dddd"}}
	//this.Data["json"]=re
	this.Data["json"] =res
	this.ServeJSON()
	// json.Marshal(Response{"code":200,"message":"验证成功","data":{"a":1,"b":3}})
	//fmt.Println(string(str))
	//fmt.Println(resp)
	//this.Data["json"] =resp
	//this.ServeJSON()
	//resp := Response{}
	//err := json.Unmarshal(str, &resp)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(resp)
	//this.Data["json"] =resp
	//this.ServeJSON()

	//url := "http://ops-environment.com/channel/ip/v1"
	//resp, _ := http.Get(url)
	//s := Employee{}
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(body)
	//resp.Body.Close()
	//json.Unmarshal([]byte(body), &s)
	//fmt.Println(fmt.Sprintf("%+v",s))

	//id := models.QueryArticleWithParam("id")
	//fmt.Println(models.HandleTagsListData(id))
	//this.Data["json"] = models.HandleTagsListData(id)
	//this.ServeJSON()
}
