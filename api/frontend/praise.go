package frontend

import "github.com/gogf/gf/v2/frame/g"

type AddPraiseReq struct {
	g.Meta   `path:"/add/praise" method:"post" tags:"前台点赞" summary:"添加点赞"`
	ObjectId uint  `json:"object_id"  dc:"对象id" v:"required#点赞的对象id必填"`
	Type     uint8 `json:"type"      dc:"点赞类型：0所有 1商品 2文章" v:"in:1,2"` //数据校验 范围约束
}
type AddPraiseRes struct {
	Id uint `json:"id"`
}

type DeletePraiseReq struct {
	g.Meta   `path:"/delete/praise" method:"post" tags:"前台点赞" summary:"移除点赞"`
	Id       uint  `json:"id"`
	ObjectId uint  `json:"object_id"`
	Type     uint8 `json:"type"`
}
type DeletePraiseRes struct {
	Id uint `json:"id"`
}

type ListPraiseReq struct {
	g.Meta `path:"/praise/list" method:"post" tags:"前台点赞" summary:"点赞列表"`
	Type   uint8 `json:"type" dc:"点赞类型：0所有 1商品 2文章" v:"in:0,1,2"`
	CommonPaginationReq
}
type ListPraiseRes struct {
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
	List  interface{} `json:"list" description:"列表"`
}

type ListPraiseItem struct {
	Id       int         `json:"id"`
	UserId   int         `json:"user_id"    dc:"用户id"`
	ObjectId int         `json:"object_id"  dc:"对象id"`
	Type     int         `json:"type"      dc:"点赞类型：0所有 1商品 2文章"`
	Goods    interface{} `json:"goods"`
	Article  interface{} `json:"article"`
}
