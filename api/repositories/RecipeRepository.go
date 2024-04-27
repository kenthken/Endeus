package repositories

import (
	"endeus/api/entities"
	exceptions "endeus/api/expections"
	"endeus/api/pagination"
	"endeus/api/payloads"

	"gorm.io/gorm"
)

type RecipeRepository interface {
	CreateRecipe(db *gorm.DB, req payloads.CreateRecipeRequest) (bool, *exceptions.BaseErrorResponse)
	GetListRecipe(db *gorm.DB, pages pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse)
	GetRecipeDetail(db *gorm.DB, ID int) (entities.Recipe, *exceptions.BaseErrorResponse)
	DeleteRecipe(db *gorm.DB, ID int) (bool, *exceptions.BaseErrorResponse)
	UpdateRecipe(db *gorm.DB, req payloads.CreateRecipeRequest) (bool, *exceptions.BaseErrorResponse)
}
