package models

type SysRole struct {
	Model
	Alias string `json:"alias" gorm:"size:20;comment:别名"`
	Label string `json:"label" gorm:"size:64;comment:名称标签"`
	Sort  int    `json:"sort" gorm:"size:11;comment:排序"`
	ModelTime
	ControlBy // 公共模字段
}
