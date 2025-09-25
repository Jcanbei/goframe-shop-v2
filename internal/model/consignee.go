package model

import "github.com/gogf/gf/v2/os/gtime"

// ConsigneeGetListInput 获取内容列表
type ConsigneeGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// ConsigneeGetListOutput 查询列表结果
type ConsigneeGetListOutput struct {
	List  []ConsigneeGetListOutputItem `json:"list" description:"列表"`
	Page  int                          `json:"page" description:"分页码"`
	Size  int                          `json:"size" description:"分页数量"`
	Total int                          `json:"total" description:"数据总数"`
}

type ConsigneeGetListOutputItem struct {
	Id        uint        `json:"id"` // 自增ID
	UserId    int         `json:"userId"     dc:"用户id"`
	IsDefault int         `json:"is_default" dc:"默认地址1  非默认0"`
	Name      string      `json:"name"       dc:"收货人名字"`
	Phone     string      `json:"phone"      dc:"收货人手机号"`
	Province  string      `json:"province"   dc:"省"`
	City      string      `json:"city"   dc:"城市"`
	Town      string      `json:"town"   dc:"县区"`
	Street    int         `json:"street"     dc:"街道乡镇"`
	Detail    string      `json:"detail"     dc:"地址详情"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
	DeletedAt *gtime.Time `json:"deleted_at" `
}
