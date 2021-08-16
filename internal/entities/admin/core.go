package admin

import (
	"errors"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/OTP"
	"github.com/freshpay/internal/entities/admin/admin_session"
	"github.com/freshpay/utilities"
)


//SignUp will be used to create a admin on signup
func SignUp(admin *Detail) (err error){
	phoneNumber :=admin.PhoneNumber

	//Validate the input
	err= ValidateInput(admin)
	if err!=nil{
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
	err= utilities.GetEncryption(admin.Password,&passwordHash)
	if err!=nil{
		return err
	}
	admin.Password= passwordHash


	//now create the admin
	admin.ID= utilities.CreateID(Prefix,IDLengthExcludingPrefix)

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

//Login will login the Admin and will create a admin _session
func LoginByPassword(phoneNumber string, password string, Session *admin_session.Detail, admin *Detail)(err error) {
	err = GetAdminByPhoneNumber(admin, phoneNumber)
	if err!=nil{
		return errors.New("Phone Number is wrong or not registered")
	}
	if !admin.IsVerified {
		return errors.New("Phone Number is not verified, please signup again")
	}
	if  !utilities.MatchPassword(password,admin.Password){
		return errors.New("Password is Wrong")
	}

	err=admin_session.GetActiveSessionByAdminId(Session,admin.ID)
	if err==nil{
		return nil
	}
	Session.AdminId = admin.ID
	err = admin_session.CreateSession(Session)
	return err
}

//Login By Using OTP
func LoginByOTP(PhoneNumber string)(err error){
	return SendOTPToRegisteredNumber(PhoneNumber)
}

//Login By OTP Verification
func LoginByOTPVerification(otp OTP.Detail,Session *admin_session.Detail, Admin *Detail)(err error){
	err=OTP.VerifyOTP(otp)
	if err!=nil{
		return err
	}
	err=GetAdminByPhoneNumber(Admin,otp.PhoneNumber)
	if err!=nil{
		return err
	}
	err=admin_session.GetActiveSessionByAdminId(Session,Admin.ID)
	if err==nil{
		return nil
	}
	Session.AdminId = Admin.ID
	err = admin_session.CreateSession(Session)
	return err
}
//Reset Password Using OTP at the registered Phone Number
func ResetPasswordByOTP(PhoneNumber string)(err error){
	return SendOTPToRegisteredNumber(PhoneNumber)
}

//Reset Password By OTP Verification
func ResetPasswordByOTPVerification(otp OTP.Detail,password string) (err error) {
	err=OTP.VerifyOTP(otp)
	if err!=nil{
		return err
	}
	var admin Detail
	err=GetAdminByPhoneNumber(&admin,otp.PhoneNumber)
	if err!=nil{
		return err
	}
	var passwordHash string
	err=utilities.GetEncryption(password,&passwordHash)
	if err!=nil {
		return err
	}
	admin.Password=passwordHash
	err=UpdateAdmin(&admin)
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

//Validate the Input
func ValidateInput(admin *Detail) (err error){
	err=utilities.ValidatePhoneNumber(admin.PhoneNumber)
	return err
}

//Function to send OTP Registered Number
func SendOTPToRegisteredNumber(PhoneNumber string)(err error){
	err=utilities.ValidatePhoneNumber(PhoneNumber)
	if err!=nil{
		return err
	}
	var tempAdmin Detail
	err=GetAdminByPhoneNumber(&tempAdmin,PhoneNumber)
	if err!=nil{
		return errors.New("Phone Number is not registered, Please Signup first")
	}
	err=OTP.SendOTP(PhoneNumber)
	return err
}