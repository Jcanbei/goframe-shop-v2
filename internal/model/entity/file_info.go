// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FileInfo is the golang structure for table file_info.
type FileInfo struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Name      string      `json:"name"      orm:"name"       description:"图片名称"`
	Src       string      `json:"src"       orm:"src"        description:"本地文件存储路径"`
	Url       string      `json:"url"       orm:"url"        description:"图片 URL"`
	UserId    int         `json:"userId"    orm:"user_id"    description:"上传用户 ID"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
