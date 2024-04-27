package entities

var CreateRecipeTable = "recipe"

type Recipe struct {
	RecipeId    int          `gorm:"column:recipe_id;size:30;not null;primaryKey" json:"recipe_id"`
	Ingredient  []Ingredient `gorm:"constraint:OnDelete:CASCADE;foreignKey:recipe_id" references:"recipe_id"`
	Method      *Method      `gorm:"constraint:OnDelete:CASCADE;foreignKey:recipe_id" references:"recipe_id"`
	Title       string       `gorm:"column:title;size:50;not null" json:"title"`
	Description string       `gorm:"column:description;size:50;not null" json:"description"`
	Photo       []byte       `gorm:"column:photo;not null" json:"photo"`
	Discussion  []Discussion `gorm:"constraint:OnDelete:CASCADE;foreignKey:recipe_id" references:"recipe_id"`
	Rating      []Rating     `gorm:"constraint:OnDelete:CASCADE;foreignKey:recipe_id" references:"recipe_id"`
}

func (*Recipe) TableName() string {
	return CreateRecipeTable
}
