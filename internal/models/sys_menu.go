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
	Status           string `json:"status" gorm:"size:4;comment:状态"`
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
	s := SysMenu{}
	st := reflect.TypeOf(s)

	// st.Len()

	// field1 := st.Field(0)
	// fmt.Printf("json:%v\n", field1.Name)
	// fmt.Printf("type:%v\n", field1.Type)
	// fmt.Printf("pkgpath:%v\n", field1.PkgPath)

	// if field1.Anonymous {
	// 	fmt.Println("field1.Type 是一个嵌入类型")
	// } else {
	// 	fmt.Println("field1.Type 不是一个嵌入类型")
	// }

	len := st.Len()
	for i := 0; i < len; i++ {
		field := st.Field(i)
		fmt.Printf("field:%v\n", field.Name)
		if field.Anonymous {
			fmt.Println("field1.Type 是一个嵌入类型")
		} else {
			fmt.Println("field1.Type 不是一个嵌入类型")
		}
	}

	// fmt.Printf("json:%v\n", field1.Tag.Get("json"))
	// fmt.Printf("xml:%v\n", field1.Tag.Get("xml"))

	// filed2 := st.Field(1)
	// fmt.Printf("json:%v\n", filed2.Tag.Get("json"))
}
