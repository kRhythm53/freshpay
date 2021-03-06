package user_management

import "C"
import (
	"github.com/freshpay/internal/entities/Error"
	"github.com/freshpay/internal/entities/OTP"
	"github.com/freshpay/internal/entities/user_management/user"
	"github.com/freshpay/internal/entities/user_management/user_session"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	SignUp will signup the user
*/
func SignUp(c *gin.Context) {
	var User user.Detail
	c.BindJSON(&User)
	err := user.SignUp(&User)
	if err != nil {
		c.JSON(http.StatusBadRequest,Error.Detail{
			"BAD_REQUEST_ERROR","Failed",err.Error(),"buisness",
			"BAD REQUEST","NA","{}",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Entity":user.EntityName,
			"Status":"OTP Verification Not Done",
			"ID":User.ID,
			"Name":User.Name,
			"PhoneNumber":User.PhoneNumber,
			"Email":User.Email,
		})
	}
}


//LoginByPassword will login the user by Password
func LoginByPassword(c *gin.Context){
	var loginInfo user.Detail
	c.BindJSON(&loginInfo)
	var Session user_session.Detail
	var User user.Detail
	err:=user.LoginByPassword(loginInfo.PhoneNumber,loginInfo.Password,&Session,&User)
	if err!=nil{
		c.JSON(401,Error.Detail{
			"UnAuthorized","Failed",err.Error(),"buisness",
			"Wrong Login Details","NA","{}",
		})
	} else{
		c.Writer.Header().Set("session_id",Session.ID)
		c.JSON(http.StatusOK,gin.H{
			"Entity": user.EntityName,
			"Status":"Success",
			"Message":"Login Successfully",
			"User": gin.H{
				"ID":User.ID,
				"Name":User.Name,
				"PhoneNumber":User.PhoneNumber,
				"Email":User.Email,
			},
		})
	}
}

func ResetPasswordByOTP(c *gin.Context){
	PhoneNumber:=c.Params.ByName("phone_number")
	err:=user.ResetPasswordByOTP(PhoneNumber)
	if err!=nil{
		c.JSON(http.StatusBadRequest,Error.Detail{
			"BAD_REQUEST_ERROR","Failed",err.Error(),"buisness",
			"BAD REQUEST","NA","{}",
		})
	} else{
		c.JSON(http.StatusOK, gin.H{
			"Entity":"Reset Passowrd",
			"Status":"OTP sent successfully,Please verify the otp",
			"PhoneNumber":PhoneNumber,
		})
	}
}

func ResetPasswordByOTPVerification(c *gin.Context){
	password:=c.Params.ByName("password")
	var otp OTP.Detail
	c.BindJSON(&otp)
	err:=user.ResetPasswordByOTPVerification(otp,password)
	if err!=nil{
		c.JSON(http.StatusBadRequest,Error.Detail{
			"BAD_REQUEST_ERROR","Failed",err.Error(),"buisness",
			"Wrong Details","NA","{}",
		})
	} else{
		c.JSON(http.StatusOK, gin.H{
			"Entity":"User",
			"Status":"Success",
			"Message":"Password has been reset successfully",
			"PhoneNumber":otp.PhoneNumber,
		})
	}
}


//LoginByOTP
func LoginByOTP(c *gin.Context){
	phoneNumber:=c.Params.ByName("phone_number")
	err:=user.LoginByOTP(phoneNumber)
	if err!=nil{
		c.JSON(http.StatusBadRequest,Error.Detail{
			"BAD_REQUEST_ERROR","Failed",err.Error(),"buisness",
			"BAD REQUEST","NA","{}",
		})
	} else{
		c.JSON(http.StatusOK, gin.H{
			"Entity":"User",
			"Status":"OTP sent successfully,Please verify the otp to login",
			"PhoneNumber":phoneNumber,
		})
	}
}


//LoginByOTPVerification
func LoginByOTPVerification(c *gin.Context){
	var otp OTP.Detail
	c.BindJSON(&otp)
	var Session user_session.Detail
	var User user.Detail
	err:=user.LoginByOTPVerification(otp,&Session,&User)
	if err!=nil{
		c.JSON(401,Error.Detail{
			"UnAuthorized","Failed",err.Error(),"buisness",
			"Wrong Login Details","NA","{}",
		})
	}else{
		c.Writer.Header().Set("session_id",Session.ID)
		c.JSON(http.StatusOK,gin.H{
			"Entity": user.EntityName,
			"Status":"Success",
			"Message":"Login Successfully",
			"User": gin.H{
				"ID":User.ID,
				"Name":User.Name,
				"PhoneNumber":User.PhoneNumber,
				"Email":User.Email,
			},
		})
	}
}