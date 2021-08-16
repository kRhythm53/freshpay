package user_management

import (
	"github.com/freshpay/internal/entities/Error"
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
		c.JSON(500,Error.Detail{"INTERNAL_SERVER_ERROR","Failed",err.Error(),
			"Internal","","NA","{}"},
		)
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
	if err!=nil{
		c.JSON(400,Error.Detail{"BAD_REQUEST_ERROR","Failed",err.Error(),
			"buisness","The Phone Number is wrong or isn't registered","NA","{}"},
		)
		c.Abort()
		return
	}
	var Wallet wallet.Detail
	err=wallet.GetWalletByUserId(&Wallet,User.ID)
	if err!=nil {
		c.JSON(400,Error.Detail{"BAD_REQUEST_ERROR","Failed",err.Error(),
			"buisness","The Phone Number is wrong or isn't registered","NA","{}"},
		)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Entity":wallet.EntityName,
		"Status":"success",
		"ID":Wallet.ID,
		"Name":User.Name,
		"PhoneNumber":phoneNumber,
	})

}
