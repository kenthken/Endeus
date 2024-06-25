package payloads

type CreateRecipeRequest struct {
	RecipeId     int           ` json:"recipe_id"`
	CategoryId   int           ` json:"category_id"`
	IngredientId int           ` json:"ingredient_id"`
	MethodId     int           ` json:"method_id"`
	Title        string        ` json:"title"`
	Ingredient   []*Ingredient `json:"ingredient"`
	Method       *Method
	Rating       *RatingData
	Discussion   []*DiscussionData
	Description  string ` json:"description"`
	Photo        string ` json:"photo"`
}

type GetRecipeList struct {
	RecipeId     int           ` json:"recipe_id"`
	CategoryId   int           ` json:"category_id"`
	CategoryName string        ` json:"title_category" gorm:"column:title_category"`
	Ingredient   []*Ingredient `gorm:"embedded"`
	// MethodData     *MethodData
	// Rating         *RatingData
	// Discussion     *DiscussionData
	Rate        float64 ` json:"rate"`
	Title       string  ` json:"title"`
	Description string  ` json:"description"`
	Photo       []byte  ` json:"photo"`
}

type Category struct {
	CategoryId    int
	TitleCategory string
}

type DiscussionData struct {
	DiscussionId    int ` json:"discussion_id"`
	User            User
	Description     string ` json:"description"`
	Photo           string ` json:"photo"`
	DiscussionReply []DiscussionReply
}

type User struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Photo  string `json:"photo"`
}

type DiscussionReply struct {
	DiscussionReplyId int    ` json:"discussion_reply_id"`
	DiscussionId      int    ` json:"discussion_id"`
	Description       string ` json:"description"`
}

type Ingredient struct {
	IngredientId     int    ` json:"ingredient_id"`
	Title            string ` json:"title"`
	RecipeId         int    `json:"recipe_id"`
	Portion          int    `json:"portion"`
	IngredientDetail string `json:"ingredient_detail"`
}

type Method struct {
	MethodId      int    `json:"method_id"`
	CookDuration  int    `json:"cook_duration"`
	Tips          string `json:"tips"`
	MethodDetails string `json:"method_details"`
}

type RatingData struct {
	RatingId int     `json:"rating_id"`
	Rate     float64 `json:"rate"`
}
