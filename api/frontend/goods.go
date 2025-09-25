package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GoodsGetListCommonReq struct {
	g.Meta `path:"/goods/list" method:"post" tags:"前台商品" summary:"商品列表接口"`
	CommonPaginationReq
}
type GoodsGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type GoodsDetailReq struct {
	g.Meta `path:"/goods/detail" method:"post" tags:"前台商品" summary:"商品详情接口"`
	Id     uint `json:"id"`
}
type GoodsDetailRes struct {
	GoodsInfoBase
	Options   []GoodsOptionsBase `json:"options"` //规格 sku
	Comments  []CommentBase      `json:"comment"`
	IsCollect bool               `json:"is_collect"`
}

type GoodsInfoBase struct {
	Id               int         `json:"id"               dc:""`
	PicUrl           string      `json:"pic_url"           dc:"图片"`
	Name             string      `json:"name"             dc:"商品名称"`
	Price            int         `json:"price"            dc:"价格 单位分"`
	Level1CategoryId int         `json:"level1_category_id" dc:"1级分类id"`
	Level2CategoryId int         `json:"level2_category_id" dc:"2级分类id"`
	Level3CategoryId int         `json:"level3_category_id" dc:"3级分类id"`
	Brand            string      `json:"brand"            dc:"品牌"`
	Stock            int         `json:"stock"            dc:"库存"`
	Sale             int         `json:"sale"             dc:"销量"`
	Tags             string      `json:"tags"             dc:"标签"`
	DetailInfo       string      `json:"detail_info"       dc:"商品详情"`
	CreatedAt        *gtime.Time `json:"created_at"        dc:""`
}

type GoodsOptionsBase struct {
	Id        int         `json:"id"        dc:""`
	GoodsId   int         `json:"goods_id"   dc:"商品id"`
	PicUrl    string      `json:"pic_url"    dc:"图片"`
	Name      string      `json:"name"      dc:"商品名称"`
	Price     int         `json:"price"     dc:"价格 单位分"`
	Stock     int         `json:"stock"     dc:"库存"`
	CreatedAt *gtime.Time `json:"created_at" dc:""`
}
