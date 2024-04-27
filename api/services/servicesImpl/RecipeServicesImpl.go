package servicesimpl

import (
	"endeus/api/entities"
	exceptions "endeus/api/expections"
	helper_test "endeus/api/helper"
	"endeus/api/pagination"
	"endeus/api/payloads"
	"endeus/api/repositories"
	"endeus/api/services"

	"gorm.io/gorm"
)

type RecipeServiceImpl struct {
	RecipeRepo repositories.RecipeRepository
	DB         *gorm.DB
}

// UpdateRecipe implements services.RecipeService.
func (r *RecipeServiceImpl) UpdateRecipe(req payloads.CreateRecipeRequest) (bool, *exceptions.BaseErrorResponse) {
	tx := r.DB.Begin()
	defer helper_test.CommitOrRollback(tx)
	results, err := r.RecipeRepo.UpdateRecipe(tx, req)
	if err != nil {
		return results, err
	}
	return results, nil
}

// CreateRecipe implements services.RecipeService.
func (r *RecipeServiceImpl) CreateRecipe(req payloads.CreateRecipeRequest) (bool, *exceptions.BaseErrorResponse) {
	tx := r.DB.Begin()
	defer helper_test.CommitOrRollback(tx)
	create, err := r.RecipeRepo.CreateRecipe(tx, req)

	if err != nil {
		return create, err
	}

	return create, nil
}

// DeleteRecipe implements services.RecipeService.
func (r *RecipeServiceImpl) DeleteRecipe(ID int) (bool, *exceptions.BaseErrorResponse) {
	tx := r.DB.Begin()
	defer helper_test.CommitOrRollback(tx)
	delete, err := r.RecipeRepo.DeleteRecipe(tx, ID)

	if err != nil {
		return delete, err
	}

	return delete, nil
}

// GetListRecipe implements services.RecipeService.
func (r *RecipeServiceImpl) GetListRecipe(pages pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse) {
	tx := r.DB.Begin()
	defer helper_test.CommitOrRollback(tx)
	get, err := r.RecipeRepo.GetListRecipe(tx, pages)

	if err != nil {
		return get, err
	}

	return get, nil
}

// GetRecipeDetail implements services.RecipeService.
func (r *RecipeServiceImpl) GetRecipeDetail(ID int) (entities.Recipe, *exceptions.BaseErrorResponse) {
	tx := r.DB.Begin()
	defer helper_test.CommitOrRollback(tx)
	get, err := r.RecipeRepo.GetRecipeDetail(tx, ID)

	if err != nil {
		return get, err
	}

	return get, nil
}

func StartRecipeService(RecipeRepo repositories.RecipeRepository, db *gorm.DB) services.RecipeService {
	return &RecipeServiceImpl{
		RecipeRepo: RecipeRepo,
		DB:         db,
	}
}
