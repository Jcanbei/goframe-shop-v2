package role

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-shop-v2/internal/dao"
	"goframe-shop-v2/internal/model"
	"goframe-shop-v2/internal/model/entity"
	"goframe-shop-v2/internal/service"
)

type sRole struct {
}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

// 添加角色
func (s *sRole) Create(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error) {
	// 获取当前时间并格式化为 MySQL 支持的标准格式
	now := gtime.Now().Format("Y-m-d H:i:s")

	// 插入数据时手动设置 created_at 和 updated_at 字段
	data := map[string]interface{}{
		"name":       in.Name,
		"desc":       in.Desc,
		"created_at": now,
		"updated_at": now,
	}

	//插入数据返回id
	lastInsertID, err := dao.RoleInfo.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RoleCreateOutput{RoleId: uint(lastInsertID)}, err
}

// 角色添加权限
func (s *sRole) AddPermission(ctx context.Context, in model.RoleAddPermissionInput) (out model.RoleAddPermissionOutput, err error) {
	// 获取当前时间并格式化为 MySQL 支持的标准格式
	now := gtime.Now().Format("Y-m-d H:i:s")

	// 插入数据时手动设置 created_at 和 updated_at 字段
	data := map[string]interface{}{
		"role_id":       in.RoleId,
		"permission_id": in.PermissionId,
		"created_at":    now,
		"updated_at":    now,
	}

	//插入数据返回id
	lastInsertId, err := dao.RolePermissionInfo.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RoleAddPermissionOutput{Id: uint(lastInsertId)}, err
}

// Delete 删除
func (s *sRole) Delete(ctx context.Context, id uint) error {
	// 删除内容
	_, err := dao.RoleInfo.Ctx(ctx).Where(g.Map{
		dao.RoleInfo.Columns().Id: id,
	}).Unscoped().Delete()
	return err
}

// 角色删除权限
func (s *sRole) DeletePermission(ctx context.Context, in model.RoleDeletePermissionInput) error {
	_, err := dao.RolePermissionInfo.Ctx(ctx).Where(g.Map{
		dao.RolePermissionInfo.Columns().RoleId:       in.RoleId,
		dao.RolePermissionInfo.Columns().PermissionId: in.PermissionId,
	}).Delete()
	if err != nil {
		return err
	}
	return err
}

// Update 修改
func (s *sRole) Update(ctx context.Context, in model.RoleUpdateInput) error {
	// 获取当前时间并格式化为 MySQL 支持的标准格式
	now := gtime.Now().Format("Y-m-d H:i:s")

	// 插入数据时手动设置updated_at 字段
	data := map[string]interface{}{
		"name":       in.Name,
		"desc":       in.Desc,
		"updated_at": now,
	}

	//更新操作
	_, err := dao.RoleInfo.
		Ctx(ctx).
		Data(data).
		FieldsEx(dao.RoleInfo.Columns().Id).
		Where(dao.RoleInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询内容列表
func (s *sRole) GetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error) {
	var (
		m = dao.RoleInfo.Ctx(ctx)
	)
	out = &model.RoleGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.OrderDesc(dao.RoleInfo.Columns().Id)
	// 执行查询
	var list []*entity.RoleInfo
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
	// 不指定item的键名用：Scan
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
