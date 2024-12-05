package models

import (
	"fmt"
	"reflect"
)

type SysMenu struct {
	Model
	ParentId         int    `json:"parentId" gorm:"size:11;comment:上级菜单ID"`
	Title            string `json:"title" gorm:"size:200;comment:标题"`
	Type             string `json:"type" gorm:"size:10;comment:类型"`
	Name             string `json:"name" gorm:"size:100;comment:名称"`
	Icon             string `json:"icon" gorm:"size:100;comment:菜单图标"`
	ActiveIcon       string `json:"activeIcon" gorm:"size:100;comment:菜单高亮图标"`
	Path             string `json:"path" gorm:"size:100;comment:路由路径"`
	Status           int    `json:"status" gorm:"size:1;comment:状态"` // ENABLE 1   DISABLE 0
	Redirect         string `json:"redirect" gorm:"size:200;comment:重定向"`
	Active           string `json:"active" gorm:"size:200;comment:菜单高亮"`
	Layout           string `json:"layout" gorm:"size:100;comment:布局组件（前端不一定支持预留）"`
	Component        string `json:"component" gorm:"size:100;comment:视图组件"`
	Color            string `json:"color" gorm:"size:10;comment:颜色"`
	Hidden           bool   `json:"hidden" gorm:";comment:隐藏菜单"`
	HiddenBreadcrumb bool   `json:"hiddenBreadcrumb" gorm:";comment:隐藏面包屑"`
	DefaultOpened    bool   `json:"defaultOpened" gorm:";comment:默认展开"`
	Affix            bool   `json:"affix" gorm:"comment:是否常驻标签页"`
	Sort             int    `json:"sort" gorm:"size:11;comment:排序"`
	ModelTime
	ControlBy
}

func init() {
	// getModelInfo(SysMenu{})
}

// / 获取模型信息
func GetModelInfo(s interface{}) {
	st := reflect.TypeOf(s)
	len := st.NumField()

	for i := 0; i < len; i++ {
		field := st.Field(i)

		if field.Anonymous {
			/// 嵌入类型
			subLen := field.Type.NumField()
			for j := 0; j < subLen; j++ {
				subField := field.Type.Field(j)
				fmt.Printf("%v %v\n", subField.Name, subField.Type)
			}
		} else {
			/// 原始类型
			fmt.Printf("%v %v\n", field.Name, field.Tag.Get("json"))
			// field.Tag.Get("json")

		}
	}
}
