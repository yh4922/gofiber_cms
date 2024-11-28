package models

import (
	g "gopress/internal/global"
)

// 初始化数据库
func InitModel() {
	g.DB.AutoMigrate(&AdminUser{})
}
