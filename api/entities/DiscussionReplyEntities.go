package entities

var CreateDiscussionReplyTable = "discussion_reply"

type DiscussionReply struct {
	DiscussionReplyId int    `gorm:"column:discussion_reply_id;size:30;not null;primaryKey" json:"discussion_reply_id"`
	DiscussionId      int    `gorm:"column:discussion_id;size:30;not null;" json:"discussion_id"`
	Description       string `gorm:"column:description;not null" json:"description"`
}

func (*DiscussionReply) TableName() string {
	return CreateDiscussionReplyTable
}
