package models

type SysUserRole struct {
	Model
	UserID int `json:"userId" gorm:"size:11;comment:用户ID"`
	RoleID int `json:"roleId" gorm:"size:11;comment:角色ID"`
}
