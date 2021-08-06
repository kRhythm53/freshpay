package user_management

import (
	"github.com/freshpay/internal/entities/user_management/wallet"
	"github.com/gin-gonic/gin"
)

func GetWalletBalance(c *gin.Context){
	userId :=c.GetString("userId")
	var Wallet wallet.Detail
	err:=wallet.GetWalletBalanceByUserId(&Wallet,userId)
	if err!=nil{
		c.AbortWithError(400,err)
	} else{
		c.JSON(200,gin.H{
			"balance": Wallet.Balance,
			"currency": Wallet.Currency,
		})
	}
}
