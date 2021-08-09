package user

import (
	"errors"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/user_management/session"
	"github.com/freshpay/internal/entities/user_management/utilities"
	"github.com/freshpay/internal/entities/user_management/wallet"

	//"github.com/freshpay/internal/entities/user_management/wallet"
)



func VerifyPhoneNumber(phoneNumber string) bool{
	if phoneNumber=="1"{
		return false
	}
	return true
}
//SignUp will be used to create a user on signup
func SignUp(user *Detail) (err error){
	phoneNumber :=user.PhoneNumber

	/*
	    Make sure PhoneNumber doesn't exist
	 */
	var userTemp Detail
	err=GetUserByPhoneNumber(&userTemp,phoneNumber)
	if err==nil{
		err=errors.New("Phone Number is already registered")
		return err
	}
	if !VerifyPhoneNumber(phoneNumber){
		err=errors.New("OTP entered is wrong, Try again")
		return err
	}
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

func UpdateTransactionCount(userID string)(err error){
	var User Detail
	err = GetUserById(&User, userID)
	if err != nil {
		return
	}
	User.NumberOfTransactions= User.NumberOfTransactions+1
	config.DB.Table("user").Save(&User)
	return nil
}