package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta       `path:"/register" method:"post" tags:"前台用户" summary:"用户注册接口"`
	Name         string `json:"name"         dc:"用户名" v:"required#用户名必填"`
	Password     string `json:"password"     dc:"密码" v:"password"`
	Avatar       string `json:"avatar"       dc:"头像"`
	UserSalt     string `json:"user_salt"     dc:"加密盐 生成密码用"`
	Sex          int    `json:"sex"          dc:"1男 2女"`
	Status       int    `json:"status"       dc:"1正常 2拉黑冻结"`
	Sign         string `json:"sign"         dc:"个性签名"`
	SecretAnswer string `json:"secret_answer" dc:"密保问题的答案"`
}
type RegisterRes struct {
	Id uint `json:"id"`
}

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"前台用户" summary:"用户登录接口"`
	Name     string `json:"name"         dc:"用户名" v:"required#用户名必填"`
	Password string `json:"password"     dc:"密码" v:"password"`
}

// for gtoken
type LoginRes struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
	UserInfoBase
}

type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"前台用户" summary:"当前登录用户信息"`
}
type UserInfoRes struct {
	UserInfoBase
}

type UserInfoBase struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Sex    uint8  `json:"sex"`
	Sign   string `json:"sign"`
	Status uint8  `json:"status"`
}

// 修改密码
type UpdatePasswordReq struct {
	g.Meta       `path:"/update/password" method:"post" tags:"前台用户" summary:"修改密码"`
	Password     string `json:"password"     dc:"密码" v:"password"`
	UserSalt     string `json:"user_salt,omitempty"     dc:"加密盐 生成密码用"`
	SecretAnswer string `json:"secret_answer" dc:"密保问题的答案"`
}
type UpdatePasswordRes struct {
	Id uint `json:"id"`
}
