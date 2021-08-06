package session

import (
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/user_management/utilities"
	"time"
)


//CreateSession will create a new session
func CreateSession(session *Detail)(err error){
	session.ID=utilities.CreateID(14)
	session.ExpireTime=uint64(time.Now().Unix()+300)

	if err=config.DB.Create(session).Error; err!=nil{
		return err
	}
	return nil
}

//GetSessionById will return the session by using the id
func GetSessionById(session *Detail, id string) (err error) {
	if err = config.DB.Where("ID = ?", id).First(session).Error; err != nil {
		return err
	}
	return nil
}
