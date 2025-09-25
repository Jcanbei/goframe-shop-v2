// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-shop-v2/internal/model"
)

type (
	IComment interface {
		AddComment(ctx context.Context, in model.AddCommentInput) (out *model.AddCommentOutput, err error)
		DeleteComment(ctx context.Context, in model.DeleteCommentInput) (out *model.DeleteCommentOutput, err error)
		// 列表
		GetList(ctx context.Context, in model.CommentListInput) (out *model.CommentListOutput, err error)
	}
)

var (
	localComment IComment
)

func Comment() IComment {
	if localComment == nil {
		panic("implement not found for interface IComment, forgot register?")
	}
	return localComment
}

func RegisterComment(i IComment) {
	localComment = i
}
