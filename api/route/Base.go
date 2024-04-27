package route

import (
	"endeus/api/controller"

	"github.com/go-chi/chi/v5"
)

func RecipeRouter(
	RecipeController controller.RecipeController,
) chi.Router {
	router := chi.NewRouter()

	router.Get("/", RecipeController.GetListRecipe)
	router.Get("/{recipe_id}", RecipeController.GetRecipeDetail)
	router.Post("/", RecipeController.CreateRecipe)
	router.Patch("/", RecipeController.UpdateRecipe)
	router.Delete("/{recipe_id}", RecipeController.DeleteRecipe)

	return router
}
