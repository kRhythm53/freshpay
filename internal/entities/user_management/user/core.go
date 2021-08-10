package user

import (
	"errors"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/user_management/session"
	"github.com/freshpay/internal/entities/user_management/utilities"
	"github.com/freshpay/internal/entities/user_management/wallet"

	//"github.com/freshpay/internal/entities/user_management/wallet"
)


//SignUp will be used to create a user on signup
func SignUp(user *Detail) (err error){

	user.ID=utilities.CreateID(Prefix,14)

	if err=config.DB.Create(user).Error; err!=nil{
		return err
	}

	err=wallet.CreateWallet(user.ID)
	if err!=nil{
		config.DB.Unscoped().Delete(&user)
		return err
	}
	return nil
}

//GetUserById will get the user infromation by using Id
func GetUserById(user *Detail, id string)(err error){
	if err = config.DB.Where("ID = ?",id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByPhoneNumber will get the user details by phone number
func GetUserByPhoneNumber(user *Detail, phoneNumber string)(err error){
	print("ad")
	if err = config.DB.Where("phone_number = ?", phoneNumber).First(user).Error; err != nil {
		return err
	}
	return nil
}

//Login will login the user and will create a session
func LoginByPassword(phoneNumber string, password string, Session *session.Detail)(err error) {
	var user Detail
	err = GetUserByPhoneNumber(&user, phoneNumber)
	if err == nil {
		if user.Password != password {
			err = errors.New("Password is Wrong")
		} else {
			Session.UserId=user.ID
			err = session.CreateSession(Session)
		}
	}
	return err
}