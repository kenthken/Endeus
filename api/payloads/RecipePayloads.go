package payloads

type CreateRecipeRequest struct {
	RecipeId     int    ` json:"recipe_id"`
	CategoryId   int    ` json:"category_id"`
	IngredientId int    ` json:"ingredient_id"`
	MethodId     int    ` json:"method_id"`
	Title        string ` json:"title"`
	Ingredient   []*Ingredient
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
	IngredientId     int                 ` json:"ingredient_id"`
	IngredientTitle  string              ` json:"ingredient_title"`
	Portion          int                 `json:"portion"`
	IngredientDetail []*IngredientDetail `gorm:"embedded" json:"ingredient_detail"`
}

type IngredientDetail struct {
	IngredientDetailId int    `json:"ingredient_detail_id"`
	IngredientId       int    ` json:"ingredient_id"`
	Description        string `json:"description"`
}

type Method struct {
	MethodId     int             `json:"method_id"`
	CookDuration int             `json:"cook_duration"`
	Tips         string          `json:"tips"`
	MethodDetail []*MethodDetail `gorm:"embedded" json:"method_detail"`
}

type MethodDetail struct {
	MethodDetailId int    `json:"method_detail_id"`
	Detail         string `json:"detail"`
}

type RatingData struct {
	RatingId int     `json:"rating_id"`
	Rate     float64 `json:"rate"`
}
