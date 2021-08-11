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
		c.JSON(500,gin.H{
			"Code": "BAD_REQUEST_ERROR",
			"Status":"failed",
			"Description":err.Error(),
			"Source": "business",
			"Reason": "input_validation_failed",
			"Step": "NA",
			"Metadata":"{}",
		})
	} else{
		c.JSON(http.StatusOK, gin.H{
			"Status":"success",
			"Entity":bank.EntityName,
			"ID":bankAccount.ID,
			"BankName":bankAccount.BankName,
			"AccountNumber":bankAccount.AccountNumber,
			"IFSCCode":bankAccount.IFSCCode,
		})
	}
}

func GetAllBankAccountByUserId(c *gin.Context){
	userId :=c.GetString("userId")
	var bankAccounts []bank.Detail
	err:=bank.GetAllBankAccountsByUserId(&bankAccounts,userId)
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
		c.JSON(http.StatusOK, gin.H{
			"Status":        "success",
			"Entity":        bank.EntityName,
			"BankDetails":    bankAccounts,
		})
	}
}