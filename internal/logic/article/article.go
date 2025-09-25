package user_coupon

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-shop-v2/internal/dao"
	"goframe-shop-v2/internal/model"
	"goframe-shop-v2/internal/model/entity"
	"goframe-shop-v2/internal/service"
)

type sArticle struct{}

func init() {
	service.RegisterArticle(New())
}

func New() *sArticle {
	return &sArticle{}
}

func (s *sArticle) Create(ctx context.Context, in model.ArticleCreateInput) (out model.ArticleCreateOutput, err error) {
	// 获取当前时间并格式化为 MySQL 支持的标准格式
	now := gtime.Now().Format("Y-m-d H:i:s")

	// 插入数据时手动设置 created_at 和 updated_at 字段
	data := map[string]interface{}{
		"user_id":    in.UserId,
		"title":      in.Title,
		"desc":       in.Desc,
		"pic_url":    in.PicUrl,
		"is_admin":   in.IsAdmin,
		"detail":     in.Detail,
		"praise":     in.Praise,
		"created_at": now,
		"updated_at": now,
	}

	lastInsertID, err := dao.ArticleInfo.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.ArticleCreateOutput{Id: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sArticle) Delete(ctx context.Context, id uint) (err error) {
	// 获取当前时间并格式化为 MySQL 支持的标准格式
	now := gtime.Now().Format("Y-m-d H:i:s")
	// 删除内容
	_, err = dao.ArticleInfo.Ctx(ctx).Where(g.Map{
		dao.ArticleInfo.Columns().Id: id,
	}).Data(g.Map{
		dao.ArticleInfo.Columns().DeletedAt: now,
		dao.ArticleInfo.Columns().UpdatedAt: now,
	}).Update()
	if err != nil {
		return err
	}
	return
}

// Update 修改
func (s *sArticle) Update(ctx context.Context, in model.ArticleUpdateInput) error {
	// 获取当前时间并格式化为 MySQL 支持的标准格式
	now := gtime.Now().Format("Y-m-d H:i:s")

	// 插入数据时手动设置 created_at 和 updated_at 字段
	data := g.Map{
		"user_id":    in.UserId,
		"title":      in.Title,
		"desc":       in.Desc,
		"pic_url":    in.PicUrl,
		"is_admin":   in.IsAdmin,
		"detail":     in.Detail,
		"praise":     in.Praise,
		"updated_at": now,
	}

	_, err := dao.ArticleInfo.
		Ctx(ctx).
		Data(data).
		FieldsEx(dao.ArticleInfo.Columns().Id).
		Where(dao.ArticleInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询文章列表
func (s *sArticle) GetList(ctx context.Context, in model.ArticleGetListInput) (out *model.ArticleGetListOutput, err error) {
	var (
		m = dao.ArticleInfo.Ctx(ctx)
	)
	out = &model.ArticleGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 执行查询
	var list []*entity.ArticleInfo
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

// 详情
func (s *sArticle) Detail(ctx context.Context, in model.ArticleDetailInput) (out *model.ArticleDetailOutput, err error) {
	err = dao.ArticleInfo.Ctx(ctx).WithAll().WherePri(in.Id).Scan(&out)
	return
}
