package user_coupon

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop-v2/internal/consts"
	"goframe-shop-v2/internal/dao"
	"goframe-shop-v2/internal/logic/collection"
	"goframe-shop-v2/internal/model"
	"goframe-shop-v2/internal/model/entity"
	"goframe-shop-v2/internal/service"
)

type sGoods struct{}

func init() {
	service.RegisterGoods(New())
}

func New() *sGoods {
	return &sGoods{}
}

func (s *sGoods) Create(ctx context.Context, in model.GoodsCreateInput) (out model.GoodsCreateOutput, err error) {
	// 获取当前时间并格式化为 MySQL 支持的标准格式
	now := gtime.Now().Format("Y-m-d H:i:s")

	// 插入数据时手动设置 created_at 和 updated_at 字段
	data := map[string]interface{}{
		"pic_url":            in.PicUrl,
		"name":               in.Name,
		"price":              in.Price,
		"level1_category_id": in.Level1CategoryId,
		"level2_category_id": in.Level2CategoryId,
		"level3_category_id": in.Level3CategoryId,
		"brand":              in.Brand,
		"stock":              in.Stock,
		"sale":               in.Sale,
		"tags":               in.Tags,
		"detail_info":        in.DetailInfo,
		"created_at":         now,
		"updated_at":         now,
	}

	lastInsertID, err := dao.GoodsInfo.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.GoodsCreateOutput{Id: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sGoods) Delete(ctx context.Context, id uint) (err error) {
	// 获取当前时间并格式化为 MySQL 支持的标准格式
	now := gtime.Now().Format("Y-m-d H:i:s")
	// 删除内容
	_, err = dao.GoodsInfo.Ctx(ctx).Where(g.Map{
		dao.GoodsInfo.Columns().Id: id,
	}).Data(g.Map{
		dao.GoodsInfo.Columns().DeletedAt: now,
		dao.GoodsInfo.Columns().UpdatedAt: now,
	}).Update()
	if err != nil {
		return err
	}
	return
}

// Update 修改
func (s *sGoods) Update(ctx context.Context, in model.GoodsUpdateInput) error {
	// 获取当前时间并格式化为 MySQL 支持的标准格式
	now := gtime.Now().Format("Y-m-d H:i:s")

	// 插入数据时手动设置 created_at 和 updated_at 字段
	data := g.Map{
		"pic_url":            in.PicUrl,
		"name":               in.Name,
		"price":              in.Price,
		"level1_category_id": in.Level1CategoryId,
		"level2_category_id": in.Level2CategoryId,
		"level3_category_id": in.Level3CategoryId,
		"brand":              in.Brand,
		"stock":              in.Stock,
		"sale":               in.Sale,
		"tags":               in.Tags,
		"detail_info":        in.DetailInfo,
		"updated_at":         now,
	}

	_, err := dao.GoodsInfo.
		Ctx(ctx).
		Data(data).
		FieldsEx(dao.GoodsInfo.Columns().Id).
		Where(dao.GoodsInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询优惠劵列表
func (s *sGoods) GetList(ctx context.Context, in model.GoodsGetListInput) (out *model.GoodsGetListOutput, err error) {
	var (
		m = dao.GoodsInfo.Ctx(ctx)
	)
	out = &model.GoodsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 执行查询
	var list []*entity.GoodsInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// 商品详情
func (s *sGoods) Detail(ctx context.Context, in model.GoodsDetailInput) (out model.GoodsDetailOutput, err error) {
	err = dao.GoodsInfo.Ctx(ctx).WithAll().WherePri(in.Id).Scan(&out)
	if err != nil {
		return model.GoodsDetailOutput{}, err
	}
	out.IsCollect, err = collection.CheckIsCollect(ctx, model.CheckIsCollectInput{
		UserId:   gconv.Uint(ctx.Value(consts.CtxUserId)),
		ObjectId: in.Id,
		Type:     consts.CollectionTypeGoods,
	})
	return
}
