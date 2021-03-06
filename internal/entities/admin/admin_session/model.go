package admin_session

import "gorm.io/gorm"

type Detail struct {
	gorm.Model
	ID         string `gorm:"type:varchar(20)"`
	AdminId     string
	ExpireTime uint64
}

const (
	TableName               ="admin_session"
	EntityName              ="admin_session"
	ExpireTime              =300000000  //time in seconds
	IDLengthExcludingPrefix =14
	Prefix                  ="sAdm"
)

func(sd *Detail) TableName() string{
	return TableName
}