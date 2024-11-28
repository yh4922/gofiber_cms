package models

import (
	"gorm.io/gorm"
)

type AdminUser struct {
	ID        string         `gorm:"primarykey;size:16"`
	Name      string         `gorm:"size:24"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
