package models

import (
	g "gopress/internal/global"
	"gopress/pkg"
	"time"

	"gorm.io/gorm"
)

type Model struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type ControlBy struct {
	CreatedBy int `json:"createBy" gorm:"index;comment:创建者"`
	UpdatedBy int `json:"updateBy" gorm:"index;comment:更新者"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}

// 初始化数据库
func InitModel() {
	// 初始化数据表
	{
		// 配置
		g.DB.AutoMigrate(&SysConfig{})
		// 用户角色
		g.DB.AutoMigrate(&SysUser{})
		g.DB.AutoMigrate(&SysRole{})
		g.DB.AutoMigrate(&SysUserRole{})
		// 菜单
		g.DB.AutoMigrate(&SysMenu{})
		g.DB.AutoMigrate(&SysRoleMenu{})
	}

	// 创建初始数据
	// 查找SysRole 是否存在 Alias = superadmin的角色
	var role SysRole
	g.DB.Where("alias = ?", "superadmin").First(&role)
	if role.Id == 0 {
		role = SysRole{Alias: "superadmin", Label: "超级管理员", Sort: 0, Status: 1}
		g.DB.Create(&role)
	}

	// 查找SysUser 是否存在 Username = superadmin 的用户
	var user SysUser
	g.DB.Where("username = ?", "superadmin").First(&user)
	if user.Id == 0 {
		user = SysUser{
			Username: "superadmin",
			Nickname: "超级管理员",
			Password: pkg.PwdEncode("123456"),
			Email:    "superadmin@example.com",
			Phone:    "1234567890",
			Status:   1,
		}
		g.DB.Create(&user)
	}

	// 查找sys_user_role 是否存在 UserID = user.Id RoleID = role.Id
	var userRole SysUserRole
	g.DB.Where("user_id = ? AND role_id = ?", user.Id, role.Id).First(&userRole)
	if userRole.Id == 0 {
		userRole = SysUserRole{UserID: user.Id, RoleID: role.Id}
		g.DB.Create(&userRole)
	}
}
