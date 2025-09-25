// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderInfo is the golang structure of table order_info for DAO operations like Where/Data.
type OrderInfo struct {
	g.Meta           `orm:"table:order_info, do:true"`
	Id               interface{} //
	Number           interface{} // 订单编号
	UserId           interface{} // 用户 id
	PayType          interface{} // 支付方式：1 = 微信，2 = 支付宝
	Remark           interface{} // 备注
	PayAt            *gtime.Time // 支付时间
	Status           interface{} // 订单状态
	ConsigneeName    interface{} // 收货人姓名
	ConsigneePhone   interface{} // 收货人手机号
	ConsigneeAddress interface{} // 收货人详细地址
	Price            interface{} // 订单金额 单位分
	CouponPrice      interface{} // 优惠劵金额 单位分
	ActualPrice      interface{} // 实际支付金额 单位分
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time //
}
