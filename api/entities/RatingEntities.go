package entities

var CreateRatingTable = "rating"

type Rating struct {
	RatingId int     `gorm:"column:rating_id;size:30;not null;primaryKey" json:"rating_id"`
	RecipeId int     `gorm:"column:recipe_id;size:30;not null" json:"recipe_id"`
	Rate     float64 `gorm:"column:rate;not null" json:"rate"`
}

func (*Rating) TableName() string {
	return CreateRatingTable
}
