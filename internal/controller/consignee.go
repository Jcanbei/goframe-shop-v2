package controller

import (
	"goframe-shop-v2/api/backend"
	"goframe-shop-v2/internal/model"
	"goframe-shop-v2/internal/service"

	"golang.org/x/net/context"
)

// Consignee 角色管理
var Consignee = cConsignee{}

type cConsignee struct{}

func (c *cConsignee) List(ctx context.Context, req *backend.ConsigneeListReq) (res *backend.ConsigneeListRes, err error) {
	getListRes, err := service.Consignee().GetList(ctx, model.ConsigneeGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &backend.ConsigneeListRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
