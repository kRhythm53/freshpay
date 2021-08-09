package phonenumber_verification

import (
	"errors"
	"github.com/freshpay/internal/entities/admin"
	"github.com/freshpay/internal/entities/user_management/user"
	"github.com/gin-gonic/gin"
)

type Phone struct{
	PhoneNumber string
}
func VerifyPhoneNumber(c *gin.Context) (err error){
	var temp Phone
	c.BindJSON(&temp)
	phoneNumber:=temp.PhoneNumber
	if(c.FullPath()=="/users/signup"){
		err= userNumberExist(phoneNumber)
	} else{
		err= adminNumberExist(phoneNumber)
	}
	if err==nil{
		return errors.New("Phone Number is already registered")
	} else{
		err= Verify(phoneNumber)
		return err
	}
}

func userNumberExist(phoneNumber string) (err error){
	var userTemp user.Detail
	err=user.GetUserByPhoneNumber(&userTemp,phoneNumber)
	return err

}
func adminNumberExist(phoneNumber string)(err error){
	var adminTemp admin.Detail
	err=admin.GetAdminByPhoneNumber(&adminTemp,phoneNumber)
	return err
}
