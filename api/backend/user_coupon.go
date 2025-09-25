package backend

import "github.com/gogf/gf/v2/frame/g"

type UserCouponCommonAddUpdate struct {
	UserId   uint  `json:"user_id" v:"required#用户id必填" dc:"用户id"`
	CouponId uint  `json:"coupon_id" v:"required#优惠劵id必填" dc:"优惠劵id"`
	Status   uint8 `json:"status" dc:"状态：1可用 2已用 3过期"`
}

type UserCouponReq struct {
	g.Meta `path:"/user/coupon/add" tags:"用户优惠劵" method:"post" summary:"添加用户优惠劵接口"`
	UserCouponCommonAddUpdate
}
type UserCouponRes struct {
	Id uint `json:"id"`
}

type UserCouponDeleteReq struct {
	g.Meta `path:"/user/coupon/delete" method:"delete" tags:"用户优惠劵" summary:"删除用户优惠劵接口"`
	Id     uint `v:"min:1#请选择需要删除的用户优惠劵id" dc:"用户优惠劵id"`
}
type UserCouponDeleteRes struct{}

type UserCouponUpdateReq struct {
	g.Meta `path:"/user/coupon/update" method:"post" tags:"用户优惠劵" summary:"修改用户优惠劵接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的用户优惠劵id" dc:"用户优惠劵Id"`
	UserCouponCommonAddUpdate
}
type UserCouponUpdateRes struct {
	Id uint `json:"id"`
}

type UserCouponGetListCommonReq struct {
	g.Meta `path:"/user/coupon/list" method:"get" tags:"用户优惠劵" summary:"用户优惠劵列表接口"`
	CommonPaginationReq
}
type UserCouponGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
