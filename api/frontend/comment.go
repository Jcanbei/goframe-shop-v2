package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AddCommentReq struct {
	g.Meta   `path:"/add/comment" method:"post" tags:"前台评论" summary:"添加评论"`
	ParentId uint   `json:"parent_id" dc:"父级评论id"`
	ObjectId uint   `json:"object_id"  dc:"对象id" v:"required#评论的对象id必填"`
	Type     uint8  `json:"type"      dc:"评论类型：0所有 1商品 2文章" v:"in:1,2"` //数据校验 范围约束
	Content  string `json:"content" dc:"评论内容" v:"required#评论内容必填"`
}
type AddCommentRes struct {
	Id uint `json:"id"`
}

type DeleteCommentReq struct {
	g.Meta `path:"/delete/comment" method:"post" tags:"前台评论" summary:"移除评论"`
	Id     uint `json:"id"`
}
type DeleteCommentRes struct {
	Id uint `json:"id"`
}

type ListCommentReq struct {
	g.Meta `path:"/comment/list" method:"post" tags:"前台评论" summary:"评论列表"`
	Type   uint8 `json:"type" dc:"评论类型：0所有 1商品 2文章" v:"in:0,1,2"`
	CommonPaginationReq
}
type ListCommentRes struct {
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
	List  interface{} `json:"list" description:"列表"`
}

type ListCommentItem struct {
	Id       int         `json:"id"`
	UserId   int         `json:"user_id"    dc:"用户id"`
	ObjectId int         `json:"object_id"  dc:"对象id"`
	Type     int         `json:"type"      dc:"评论类型：0所有 1商品 2文章"`
	Goods    interface{} `json:"goods"`
	Article  interface{} `json:"article"`
}

type CommentBase struct {
	Id        int          `json:"id"        dc:""`
	ParentId  int          `json:"parent_id"  dc:"父级评论 id"`
	UserId    int          `json:"user_id"    dc:""`
	User      UserInfoBase `json:"user" dc:"用户信息"`
	ObjectId  int          `json:"object_id"  dc:""`
	Type      int          `json:"type"      dc:"评论类型：1商品 2文章"`
	Content   string       `json:"content"   dc:"评论内容"`
	CreatedAt *gtime.Time  `json:"created_at" dc:"创建时间"`
}
