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
	if err==nil{
		err=user.SetVerifiedUserByPhoneNumber(otp.PhoneNumber)
		if err==nil{
			c.JSON(http.StatusOK, gin.H{
				"Entity":user.EntityName,
				"Status":"Success",
				"Message":"User Account Registered Succesfully",
				"PhoneNumber":otp.PhoneNumber,
			})
		}
	}
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"Code": "BAD_REQUEST_ERROR",
			"Status":"failed",
			"Description":err.Error(),
			"Source": "business",
			"Reason": "Wrong OTP",
			"Step": "NA",
			"Metadata":"{}",
		})
	}
}
func VerifyOTPAdmin(c *gin.Context){
	var otp OTP.Detail
	c.BindJSON(&otp)

	err:= OTP.VerifyOTP(otp)
	if err==nil{
		err=admin.SetVerifiedAdminByPhoneNumber(otp.PhoneNumber)
		if err==nil{
			c.JSON(http.StatusOK, gin.H{
				"Entity":admin.EntityName,
				"Status":"Success",
				"Message":"User Account Registered Succesfully",
				"PhoneNumber":otp.PhoneNumber,
			})
		}
	}
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"Code": "BAD_REQUEST_ERROR",
			"Status":"failed",
			"Description":err.Error(),
			"Source": "business",
			"Reason": "Wrong OTP or phone number",
			"Step": "NA",
			"Metadata":"{}",
		})
	}
}
