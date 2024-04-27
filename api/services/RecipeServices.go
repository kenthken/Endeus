package services

import (
	"endeus/api/entities"
	exceptions "endeus/api/expections"
	"endeus/api/pagination"
	"endeus/api/payloads"
)

type RecipeService interface {
	CreateRecipe(req payloads.CreateRecipeRequest) (bool, *exceptions.BaseErrorResponse)
	GetListRecipe(pages pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse)
	GetRecipeDetail(ID int) (entities.Recipe, *exceptions.BaseErrorResponse)
	DeleteRecipe(ID int) (bool, *exceptions.BaseErrorResponse)
	UpdateRecipe(req payloads.CreateRecipeRequest) (bool, *exceptions.BaseErrorResponse)
}
