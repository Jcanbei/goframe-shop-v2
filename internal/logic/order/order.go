package order

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop-v2/internal/consts"
	"goframe-shop-v2/internal/dao"
	"goframe-shop-v2/internal/model"
	"goframe-shop-v2/internal/service"
	"goframe-shop-v2/utility"
)

type sOrder struct{}

func init() {
	service.RegisterOrder(New())
}

func New() *sOrder {
	return &sOrder{}
}

// 下单
func (s *sOrder) Add(ctx context.Context, in model.OrderAddInput) (out *model.OrderAddOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	in.Number = utility.GetOrderNum()
	out = &model.OrderAddOutput{}

	// 获取当前时间，格式化为 MySQL 支持的 datetime 格式（不含毫秒）
	now := gtime.Now().Format("Y-m-d H:i:s")

	//官方建议的事务闭包处理
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//生成主订单
		orderData := g.Map{
			"number":            in.Number,
			"user_id":           in.UserId,
			"remark":            in.Remark,
			"consignee_name":    in.ConsigneeName,
			"consignee_phone":   in.ConsigneePhone,
			"consignee_address": in.ConsigneeAddress,
			"price":             in.Price,
			"coupon_price":      in.CouponPrice,
			"actual_price":      in.ActualPrice,
			"created_at":        now,
			"updated_at":        now,
		}

		lastInsertId, err := dao.OrderInfo.Ctx(ctx).TX(tx).Data(orderData).InsertAndGetId(in)
		if err != nil {
			return err
		}
		//生成商品订单
		for _, info := range in.OrderAddGoodsInfos {
			goodsData := g.Map{
				"order_id":         lastInsertId,
				"goods_id":         info.GoodsId,
				"goods_options_id": info.GoodsOptionsId,
				"count":            info.Count,
				"remark":           info.Remark,
				"price":            info.Price,
				"coupon_price":     info.CouponPrice,
				"actual_price":     info.ActualPrice,
				"created_at":       now,
				"updated_at":       now,
			}
			_, err := dao.OrderGoodsInfo.Ctx(ctx).TX(tx).Data(goodsData).Insert(info)
			if err != nil {
				return err
			}
		}
		//更新商品销量和库存
		for _, info := range in.OrderAddGoodsInfos {
			//商品增加销量
			_, err := dao.GoodsInfo.Ctx(ctx).TX(tx).WherePri(info.GoodsId).Increment(dao.GoodsInfo.Columns().Sale, info.Count)
			if err != nil {
				return err
			}
			//商品减少库存
			_, err2 := dao.GoodsInfo.Ctx(ctx).TX(tx).WherePri(info.GoodsId).Decrement(dao.GoodsInfo.Columns().Stock, info.Count)
			if err2 != nil {
				return err
			}
			//商品规格减少库存
			_, err3 := dao.GoodsOptionsInfo.Ctx(ctx).TX(tx).WherePri(info.GoodsOptionsId).Decrement(dao.GoodsOptionsInfo.Columns().Stock, info.Count)
			if err3 != nil {
				return err
			}
		}
		out.Id = uint(lastInsertId)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}
