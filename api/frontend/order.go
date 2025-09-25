package frontend

import "github.com/gogf/gf/v2/frame/g"

type AddOrderReq struct {
	g.Meta `path:"/add/order" method:"post" tags:"前台订单" summary:"创建订单"`
	//主订单维度
	Price            int    `json:"price"            dc:"订单金额 单位分"`
	CouponPrice      int    `json:"coupon_price"      dc:"优惠劵金额 单位分"`
	ActualPrice      int    `json:"actual_price"      dc:"实际支付金额 单位分"`
	ConsigneeName    string `json:"consignee_name"    dc:"收货人姓名"`
	ConsigneePhone   string `json:"consignee_phone"   dc:"收货人手机号"`
	ConsigneeAddress string `json:"consignee_address" dc:"收货人详细地址"`
	Remark           string `json:"remark"           dc:"备注"`
	//商品订单维度
	OrderAddGoodsInfos []*OrderAddGoodsInfo `json:"order_add_goods_infos"`
}
type AddOrderRes struct {
	Id uint `json:"id"`
}

type OrderAddGoodsInfo struct {
	GoodsId        int    `json:"goods_id"         dc:"商品id"`
	GoodsOptionsId int    `json:"goods_options_id" dc:"商品规格id sku id"`
	Count          int    `json:"count"            dc:"商品数量"`
	Remark         string `json:"remark"           dc:"备注"`
	Price          int    `json:"price"            dc:"订单金额 单位分"`
	CouponPrice    int    `json:"coupon_price"     dc:"优惠券金额 单位分"`
	ActualPrice    int    `json:"actual_price"     dc:"实际支付金额 单位分"`
}
