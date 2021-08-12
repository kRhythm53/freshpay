package user_management

import (
	"github.com/freshpay/internal/entities/Error"
	"github.com/freshpay/internal/entities/user_management/beneficiary"
	"github.com/gin-gonic/gin"
	"net/http"
)

//AddBankAccount will add the bank account to the user
func AddBeneficiary(c *gin.Context){
	userId :=c.GetString("userId")
	var Beneficiary beneficiary.Detail
	c.BindJSON(&Beneficiary)
	err:=beneficiary.CreateBeneficiary(&Beneficiary,userId)
	if err!=nil{
		c.JSON(400,Error.Detail{"BAD_REQUEST_ERROR","Failed",err.Error(),
			"buisness","Input Validation Failed","NA","{}"})
	} else{
		c.JSON(http.StatusOK, gin.H{
			"Status":"Success",
			"Entity":beneficiary.EntityName,
			"ID":Beneficiary.ID,
			"BankName":Beneficiary.BankName,
			"AccountNumber":Beneficiary.AccountNumber,
			"IFSCCode":Beneficiary.IFSCCode,
		})
	}
}

func GetAllBeneficiaryByUserId(c *gin.Context){
	userId :=c.GetString("userId")
	var Beneficiary []beneficiary.Detail
	err:=beneficiary.GetAllBeneficiaryAccountsByUserId(&Beneficiary,userId)
	if err!=nil{
		c.JSON(500,Error.Detail{"INTERNAL_SERVER_ERROR","Failed",err.Error(),
			"Internal","","NA","{}"},
		)
	} else{
		c.JSON(http.StatusOK, gin.H{
			"Status":        "Success",
			"Entity":        beneficiary.EntityName,
			"BankDetails":    Beneficiary,
		})
	}
}
