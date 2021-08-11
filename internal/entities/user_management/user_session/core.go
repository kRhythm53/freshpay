package user_session

import (
	"github.com/freshpay/internal/config"
	utilities2 "github.com/freshpay/utilities"
	"time"
)


//CreateSession will create a new user_session
func CreateSession(session *Detail)(err error){
	session.ID= utilities2.CreateID(Prefix, IDLengthExcludingPrefix)
	session.ExpireTime=uint64(time.Now().Unix()+ExpireTime)

	if err=config.DB.Create(session).Error; err!=nil{
		return err
	}
	return nil
}

//GetSessionById will return the user_session by using the id
func GetSessionById(session *Detail, id string) (err error) {
	if err = config.DB.Where("ID = ?", id).First(session).Error; err != nil {
		return err
	}
	return nil
}

//GetActiveSessionByUserId will return active session for that user
func GetActiveSessionByUserId(session *Detail,user_id string)(err error){
	err=config.DB.Where("user_id = ? AND expire_time >= ?", user_id, time.Now().Unix()).First(&session).Error
	return err
}
