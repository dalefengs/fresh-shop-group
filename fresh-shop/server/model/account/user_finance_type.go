package account

import (
	"fresh-shop/server/global"
)

// UserFinanceType 结构体
type UserFinanceType struct {
	global.DbModel
	ParentId *int              `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父级ID;size:20;"`
	Name     string            `json:"name" form:"name" gorm:"column:name;comment:流水类型名称;size:20;"`
	Remarks  string            `json:"remarks" form:"remarks" gorm:"column:remarks;comment:备注;size:255;"`
	Children []UserFinanceType `json:"children" gorm:"foreignKey:parent_id"`
}

// TableName UserFinanceType 表名
func (UserFinanceType) TableName() string {
	return "user_finance_type"
}
