package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type AddPraiseInput struct {
	UserId   uint  `json:"user_id"    dc:"用户id"`
	ObjectId uint  `json:"object_id"  dc:"对象id" v:"required#点赞的对象id必填"`
	Type     uint8 `json:"type"      dc:"点赞类型：1商品 2文章" v:"in:1,2"` //数据校验 范围约束
}
type AddPraiseOutput struct {
	Id uint `json:"id"`
}

type DeletePraiseInput struct {
	Id       uint  `json:"id"`
	UserId   uint  `json:"user_id"    dc:"用户id"`
	ObjectId uint  `json:"object_id"  dc:"对象id" v:"required#点赞的对象id必填"`
	Type     uint8 `json:"type"      dc:"点赞类型：1商品 2文章" v:"in:1,2"` //数据校验 范围约束
}
type DeletePraiseOutput struct {
	Id uint `json:"id"`
}

// PraiseListInput 获取内容列表
type PraiseListInput struct {
	Page int   // 分页号码
	Size int   // 分页数量，最大50
	Type uint8 //点赞类型：0所有 1商品 2文章
}

// PraiseListOutput 查询列表结果
type PraiseListOutput struct {
	List  []PraiseListOutputItem `json:"list" description:"列表"`
	Page  int                    `json:"page" description:"分页码"`
	Size  int                    `json:"size" description:"分页数量"`
	Total int                    `json:"total" description:"数据总数"`
}

type PraiseListOutputItem struct {
	Id        int         `json:"id"`
	UserId    int         `json:"user_id"    dc:"用户id"`
	ObjectId  int         `json:"object_id"  dc:"对象id"`
	Type      int         `json:"type"      dc:"点赞类型：1商品 2文章"`
	Goods     GoodsItem   `json:"goods" orm:"with:id=object_id"`
	Article   ArticleItem `json:"article" orm:"with:id=object_id"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}
