package entities

var CreateIngredientDetailTable = "ingredient_detail"

type IngredientDetail struct {
	IngredientDetailId int    `gorm:"column:ingredient_detail_id;size:30;not null;primaryKey;autoIncrement" json:"ingredient_detail_id"`
	IngredientId       int    `gorm:"column:ingredient_id;size:30;not null" json:"ingredient_id"`
	Description        string `gorm:"column:description;not null" json:"description"`
}

func (*IngredientDetail) TableName() string {
	return CreateIngredientDetailTable
}
