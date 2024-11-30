package models

type SysConfig struct {
	Model
	Key    string `json:"key" gorm:"size:20;comment:键"`
	Label  string `json:"label" gorm:"size:20;comment:标签"`
	Value  string `json:"value" gorm:"size:20;comment:值"`
	Remark string `json:"remark" gorm:"size:20;comment:备注"`
	Status string `json:"status" gorm:"size:4;comment:状态"`
	ModelTime
	ControlBy // 公共模字段
}
