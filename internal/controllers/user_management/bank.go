package user_management

import (
	"github.com/freshpay/internal/entities/user_management/bank"
	"github.com/gin-gonic/gin"

	"net/http"
)


//AddBankAccount will add the bank account to the user
func AddBankAccount(c *gin.Context){
	userId :=c.GetString("userId")
	var bankAccount bank.Detail
	c.BindJSON(&bankAccount)
	err:=bank.CreateBank(&bankAccount,userId)
	if err!=nil{
		c.AbortWithStatus(http.StatusBadRequest)
	} else{
		c.JSON(http.StatusOK,bankAccount)
	}
}

func GetAllBankAccountByUserId(c *gin.Context){
	userId :=c.GetString("userId")
	var bankAccount []bank.Detail
	err:=bank.GetAllBankAccountsByUserId(&bankAccount,userId)
	if err!=nil{
		c.AbortWithStatus(http.StatusBadRequest)
	} else{
		c.JSON(http.StatusOK,bankAccount)
	}
}

