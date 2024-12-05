package models

type SysUser struct {
	Model
	Username string `json:"username" gorm:"size:20;comment:用户名"`
	Nickname string `json:"nickname" gorm:"size:64;comment:昵称"`
	Password string `json:"password" gorm:"size:255;comment:密码"`
	Email    string `json:"email" gorm:"size:64;comment:邮箱"`
	Phone    string `json:"phone" gorm:"size:20;comment:手机号"`
	Status   int    `json:"status" gorm:"size:1;comment:状态"` // ENABLE 1   DISABLE 0
	ModelTime
	ControlBy
}
