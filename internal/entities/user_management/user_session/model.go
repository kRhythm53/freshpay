package user_session

import "gorm.io/gorm"

type Detail struct {
	gorm.Model
	ID         string `gorm:"type:varchar(20)"`
	UserId     string
	ExpireTime uint64
}

const (
	TableName               ="user_session"
	EntityName              ="user_session"
	ExpireTime              =300000  //time in seconds
	IDLengthExcludingPrefix =14
	Prefix                  ="sUsr"
)

func(sd *Detail) TableName() string{
	return TableName
}