package model

import "github.com/gogf/gf/v2/os/gtime"

// CouponCreateUpdateBase 创建/修改内容基类
type CouponCreateUpdateBase struct {
	Name       string //名称
	Price      int    // 优惠劵金额
	GoodsIds   string //可用的商品id
	CategoryId uint   //可用的商品分类
}

// CouponCreateInput 创建内容
type CouponCreateInput struct {
	CouponCreateUpdateBase
}

// CouponCreateOutput 创建内容返回结果
type CouponCreateOutput struct {
	CouponId uint `json:"coupon_id"`
}

// CouponUpdateInput 修改内容
type CouponUpdateInput struct {
	CouponCreateUpdateBase
	Id uint
}

// CouponGetListInput 获取内容列表
type CouponGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CouponGetListOutput 查询列表结果
type CouponGetListOutput struct {
	List  []CouponGetListOutputItem `json:"list" description:"列表"`
	Page  int                       `json:"page" description:"分页码"`
	Size  int                       `json:"size" description:"分页数量"`
	Total int                       `json:"total" description:"数据总数"`
}

type CouponGetListOutputItem struct {
	Id         uint        `json:"id"`          // 自增ID
	Name       string      `json:"name"`        //名称
	Price      int         `json:"price"`       // 优惠劵金额
	GoodsIds   string      `json:"goods_ids"`   //可用的商品id
	CategoryId uint        `json:"category_id"` //可用的商品分类
	CreatedAt  *gtime.Time `json:"created_at"`  // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at"`  // 修改时间
}

type CouponSearchOutputItem struct {
	CouponGetListOutputItem
}
