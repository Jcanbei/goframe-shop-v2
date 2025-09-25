package frontend

import "github.com/gogf/gf/v2/frame/g"

type AddCollectionReq struct {
	g.Meta   `path:"/add/collection" method:"post" tags:"前台收藏" summary:"添加收藏"`
	ObjectId uint  `json:"object_id"  dc:"对象id" v:"required#收藏的对象id必填"`
	Type     uint8 `json:"type"      dc:"收藏类型：0所有 1商品 2文章" v:"in:1,2"` //数据校验 范围约束
}
type AddCollectionRes struct {
	Id uint `json:"id"`
}

type DeleteCollectionReq struct {
	g.Meta   `path:"/delete/collection" method:"post" tags:"前台收藏" summary:"移除收藏"`
	Id       uint  `json:"id"`
	ObjectId uint  `json:"object_id"`
	Type     uint8 `json:"type"`
}
type DeleteCollectionRes struct {
	Id uint `json:"id"`
}

type ListCollectionReq struct {
	g.Meta `path:"/collection/list" method:"post" tags:"前台收藏" summary:"收藏列表"`
	Type   uint8 `json:"type" dc:"收藏类型：0所有 1商品 2文章" v:"in:0,1,2"`
	CommonPaginationReq
}
type ListCollectionRes struct {
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
	List  interface{} `json:"list" description:"列表"`
}

type ListCollectionItem struct {
	Id       int         `json:"id"`
	UserId   int         `json:"user_id"    dc:"用户id"`
	ObjectId int         `json:"object_id"  dc:"对象id"`
	Type     int         `json:"type"      dc:"收藏类型：0所有 1商品 2文章"`
	Goods    interface{} `json:"goods"`
	Article  interface{} `json:"article"`
}
