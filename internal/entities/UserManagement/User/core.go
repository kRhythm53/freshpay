package User

import "github.com/kshitij-nawandar9/freshpay/internal/Config"

//CreateUser will be used to create a user on signup
func CreateUser(user *Detail) (err error){
	if err=Config.DB.Create(user).Error; err!=nil{
		return err
	}
	return nil
}

//GetUserById will get the user infromation by using Id
func GetUserByID(user *Detail, entity_id string)(err error){
	if err = Config.DB.Where("entity_id = ?", entity_id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByPhoneNumber will get the user details by phone number
func GetUserByPhoneNumber(user *Detail, phone_number string)(err error){
	if err = Config.DB.Where("phone_number = ?", phone_number).First(user).Error; err != nil {
		return err
	}
	return nil
}