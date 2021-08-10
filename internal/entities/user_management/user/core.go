package user

import (
	"errors"
	"fmt"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/OTP"
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
	if err==nil && userTemp.IsVerified{
		err=errors.New("Phone Number is already registered")
		return err
	} else if err==nil{
		err=DeleteUser(&userTemp)
		if err!=nil{
			return err
		}
	}
	err= OTP.SendOTP(phoneNumber)
	if err!=nil{
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


//update the user
func UpdateUser(user *Detail)(err error){
	err=config.DB.Save(user).Error
	if err!=nil{
		return err
	}
	return nil
}


//Delete a user
func DeleteUser(user *Detail)(err error){
	err=config.DB.Where("id = ?",user.ID).Delete(user).Error
	return err
}
//Login will login the user and will create a session
func LoginByPassword(phoneNumber string, password string, Session *session.Detail)(err error) {
	var user Detail
	err = GetUserByPhoneNumber(&user, phoneNumber)
	if err == nil {
		if !user.IsVerified{
			err=errors.New("Phone Number is not verified, please signup again")
			/*
			   need to remove this line
			*/
			//return err
		}
		if user.Password != password {
			err = errors.New("Password is Wrong")
		} else {
			Session.UserId=user.ID
			err = session.CreateSession(Session)
		}
	}
	return err
}

//set verified user by phone number
func SetVerifiedUserByPhoneNumber(phoneNumber string)(err error){
	var user Detail
	err = GetUserByPhoneNumber(&user, phoneNumber)
	if err==nil{
		user.IsVerified=true
		err=UpdateUser(&user)
	}
	fmt.Println(user)
	return err
}