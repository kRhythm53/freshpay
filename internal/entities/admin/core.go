package admin

import (
	"errors"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/admin/admin_session"
	"github.com/freshpay/internal/entities/user_management/utilities"
)


//SignUp will be used to create a admin on signup
func SignUp(admin *Detail) (err error){

	admin.ID=utilities.CreateID(Prefix,14)

	if err=config.DB.Create(admin).Error; err!=nil{
		return err
	}

	if err!=nil{
		config.DB.Unscoped().Delete(&admin)
		return err
	}
	return nil
}

//GetAdminById will get the admin infromation by using Id
func GetAdminById(admin *Detail, id string)(err error){
	if err = config.DB.Where("ID = ?",id).First(admin).Error; err != nil {
		return err
	}
	return nil
}

//GetAdminByPhoneNumber will get the Admin details by phone number
func GetAdminByPhoneNumber(admin *Detail, phoneNumber string)(err error){
	if err = config.DB.Where("phone_number = ?", phoneNumber).First(admin).Error; err != nil {
		return err
	}
	return nil
}

//Login will login the Admin and will create a session
func LoginByPassword(phoneNumber string, password string, Session *admin_session.Detail)(err error) {
	var admin Detail
	err = GetAdminByPhoneNumber(&admin, phoneNumber)
	if err == nil {
		if admin.Password != password {
			err = errors.New("Password is Wrong")
		} else {
			Session.AdminId=admin.ID
			err = admin_session.CreateSession(Session)
		}
	}
	return err
}