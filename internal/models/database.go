package models

import (
	g "gopress/internal/global"
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
}
