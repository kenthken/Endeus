package entities

var CreateUserTable = "user"

type User struct {
	UserId     int        `gorm:"column:user_id;size:30;not null;primaryKey" json:"user_id"`
	Name       string     `gorm:"column:name;size:30;not null" json:"name"`
	Photo      []byte     `gorm:"column:photo;not null" json:"photo"`
	Discussion Discussion `gorm:"foreignKey:user_id; references:user_id"`
}

func (*User) TableName() string {
	return CreateUserTable
}
