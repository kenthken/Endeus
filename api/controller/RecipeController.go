package controller

import (
	exceptions "endeus/api/expections"
	helper_test "endeus/api/helper"
	jsonchecker "endeus/api/helper/json/json-checker"
	"endeus/api/pagination"
	"endeus/api/payloads"
	"endeus/api/services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RecipeController interface {
	CreateRecipe(writer http.ResponseWriter, request *http.Request)
	GetListRecipe(writer http.ResponseWriter, request *http.Request)
	GetRecipeDetail(writer http.ResponseWriter, request *http.Request)
	DeleteRecipe(writer http.ResponseWriter, request *http.Request)
	UpdateRecipe(writer http.ResponseWriter, request *http.Request)
}

type RecipeControllerImpl struct {
	recipeservice services.RecipeService
}

// UpdateRecipe implements RecipeController.
func (r *RecipeControllerImpl) UpdateRecipe(writer http.ResponseWriter, request *http.Request) {
	var formRequest payloads.CreateRecipeRequest

	err := jsonchecker.ReadFromRequestBody(request, &formRequest)

	if err != nil {
		exceptions.NewBadRequestException(writer, request, err)
		return
	}

	create, err := r.recipeservice.UpdateRecipe(formRequest)

	if err != nil {
		helper_test.ReturnError(writer, request, err)
		return
	}

	payloads.NewHandleSuccess(writer, create, "Update Data Successfully!", http.StatusOK)
}

// CreateRecipe implements RecipeController.
func (r *RecipeControllerImpl) CreateRecipe(writer http.ResponseWriter, request *http.Request) {
	var formRequest payloads.CreateRecipeRequest

	err := jsonchecker.ReadFromRequestBody(request, &formRequest)

	if err != nil {
		exceptions.NewBadRequestException(writer, request, err)
		return
	}

	create, err := r.recipeservice.CreateRecipe(formRequest)

	if err != nil {
		helper_test.ReturnError(writer, request, err)
		return
	}

	payloads.NewHandleSuccess(writer, create, "Create Data Successfully!", http.StatusOK)
}

// DeleteRecipe implements RecipeController.
func (r *RecipeControllerImpl) DeleteRecipe(writer http.ResponseWriter, request *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(request, "recipe_id"))

	result, err := r.recipeservice.DeleteRecipe(id)

	if err != nil {
		helper_test.ReturnError(writer, request, err)
		return
	}

	payloads.NewHandleSuccess(writer, result, "Delete Data Successfully!", http.StatusOK)

}

// GetListRecipe implements RecipeController.
func (r *RecipeControllerImpl) GetListRecipe(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()

	limit, _ := strconv.Atoi(queryValues.Get("limit"))
	page, _ := strconv.Atoi(queryValues.Get("page"))

	pagination := pagination.Pagination{
		Limit:  limit,
		Page:   page,
		SortOf: queryValues.Get("sort_of"),
		SortBy: queryValues.Get("sort_by"),
	}

	result, err := r.recipeservice.GetListRecipe(pagination)
	if err != nil {
		exceptions.NewNotFoundException(writer, request, err)
		return
	}
	payloads.NewHandleSuccess(writer, result, "Get Data Successfully!", http.StatusOK)
}

// GetRecipeDetail implements RecipeController.
func (r *RecipeControllerImpl) GetRecipeDetail(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(request, "recipe_id"))

	result, err := r.recipeservice.GetRecipeDetail(id)

	if err != nil {
		helper_test.ReturnError(writer, request, err)
		return
	}

	payloads.NewHandleSuccess(writer, result, "Get Data Successfully!", http.StatusOK)

}

func NewRecipeController(recipeService services.RecipeService) RecipeController {
	return &RecipeControllerImpl{
		recipeservice: recipeService,
	}
}
