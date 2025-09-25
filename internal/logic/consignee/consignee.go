package role

import (
	"goframe-shop-v2/internal/dao"
	"goframe-shop-v2/internal/model"
	"goframe-shop-v2/internal/model/entity"
	"goframe-shop-v2/internal/service"

	"golang.org/x/net/context"
)

type sConsignee struct{}

func init() {
	service.RegisterConsignee(New())
}

func New() *sConsignee {
	return &sConsignee{}
}

// GetList 查询内容列表
func (s *sConsignee) GetList(ctx context.Context, in model.ConsigneeGetListInput) (out *model.ConsigneeGetListOutput, err error) {
	var (
		m = dao.ConsigneeInfo.Ctx(ctx)
	)
	out = &model.ConsigneeGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.ConsigneeInfo
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
	//不指定item的键名用：Scan
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
