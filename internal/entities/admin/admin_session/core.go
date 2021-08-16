package admin_session

import (
	"github.com/freshpay/internal/config"
	"github.com/freshpay/utilities"
	"time"
)


//CreateSession will create a new user_session
func CreateSession(session *Detail)(err error){
	session.ID= utilities.CreateID(Prefix, IDLengthExcludingPrefix)
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
//GetActiveSessionByAdminId will return active session for that user
func GetActiveSessionByAdminId(session *Detail, adminId string)(err error){
	err=config.DB.Where("admin_id = ? AND expire_time >= ?", adminId, time.Now().Unix()-ExpireTime).Last(&session).Error
	return err
}
