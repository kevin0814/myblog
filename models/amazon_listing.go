package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"myblog/utils"
	"strconv"
)
var db *gorm.DB
/**
 * amazon结构体
 */
//type AmanzonListing struct {
//	*Model
//	 id               int
//	 spu              string   `json:产品spu编码（商品中心获取)`
//	 product_sku      string   `json:商品中心sku`
//	 parent_sku       string `json:店铺id+站点二位简码（UK）+销售名简称（陆洲-LZ）+product_sku（多变种取第一个)`
//	 shop_id          int `json:店铺id`
//	 site             string `json:站点`
//	 category_id      string `json:最后一级分类id`
//	 category_name	  string `产品分类一级名称`
//	 product_cate_ids  string `json:产品分类id组合，用下划线分隔，例如：1_2_3`
//	 product_cate_name  string `json:产品分类值组合，以下划线间隔，例如：Video Games & Consoles_Video Game`
//	 sell_type          int `json:产品售卖形式，0：单品 1：变种`
//	 brand_biz_in       int `json:品牌商入驻，0：否 1：是`
//	 submit_status      int `json:产品上传状态0：未提交 1：已提交`
//	 relation_status	    int `父子关联状态0：未关联 1：已关联`
//	 relation_error_msg  string `json:上传关联失败信息`
//	 relation_error_code  int `json:上传关联失败码`
//	 upc  string `json:upc编码（amazon平台拉取）`
//	 product_state string `json:产品状态（amazon平台拉取）`
//	 product_state_desc string
//	 product_mfrs string
//	 mfr_part_number string
//	 product_attrs string
//	 other_attrs string
//	 product_price string
//	 prom_price string
//	 prom_start int64
//	 prom_end int64
//	 product_num int
//	 buy_price float32 `json:产品采购价(￥)人民币`
//	 pack_weight float32
//	 process_days int
//	 trans_template string
//	 pack_service int
//	 product_desc string
//	 key_words string
//	 is_plan_release int
//	 plan_update int64
//	 cat_id_one int
//	 cat_id_one_name string
//	 thumbnail string
//	 push_status int
//	 price_status int
//	 inventory_status int
//	 image_status int
//	 push_date float32
//	 title string
//	 asin string
//	 variant_theme string
//	 product_type_id string
//	 create_user string
//	 update_user string
//	 created_at int64
//	 updated_at int64
//	 remarks string
//	 target_audience string
//	 item_type string
//	 //Account  *Account `orm:rel(one)`
//}
type AmanzonListing struct {
	*Model
	Id         int `json:"id"`
	Spu        string `json:"spu"`
	Remarks string `json:"remarks"`
	SubmitStatus int  `json:"submit_status"`
	Title string `json:"title"`
	ShopId     int `json:"shop_Id"`
	Site       string `json:"site"`
	CreatedAt  int64 `json:"created_at"`

}

func (a AmanzonListing) TableName() string{
	return "amazon_listing"
}
type Amazon struct {
	ID int `json:"id"`
	ProductSku string `json:"product_sku"`
	ShopId int `json:"shop_Id"`
	Site  string `json:"site"`
	CreatedAt  int64 `json:"created_at"`
}
//请求参数
type CreateAmazonListingRequest  struct{
	Id          int `form:"id" binding:"required"`
	Spu         string `form:"spu"`
	//Title         string `form:"title" binding:"required,min=2,max=100"`
	//Desc          string `form:"desc" binding:"required,min=2,max=255"`
	//Content       string `form:"content" binding:"required,min=2,max=4294967295"`
	//CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
	//CreatedBy     string `form:"created_by" binding:"required,min=2,max=100"`
	//State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type Account struct{
	shop_Id           int
	shop_name         string
}
//查询单条记录
func GetAmazonListing(param CreateAmazonListingRequest) (*Amazon,error){
	 sql :="select id,product_sku,shop_id,site,created_at from amazon_listing where id ="+ strconv.Itoa(param.Id)
	 row :=utils.QueryRowDB(sql)
	 id := 0
	 product_sku := ""
	 shop_id := 0
	 site := ""
	 var created_at int64
	 created_at = 0
	 row.Scan(&id, &product_sku, &shop_id, &site, &created_at)
	 amazon := Amazon{id, product_sku, shop_id, site, created_at}
	return &Amazon{
		ID:                 amazon.ID,
		ProductSku:         amazon.ProductSku,
		ShopId:             amazon.ShopId,
		Site:               amazon.Site,
		CreatedAt:          amazon.CreatedAt,
	}, nil
}
//查询优化
func (a AmanzonListing) Select(db *gorm.DB) (*Amazon,error){
	if a.Id != 0 {
		db =db.Where("id =?",a.Id)
	}
	if a.Spu !=""{
		db= db.Where("spu = ?",a.Spu)
	}
  return &Amazon{},nil
	//if err := db.Model(&a).FirstOrCreate(&count).Error; err != nil {
	//	return 0, err
	//}
}
type Role struct {
	Id     string `json:"id"`
	Remarks   string `json:"remarks"`
	SubmitStatus int `json:"submit_status"`
	Title   string `json:"title"`
}
func AddAmazonListing(id string,remark string, submitStatus int,title string) error {
	//db.Create(&AmanzonListing{remarks: "L1212", title:"标题需要注意内容贴切"})
	//amazon := &AmanzonListing{
	//	Remarks: remark,
	//	SubmitStatus:  submitStatus,
	//	Title: title,
	//}
	amazon :=map[string]interface{}{
		"remarks":remark,
		"submit_status":submitStatus,
		"title":title,
	}
	fmt.Println(amazon)
	fmt.Println(id)
	if err :=db.Model(&AmanzonListing{}).Where("id= ? ", id).Limit(1).Updates(amazon).Error; err != nil {
		return err
	}
	return nil
}

////-------插入sku---------------
//func InsertAmanzonListing(AmanzonListing AmanzonListing) (int64, error) {
//	return utils.ModifyDB("insert into amanzon_listing(filepath,filename,status,createtime)values(?,?,?,?)",
//		AmanzonListing.product_sku, AmanzonListing.variant_theme, AmanzonListing.product_type_id, AmanzonListing.update_user)
//}



//func (d *Dao) CreateArticle(param *Articles) (*AmanzonListing, error) {
//	amanzonListing := AmanzonListing{
//		id:            param.Title,
//		Desc:          param.Desc,
//		Content:       param.Content,
//		CoverImageUrl: param.CoverImageUrl,
//		State:         param.State,
//		Model:         &model.Model{CreatedBy: param.CreatedBy},
//	}
//	return amanzonListing.Create(d.engine)
//}

////--------查询图片----------
//func FindAllAlbums() ([]Album, error) {
//	rows, err := utils.QueryDB("select id,filepath,filename,status,createtime from album")
//	if err != nil {
//		return nil, err
//	}
//	var albums []Album
//	for rows.Next() {
//		id := 0
//		filepath := ""
//		filename := ""
//		status := 0
//		var createtime int64
//		createtime = 0
//		rows.Scan(&id, &filepath, &filename, &status, &createtime)
//		album := Album{id, filepath, filename, status, createtime}
//		albums = append(albums, album)
//	}
//	return albums, nil
//}
