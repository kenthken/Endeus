package entities

var CreateDiscussionTable = "discussion"

type Discussion struct {
	DiscussionId    int               `gorm:"column:discussion_id;size:30;not null;primaryKey" json:"discussion_id"`
	RecipeId        int               `gorm:"column:recipe_id;size:30;not null" json:"recipe_id"`
	UserId          int               `gorm:"column:user_id;size:30;not null" json:"user_id"`
	Description     string            `gorm:"column:description;not null" json:"description"`
	Photo           []byte            `gorm:"column:photo;" json:"photo"`
	DiscussionReply []DiscussionReply `gorm:"foreignKey:discussion_id; references:discussion_id"`
}

func (*Discussion) TableName() string {
	return CreateDiscussionTable
}
