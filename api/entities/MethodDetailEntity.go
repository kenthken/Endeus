package entities

var CreateMethodDetailTable = "method_detail"

type MethodDetail struct {
	MethodDetailId int    `gorm:"column:method_detail_id;size:30;not null;primaryKey;autoIncrement" json:"method_detail_id"`
	MethodId       int    `gorm:"column:method_id;size:30;not null" json:"method_id"`
	Detail         string `gorm:"column:detail;size:50" json:"detail"`
}

func (*MethodDetail) TableName() string {
	return CreateMethodDetailTable
}
