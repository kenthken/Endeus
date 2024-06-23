package entities

var CreateMethodTable = "method"

type Method struct {
	MethodId      int    `gorm:"column:method_id;size:30;not null;primaryKey;autoIncrement" json:"method_id"`
	RecipeId      int    `gorm:"column:recipe_id;size:30;not null" json:"recipe_id"`
	CookDuration  int    `gorm:"column:cook_duration;size:30" json:"cook_duration"`
	Tips          string `gorm:"column:tips;size:50" json:"tips"`
	MethodDetails string `gorm:"column:method_details;size:50" json:"method_details"`
}

func (*Method) TableName() string {
	return CreateMethodTable
}
