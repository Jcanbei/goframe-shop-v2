package model

import "github.com/gogf/gf/v2/os/gtime"

// CategoryCreateUpdateBase 创建/修改内容基类
type CategoryCreateUpdateBase struct {
	ParentId uint   //父级id
	Name     string //名称
	PicUrl   string // 图片链接
	Level    uint8  //等级 默认1级分类
	Sort     uint8  //排序
}

// CategoryCreateInput 创建内容
type CategoryCreateInput struct {
	CategoryCreateUpdateBase
}

// CategoryCreateOutput 创建内容返回结果
type CategoryCreateOutput struct {
	CategoryId uint `json:"category_id"`
}

// CategoryUpdateInput 修改内容
type CategoryUpdateInput struct {
	CategoryCreateUpdateBase
	Id uint
}

// CategoryGetListInput 获取内容列表
type CategoryGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CategoryGetListOutput 查询列表结果
type CategoryGetListOutput struct {
	List  []CategoryGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}

type CategoryGetListOutputItem struct {
	Id        uint        `json:"id"`         // 自增ID
	ParentId  int         `json:"parent_id"`  //父级id
	Name      string      `json:"name"`       //名称
	PicUrl    string      `json:"pic_url"`    // 图片链接
	Level     uint8       `json:"level"`      //等级 默认1级分类
	Sort      uint8       `json:"sort"`       //排序
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type CategorySearchOutputItem struct {
	CategoryGetListOutputItem
}
