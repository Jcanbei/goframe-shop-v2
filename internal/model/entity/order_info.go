// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderInfo is the golang structure for table order_info.
type OrderInfo struct {
	Id               int         `json:"id"               orm:"id"                description:""`
	Number           string      `json:"number"           orm:"number"            description:"订单编号"`
	UserId           int         `json:"userId"           orm:"user_id"           description:"用户 id"`
	PayType          int         `json:"payType"          orm:"pay_type"          description:"支付方式：1 = 微信，2 = 支付宝"`
	Remark           string      `json:"remark"           orm:"remark"            description:"备注"`
	PayAt            *gtime.Time `json:"payAt"            orm:"pay_at"            description:"支付时间"`
	Status           string      `json:"status"           orm:"status"            description:"订单状态"`
	ConsigneeName    string      `json:"consigneeName"    orm:"consignee_name"    description:"收货人姓名"`
	ConsigneePhone   string      `json:"consigneePhone"   orm:"consignee_phone"   description:"收货人手机号"`
	ConsigneeAddress string      `json:"consigneeAddress" orm:"consignee_address" description:"收货人详细地址"`
	Price            int         `json:"price"            orm:"price"             description:"订单金额 单位分"`
	CouponPrice      int         `json:"couponPrice"      orm:"coupon_price"      description:"优惠劵金额 单位分"`
	ActualPrice      int         `json:"actualPrice"      orm:"actual_price"      description:"实际支付金额 单位分"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"        description:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"        description:""`
}
