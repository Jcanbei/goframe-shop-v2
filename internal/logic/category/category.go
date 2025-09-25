package category

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-shop-v2/internal/dao"
	"goframe-shop-v2/internal/model"
	"goframe-shop-v2/internal/model/entity"
	"goframe-shop-v2/internal/service"
)

type sCategory struct{}

func init() {
	service.RegisterCategory(New())
}

func New() *sCategory {
	return &sCategory{}
}

func (s *sCategory) Create(ctx context.Context, in model.CategoryCreateInput) (out model.CategoryCreateOutput, err error) {
	// 获取当前时间并格式化为 MySQL 支持的标准格式
	now := gtime.Now().Format("Y-m-d H:i:s")

	// 插入数据时手动设置 created_at 和 updated_at 字段
	data := map[string]interface{}{
		"parent_id":  in.ParentId,
		"name":       in.Name,
		"pic_url":    in.PicUrl,
		"level":      in.Level,
		"sort":       in.Sort,
		"created_at": now,
		"updated_at": now,
	}

	lastInsertID, err := dao.CategoryInfo.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.CategoryCreateOutput{CategoryId: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sCategory) Delete(ctx context.Context, id uint) (err error) {
	// 获取当前时间并格式化为 MySQL 支持的标准格式
	now := gtime.Now().Format("Y-m-d H:i:s")
	// 删除内容
	_, err = dao.CategoryInfo.Ctx(ctx).Where(g.Map{
		dao.CategoryInfo.Columns().Id: id,
	}).Data(g.Map{
		dao.CategoryInfo.Columns().DeletedAt: now,
	}).Update()
	if err != nil {
		return err
	}
	return
}

// Update 修改
func (s *sCategory) Update(ctx context.Context, in model.CategoryUpdateInput) error {
	// 获取当前时间并格式化为 MySQL 支持的标准格式
	now := gtime.Now().Format("Y-m-d H:i:s")

	// 插入数据时手动设置 created_at 和 updated_at 字段
	data := g.Map{
		"parent_id":  in.ParentId,
		"name":       in.Name,
		"pic_url":    in.PicUrl,
		"level":      in.Level,
		"sort":       in.Sort,
		"updated_at": now,
	}

	_, err := dao.CategoryInfo.
		Ctx(ctx).
		Data(data).
		FieldsEx(dao.CategoryInfo.Columns().Id).
		Where(dao.CategoryInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询商品分类列表
func (s *sCategory) GetList(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error) {
	var (
		m = dao.CategoryInfo.Ctx(ctx)
	)
	out = &model.CategoryGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.OrderDesc(dao.CategoryInfo.Columns().Sort)
	// 执行查询
	var list []*entity.CategoryInfo
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

// GetListAll 查询商品分类列表全部信息(不分页)
func (s *sCategory) GetListAll(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error) {
	var (
		m = dao.CategoryInfo.Ctx(ctx)
	)
	out = &model.CategoryGetListOutput{}

	listModel := m
	// 排序方式
	listModel = listModel.OrderDesc(dao.CategoryInfo.Columns().Sort)
	// 执行查询
	var list []*entity.CategoryInfo
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
