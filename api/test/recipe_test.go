package test

import (
	"endeus/api/config"
	"endeus/api/entities"
	"fmt"
	"testing"
)

func TestCreateRecipe(t *testing.T) {

	config.InitEnvConfigs(true)
	db := config.InitDB()

	ingredientD := []entities.IngredientDetail{
		{Description: "ayam 1 ekor"},
		{Description: "garam"},
	}

	ingredient := []entities.Ingredient{
		{
			Title:            "Bahan",
			Portion:          4,
			IngredientDetail: ingredientD,
		},
	}

	methodD := []entities.MethodDetail{
		{Detail: "masak ayammya 50 menit"},
		{Detail: "taburin garam"},
	}

	method := entities.Method{
		CookDuration: 4,
		Tips:         "cook it well",
		MethodDetail: methodD,
	}

	entities := entities.Recipe{

		RecipeId: 0,

		Ingredient:  ingredient,
		Method:      &method,
		Title:       "bebek goreng",
		Description: "bebek goreng enak",
		Photo:       []byte{0x12, 0x34, 0x56, 0x78},
	}

	if err := db.Save(&entities).Error; err != nil {
		fmt.Print("error create recipe ", err)

	}

}
