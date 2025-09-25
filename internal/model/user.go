package model

import "github.com/gogf/gf/v2/frame/g"

type RegisterInput struct {
	Name         string `json:"name"         dc:"用户名" v:"required#用户名必填"`
	Password     string `json:"password"     dc:"密码" v:"password"`
	Avatar       string `json:"avatar"       dc:"头像"`
	UserSalt     string `json:"userSalt"     dc:"加密盐 生成密码用"`
	Sex          int    `json:"sex"          dc:"1男 2女"`
	Status       int    `json:"status"       dc:"1正常 2拉黑冻结"`
	Sign         string `json:"sign"         dc:"个性签名"`
	SecretAnswer string `json:"secret_answer" dc:"密保问题的答案"`
}
type RegisterOutput struct {
	Id uint
}

type LoginInput struct {
	Name     string `json:"name"         dc:"用户名" v:"required#用户名必填"`
	Password string `json:"password"     dc:"密码" v:"password"`
}

type UpdatePasswordInput struct {
	Password     string `json:"password"     dc:"密码" v:"password"`
	UserSalt     string `json:"user_salt,omitempty"     dc:"加密盐 生成密码用"`
	SecretAnswer string `json:"secret_answer" dc:"密保问题的答案"`
}
type UpdatePasswordOutput struct {
	Id uint `json:"id"`
}

type UserInfoBase struct {
	g.Meta `orm:"table:user_info"`
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Sex    uint8  `json:"sex"`
	Sign   string `json:"sign"`
	Status uint8  `json:"status"`
}
