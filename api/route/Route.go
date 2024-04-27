package route

import (
	"endeus/api/config"
	"endeus/api/controller"
	repositoriesimpl "endeus/api/repositories/repositoriesImpl"
	servicesimpl "endeus/api/services/servicesImpl"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func StartRouting(db *gorm.DB) {

	recipeRepository := repositoriesimpl.StartRecipeRepositoryImpl()
	recipeService := servicesimpl.StartRecipeService(recipeRepository, db)
	recipeController := controller.NewRecipeController(recipeService)

	recipeRouter := RecipeRouter(recipeController)

	r := chi.NewRouter()

	r.Mount("/recipe", recipeRouter)

	server := http.Server{
		Addr:    config.EnvConfigs.ClientOrigin,
		Handler: r,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
