package entities

var CreateIngredientTable = "ingredient"

type Ingredient struct {
	IngredientId     int                `gorm:"column:ingredient_id;size:30;not null;primaryKey" json:"ingredient_id"`
	RecipeId         int                `gorm:"column:recipe_id;size:30;not null;index" json:"recipe_id"`
	Title            string             `gorm:"column:title;size:30;not null" json:"title"`
	Portion          int                `gorm:"column:portion;size:30;not null" json:"portion"`
	IngredientDetail []IngredientDetail `gorm:"foreignKey:ingredient_id; references:ingredient_id"`
}

func (*Ingredient) TableName() string {
	return CreateIngredientTable
}
