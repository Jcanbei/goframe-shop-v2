package model

type OrderAddInput struct {
	//主订单维度
	UserId           uint
	Number           string
	Price            int    `dc:"订单金额 单位分"`
	CouponPrice      int    `dc:"优惠劵金额 单位分"`
	ActualPrice      int    `dc:"实际支付金额 单位分"`
	ConsigneeName    string `dc:"收货人姓名"`
	ConsigneePhone   string `dc:"收货人手机号"`
	ConsigneeAddress string `dc:"收货人详细地址"`
	Remark           string `dc:"备注"`
	//商品订单维度
	OrderAddGoodsInfos []*OrderAddGoodsInfo
}

type OrderAddGoodsInfo struct {
	Id             int
	OrderId        int    `dc:"关联的主订单表id"`
	GoodsId        int    `dc:"商品id"`
	GoodsOptionsId int    `dc:"商品规格id sku id"`
	Count          int    `dc:"商品数量"`
	Remark         string `dc:"备注"`
	Price          int    `dc:"订单金额 单位分"`
	CouponPrice    int    `dc:"优惠券金额 单位分"`
	ActualPrice    int    `dc:"实际支付金额 单位分"`
}

type OrderAddOutput struct {
	Id uint `json:"id"`
}
