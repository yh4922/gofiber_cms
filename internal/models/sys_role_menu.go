package models

type SysRoleMenu struct {
	Model
	RoleID int `json:"roleId" gorm:"size:11;comment:角色ID"`
	MenuID int `json:"menuId" gorm:"size:11;comment:菜单ID"`
}
