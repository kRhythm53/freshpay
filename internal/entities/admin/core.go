package admin

import (
	"errors"
	"fmt"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/OTP"
	"github.com/freshpay/internal/entities/admin/admin_session"
	"github.com/freshpay/internal/entities/user_management/utilities"
)


//SignUp will be used to create a admin on signup
func SignUp(admin *Detail) (err error){
	phoneNumber :=admin.PhoneNumber
	/*
	    Validate the phoneNumber-> Phone number is of 10 digits and only have 0-9 characters
	 */
	if len(phoneNumber)!=10 || phoneNumber[0]=='0'{
		err=errors.New("phone number should be 10 digit long")
		return err
	}
	if !utilities.IsNumeric(phoneNumber){
		err=errors.New("Phone number can contain characters 0-9")
		return err
	}


	/*
	   Make sure PhoneNumber doesn't exist
	*/
	var adminTemp Detail
	err=GetAdminByPhoneNumber(&adminTemp,phoneNumber)
	if err==nil && adminTemp.IsVerified{
		err=errors.New("Phone Number is already registered")
		return err
	} else if err==nil{
		err=DeleteAdmin(&adminTemp)
		if err!=nil{
			return err
		}
	}

	err= OTP.SendOTP(phoneNumber)
	if err!=nil{
		return err
	}

	/*
	 Encrypt the password
	*/
	var passwordHash string
	err=utilities.GetEncryption(admin.Password,&passwordHash)
	if err!=nil{
		return err
	}
	admin.Password= passwordHash


	//now create the admin
	admin.ID=utilities.CreateID(Prefix,IDLengthExcludingPrefix)

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

//Login will login the Admin and will create a user_session
func LoginByPassword(phoneNumber string, password string, Session *admin_session.Detail, admin *Detail)(err error) {
	err = GetAdminByPhoneNumber(admin, phoneNumber)

	if err == nil {
		if !admin.IsVerified{
			err=errors.New("Phone Number is not verified, please signup again")
			fmt.Println(err)
			/*
			   need to remove this line
			*/
			return err
		}
		if !utilities.MatchPassword(password,admin.Password) {
			err = errors.New("Password is Wrong")
		} else {
			err=admin_session.GetActiveSessionByAdminId(Session,admin.ID)
			if err==nil{
				return nil
			}
			Session.AdminId=admin.ID
			err = admin_session.CreateSession(Session)
		}
	} else{
		err=errors.New("Phone Number is wrong or not registered")
	}
	return err
}




//set verified admin by phone number
func SetVerifiedAdminByPhoneNumber(phoneNumber string)(err error){
	var admin Detail
	err = GetAdminByPhoneNumber(&admin, phoneNumber)
	if err==nil{
		admin.IsVerified=true
		err=UpdateAdmin(&admin)
	}
	return err
}

//update the admin
func UpdateAdmin(admin *Detail)(err error){
	err=config.DB.Save(admin).Error
	if err!=nil{
		return err
	}
	return nil
}


//Delete a admin
func DeleteAdmin(admin *Detail)(err error){
	err=config.DB.Where("id = ?",admin.ID).Delete(admin).Error
	return err
}