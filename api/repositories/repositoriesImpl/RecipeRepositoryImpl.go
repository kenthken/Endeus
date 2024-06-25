package repositoriesimpl

import (
	"endeus/api/entities"
	exceptions "endeus/api/expections"
	"endeus/api/pagination"
	"endeus/api/payloads"
	"endeus/api/repositories"
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type RecipeRepositoryImpl struct {
}

// UpdateRecipe implements repositories.RecipeRepository.
func (r *RecipeRepositoryImpl) UpdateRecipe(db *gorm.DB, req payloads.CreateRecipeRequest) (bool, *exceptions.BaseErrorResponse) {

	for _, value := range req.Ingredient {
		ingredient := entities.Ingredient{Title: value.Title, Portion: value.Portion, IngredientId: value.IngredientId, RecipeId: req.RecipeId, IngredientDetail: value.IngredientDetail}
		fmt.Print("sini?", ingredient)
		if ingredient.IngredientId != 0 {
			if err := db.Where(entities.Ingredient{IngredientId: value.IngredientId}).Updates(&ingredient).Error; err != nil {
				return false, &exceptions.BaseErrorResponse{
					StatusCode: http.StatusBadRequest,
					Err:        err,
				}
			}
		} else {
			if err := db.Save(&ingredient).Error; err != nil {
				return false, &exceptions.BaseErrorResponse{
					StatusCode: http.StatusBadRequest,
					Err:        err,
				}
			}
		}
	}

	method := entities.Method{
		MethodId:      req.Method.MethodId,
		RecipeId:      req.RecipeId,
		CookDuration:  req.Method.CookDuration,
		MethodDetails: req.Method.MethodDetails,
		Tips:          req.Method.Tips,
	}

	if req.Method.MethodId != 0 {
		if err := db.Where(entities.Method{MethodId: req.Method.MethodId}).Updates(&method).Error; err != nil {
			return false, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusBadRequest,
				Err:        err,
			}
		}
	}

	recipeEntities := entities.Recipe{
		RecipeId:    req.RecipeId,
		Title:       req.Title,
		Description: req.Description,
		Photo:       []byte(req.Photo),
	}

	if err := db.Where("recipe_id = ?", req.RecipeId).Updates(&recipeEntities).Error; err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}

	}

	return true, nil
}

// CreateRecipe implements repositories.RecipeRepository.
func (r *RecipeRepositoryImpl) CreateRecipe(db *gorm.DB, req payloads.CreateRecipeRequest) (bool, *exceptions.BaseErrorResponse) {

	var ingredient []entities.Ingredient

	for _, value := range req.Ingredient {

		ingredient = append(ingredient, entities.Ingredient{Title: value.Title, Portion: value.Portion, IngredientDetail: value.IngredientDetail})
	}

	method := entities.Method{
		CookDuration:  req.Method.CookDuration,
		Tips:          req.Method.Tips,
		MethodDetails: req.Method.MethodDetails,
	}

	recipeEntities := entities.Recipe{
		RecipeId:    0,
		Ingredient:  ingredient,
		Method:      &method,
		Title:       req.Title,
		Description: req.Description,
		Photo:       []byte(req.Photo),
	}

	if err := db.Save(&recipeEntities).Error; err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}

	}

	return true, nil
}

// DeleteRecipe implements repositories.RecipeRepository.
func (r *RecipeRepositoryImpl) DeleteRecipe(db *gorm.DB, ID int) (bool, *exceptions.BaseErrorResponse) {

	result, err := r.GetRecipeDetail(db, ID)

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusConflict,
			Err:        err.Err,
		}
	}

	//delete ingredient

	ingredient := entities.Ingredient{}

	if err := db.Where("recipe_id = ?", ID).Delete(ingredient).Error; err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusConflict,
			Err:        err,
		}
	}

	//delete method

	method := entities.Method{}

	if err := db.Where("recipe_id = ?", ID).Delete(method).Error; err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusConflict,
			Err:        err,
		}
	}

	//delete discussiond

	discussionD := entities.DiscussionReply{}

	for _, value := range result.Discussion {
		if err := db.Where("discussion_id = ?", value.DiscussionId).Delete(discussionD).Error; err != nil {
			return false, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusConflict,
				Err:        err,
			}
		}
	}

	//delete discussion

	discussion := entities.Discussion{}

	if err := db.Where("recipe_id = ?", ID).Delete(discussion).Error; err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusConflict,
			Err:        err,
		}
	}

	//delete rating

	rating := entities.Rating{}

	if err := db.Where("recipe_id = ?", ID).Delete(rating).Error; err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusConflict,
			Err:        err,
		}
	}

	//delete recipe

	model := &entities.Recipe{}

	if err := db.Where("recipe_id = ?", ID).Delete(model).Error; err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusConflict,
			Err:        err,
		}
	}

	return true, nil
}

// GetListRecipe implements repositories.RecipeRepository.
func (r *RecipeRepositoryImpl) GetListRecipe(db *gorm.DB, pages pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse) {
	recipe := []entities.Recipe{}

	// payloads := []payloads.GetRecipeList{}

	// query := db.Model(&recipe).
	// 	Joins("left join method on recipe.recipe_id = method.recipe_id").
	// 	Joins("left join ingredient on recipe.recipe_id = ingredient.recipe_id").
	// 	Joins("left join discussion on recipe.recipe_id = discussion.recipe_id").
	// 	Joins("left join discussion_reply on discussion.discussion_id = discussion_reply.discussion_id").
	// 	Joins("left join rating on recipe.recipe_id = rating.recipe_id")

	//GET RECIPE
	query := db.Model(recipe)

	err := query.Scopes(pagination.Paginate(&recipe, &pages, query)).Scan(&recipe).Error

	//GET INGREDIENT && Method
	ingredientEntity := entities.Ingredient{}
	methodEntity := entities.Method{}

	for key, value := range recipe {
		ingredientResponses := []entities.Ingredient{}
		methodResponses := entities.Method{}
		if err := db.Model(ingredientEntity).Where(entities.Ingredient{RecipeId: value.RecipeId}).Scan(&ingredientResponses).Error; err != nil {
			return pages, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}

		if err := db.Model(methodEntity).Where(entities.Method{RecipeId: value.RecipeId}).First(&methodResponses).Error; err != nil {
			return pages, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}

		recipe[key].Ingredient = ingredientResponses
		recipe[key].Method = &methodResponses

	}

	fmt.Print(recipe)

	if err != nil {
		return pages, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if len(recipe) == 0 {
		return pages, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        errors.New(""),
		}
	}

	pages.Rows = recipe

	return pages, nil

}

// GetRecipeDetail implements repositories.RecipeRepository.
func (r *RecipeRepositoryImpl) GetRecipeDetail(db *gorm.DB, ID int) (entities.Recipe, *exceptions.BaseErrorResponse) {
	recipe := entities.Recipe{}

	if err := db.Preload("Method").Preload("Ingredient").Preload("Discussion").Preload("Rating").Where("recipe_id = ?", ID).First(&recipe).Error; err != nil {
		return recipe, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        errors.New(""),
		}
	}

	return recipe, nil

}

func StartRecipeRepositoryImpl() repositories.RecipeRepository {
	return &RecipeRepositoryImpl{}
}
