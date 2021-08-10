package otp_verification

import (
	"github.com/freshpay/internal/entities/OTP"
	"github.com/freshpay/internal/entities/admin"
	"github.com/freshpay/internal/entities/user_management/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyOTPUser(c *gin.Context){
	var otp OTP.Detail
	c.BindJSON(&otp)

	err:= OTP.VerifyOTP(otp)
	if err!=nil{
		c.AbortWithError(400,err)
	} else{
		err=user.SetVerifiedUserByPhoneNumber(otp.PhoneNumber)
		if err!=nil{
			c.AbortWithError(400,err)
		} else{
			c.JSON(http.StatusOK,"Your account has been registered successfully")
		}
	}
}
func VerifyOTPAdmin(c *gin.Context){
	var otp OTP.Detail
	c.BindJSON(&otp)

	err:= OTP.VerifyOTP(otp)
	if err!=nil{
		c.AbortWithError(400,err)
	} else{
		err=admin.SetVerifiedAdminByPhoneNumber(otp.PhoneNumber)
		if err!=nil{
			c.AbortWithError(400,err)
		} else{
			c.JSON(http.StatusOK,"Your account has been registered successfully")
		}
	}
}
