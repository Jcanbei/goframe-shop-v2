package backend

import "github.com/gogf/gf/v2/frame/g"

type ArticleCommonAddUpdate struct {
	Title   string `json:"title" dc:"文章标题" v:"required#文章标题必传"`
	Desc    string `json:"desc" dc:"文章概要"`
	PicUrl  string `json:"pic_url" dc:"图片"`
	IsAdmin uint   `d:"1" dc:"1后台管理员发布 2前台管理员发布"`
	Detail  string `json:"detail" dc:"文章详情" v:"required#文章详情必填"`
	Praise  int    `json:"praise" dc:"点赞数量"`
}

type ArticleReq struct {
	g.Meta `path:"/article/add" tags:"文章" method:"post" summary:"添加文章接口"`
	ArticleCommonAddUpdate
}
type ArticleRes struct {
	Id uint `json:"id"`
}

type ArticleDeleteReq struct {
	g.Meta `path:"/article/delete" method:"delete" tags:"文章" summary:"删除文章接口"`
	Id     uint `v:"min:1#请选择需要删除的文章id" dc:"文章id"`
}
type ArticleDeleteRes struct{}

type ArticleUpdateReq struct {
	g.Meta `path:"/article/update" method:"post" tags:"文章" summary:"修改文章接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的文章id" dc:"文章Id"`
	ArticleCommonAddUpdate
}
type ArticleUpdateRes struct {
	Id uint `json:"id"`
}

type ArticleGetListCommonReq struct {
	g.Meta `path:"/article/list" method:"get" tags:"文章" summary:"文章列表接口"`
	CommonPaginationReq
}
type ArticleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
