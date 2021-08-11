package user_management

import (
	"github.com/freshpay/internal/entities/user_management/user"
	"github.com/freshpay/internal/entities/user_management/wallet"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetWalletBalance(c *gin.Context){
	userId :=c.GetString("userId")
	var Wallet wallet.Detail
	err:=wallet.GetWalletByUserId(&Wallet,userId)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Code": "Internal_Server_Error",
			"Status":"failed",
			"Description":err.Error(),
			"Source": "internal",
			"Reason": "",
			"Step": "NA",
			"Metadata":"{}",
		})
	} else{
		c.JSON(200,gin.H{
			"Entity":wallet.EntityName,
			"Status":"success",
			"ID":Wallet.ID,
			"balance": Wallet.Balance,
			"currency": Wallet.Currency,
		})
	}
}

func GetWalletByPhoneNumber(c *gin.Context){
	phoneNumber:=c.Params.ByName("phone_number")
	var User user.Detail
	err:=user.GetUserByPhoneNumber(&User,phoneNumber)
	if err==nil{
		var Wallet wallet.Detail
		err=wallet.GetWalletByUserId(&Wallet,User.ID)
		if err==nil{
			c.JSON(http.StatusOK, gin.H{
				"Entity":wallet.EntityName,
				"Status":"success",
				"ID":Wallet.ID,
				"Name":User.Name,
				"PhoneNumber":phoneNumber,
			})
		}
	}
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{
			"Code": "BAD_REQUEST_ERROR",
			"Status":"failed",
			"Description":err.Error(),
			"Source": "business",
			"Reason": "The Phone Number isn't registered",
			"Step": "NA",
			"Metadata":"{}",
		})
	}
}
