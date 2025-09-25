package backend

import "github.com/gogf/gf/v2/frame/g"

type GoodsOptionsCommonAddUpdate struct {
	GoodsId uint   `json:"goods_id" dc:"主商品id"`
	PicUrl  string `json:"pic_url"  dc:"图片"`
	Name    string `json:"name"     dc:"商品名称"   v:"required#名称必传"`
	Price   int    `json:"price"    dc:"价格 单位分" v:"required#价格必传"`
	Stock   int    `json:"stock"    dc:"库存"`
}

type GoodsOptionsReq struct {
	g.Meta `path:"/goods/options/add" tags:"商品规格" method:"post" summary:"添加商品规格接口"`
	GoodsOptionsCommonAddUpdate
}
type GoodsOptionsRes struct {
	Id uint `json:"id"`
}

type GoodsOptionsDeleteReq struct {
	g.Meta `path:"/goods/options/delete" method:"delete" tags:"商品规格" summary:"删除商品规格接口"`
	Id     uint `v:"min:1#请选择需要删除的商品规格id" dc:"商品规格id"`
}
type GoodsOptionsDeleteRes struct{}

type GoodsOptionsUpdateReq struct {
	g.Meta `path:"/goods/options/update" method:"post" tags:"商品规格" summary:"修改商品规格接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的商品规格id" dc:"商品规格Id"`
	GoodsOptionsCommonAddUpdate
}
type GoodsOptionsUpdateRes struct {
	Id uint `json:"id"`
}

type GoodsOptionsGetListCommonReq struct {
	g.Meta `path:"/goods/options/list" method:"get" tags:"商品规格" summary:"商品规格列表接口"`
	CommonPaginationReq
}
type GoodsOptionsGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
