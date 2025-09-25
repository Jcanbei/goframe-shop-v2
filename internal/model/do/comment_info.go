// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CommentInfo is the golang structure of table comment_info for DAO operations like Where/Data.
type CommentInfo struct {
	g.Meta    `orm:"table:comment_info, do:true"`
	Id        interface{} //
	ParentId  interface{} // 父级评论 id
	UserId    interface{} //
	ObjectId  interface{} //
	Type      interface{} // 评论类型：1商品 2文章
	Content   interface{} // 评论内容
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
