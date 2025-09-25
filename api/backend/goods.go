package backend

import "github.com/gogf/gf/v2/frame/g"

type GoodsCommonAddUpdate struct {
	Name             string `json:"name" v:"required#商品名称必填" dc:"商品名称"`
	Price            int    `json:"price" v:"required#价格必填" dc:"价格 单位分"`
	PicUrl           string `json:"pic_url" dc:"图片"`
	Level1CategoryId int    `json:"level1_category_id"  dc:"1级分类id"`
	Level2CategoryId int    `json:"level2_category_id"  dc:"2级分类id"`
	Level3CategoryId int    `json:"level3_category_id"  dc:"3级分类id"`
	Brand            string `json:"brand" v:"max-length:30#品牌名称最多30个字" dc:"品牌"`
	Stock            int    `json:"stock" dc:"库存"`
	Sale             int    `json:"sale" dc:"销量"`
	Tags             string `json:"tags" dc:"标签"`
	DetailInfo       string `json:"detail_info" dc:"商品详情"`
}

type GoodsReq struct {
	g.Meta `path:"/goods/add" tags:"商品" method:"post" summary:"添加商品接口"`
	GoodsCommonAddUpdate
}
type GoodsRes struct {
	Id uint `json:"id"`
}

type GoodsDeleteReq struct {
	g.Meta `path:"/goods/delete" method:"delete" tags:"商品" summary:"删除商品接口"`
	Id     uint `v:"min:1#请选择需要删除的商品id" dc:"商品id"`
}
type GoodsDeleteRes struct{}

type GoodsUpdateReq struct {
	g.Meta `path:"/goods/update" method:"post" tags:"商品" summary:"修改商品接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的商品id" dc:"商品Id"`
	GoodsCommonAddUpdate
}
type GoodsUpdateRes struct {
	Id uint `json:"id"`
}

type GoodsGetListCommonReq struct {
	g.Meta `path:"/goods/list" method:"get" tags:"商品" summary:"商品列表接口"`
	CommonPaginationReq
}
type GoodsGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
