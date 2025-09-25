package user

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"goframe-shop-v2/internal/consts"
	"goframe-shop-v2/internal/dao"
	"goframe-shop-v2/internal/model"
	"goframe-shop-v2/internal/model/do"
	"goframe-shop-v2/internal/service"
	"goframe-shop-v2/utility"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// 注册
func (s *sUser) Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutput, err error) {
	//处理加密盐和密码的逻辑
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt

	// 获取当前时间并格式化为 MySQL 支持的标准格式
	now := gtime.Now().Format("Y-m-d H:i:s")

	// 插入数据时手动设置 created_at 和 updated_at 字段
	data := map[string]interface{}{
		"name":          in.Name,
		"avatar":        in.Avatar,
		"password":      in.Password,
		"user_salt":     in.UserSalt,
		"sex":           in.Sex,
		"status":        in.Status,
		"sign":          in.Sign,
		"secret_answer": in.SecretAnswer,
		"created_at":    now,
		"updated_at":    now,
	}

	//插入数据返回id
	lastInsertID, err := dao.UserInfo.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RegisterOutput{Id: uint(lastInsertID)}, err
}

// 修改密码
func (s *sUser) UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) (out model.UpdatePasswordOutput, err error) {
	// 验证密保问题
	userInfo := do.UserInfo{}
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.UserInfo.Ctx(ctx).WherePri(userId).Scan(&userInfo)
	if err != nil {
		return model.UpdatePasswordOutput{}, err
	}
	if gconv.String(userInfo.SecretAnswer) != in.SecretAnswer {
		return out, errors.New(consts.ErrSecretAnswerMsg)
	}
	userSalt := grand.S(10)
	in.UserSalt = userSalt
	in.Password = utility.EncryptPassword(in.Password, userSalt)

	// 获取当前时间并格式化为 MySQL 支持的标准格式
	now := gtime.Now().Format("Y-m-d H:i:s")

	_, err = dao.UserInfo.Ctx(ctx).WherePri(userId).Data(g.Map{
		"password":   in.Password,
		"user_salt":  in.UserSalt,
		"updated_at": now,
	}).Update()
	if err != nil {
		return model.UpdatePasswordOutput{}, err
	}
	return model.UpdatePasswordOutput{Id: userId}, nil
}
