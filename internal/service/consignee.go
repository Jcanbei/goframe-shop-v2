// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"goframe-shop-v2/internal/model"

	"golang.org/x/net/context"
)

type (
	IConsignee interface {
		// GetList 查询内容列表
		GetList(ctx context.Context, in model.ConsigneeGetListInput) (out *model.ConsigneeGetListOutput, err error)
	}
)

var (
	localConsignee IConsignee
)

func Consignee() IConsignee {
	if localConsignee == nil {
		panic("implement not found for interface IConsignee, forgot register?")
	}
	return localConsignee
}

func RegisterConsignee(i IConsignee) {
	localConsignee = i
}
