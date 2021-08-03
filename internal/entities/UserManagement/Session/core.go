package Session

import "github.com/kshitij-nawandar9/freshpay/internal/Config"

//CreateSession will create a new session
func CreateSession(session *Detail)(err error){
	if err=Config.DB.Create(session).Error; err!=nil{
		return err
	}
	return nil
}

//GetSessionById will return the session by using the id
func GetSessionByID(session *Detail, entity_id string)(err error){
	if err = Config.DB.Where("entity_id = ?", entity_id).First(session).Error; err != nil {
		return err
	}
	return nil
}
