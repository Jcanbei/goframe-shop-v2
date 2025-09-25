package backend

import "github.com/gogf/gf/v2/frame/g"

type CouponCommonAddUpdate struct {
	Name       string `json:"name" v:"required#优惠劵名称必填" dc:"优惠劵名称"`
	Price      int    `json:"price" v:"required#优惠劵金额" dc:"优惠劵金额"`
	GoodsIds   string `json:"goods_ids" dc:"可用的商品id 多个逗号分隔"`
	CategoryId uint   `json:"category_id" dc:"可用的商品分类"`
}

type CouponReq struct {
	g.Meta `path:"/coupon/add" tags:"优惠劵" method:"post" summary:"添加优惠劵接口"`
	CouponCommonAddUpdate
}
type CouponRes struct {
	CouponId uint `json:"coupon_id"`
}

type CouponDeleteReq struct {
	g.Meta `path:"/coupon/delete" method:"delete" tags:"优惠劵" summary:"删除优惠劵接口"`
	Id     uint `v:"min:1#请选择需要删除的优惠劵id" dc:"优惠劵id"`
}
type CouponDeleteRes struct{}

type CouponUpdateReq struct {
	g.Meta `path:"/coupon/update" method:"post" tags:"优惠劵" summary:"修改优惠劵接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的优惠劵id" dc:"优惠劵Id"`
	CouponCommonAddUpdate
}
type CouponUpdateRes struct {
	Id uint `json:"id"`
}

type CouponGetListCommonReq struct {
	g.Meta `path:"/coupon/list" method:"get" tags:"优惠劵" summary:"优惠劵列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type CouponGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
