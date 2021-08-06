package user_management

import (
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
		c.AbortWithStatus(http.StatusBadRequest)
	} else{
		c.JSON(http.StatusOK,Beneficiary)
	}
}

func GetAllBeneficiaryByUserId(c *gin.Context){
	userId :=c.GetString("userId")
	var Beneficiary []beneficiary.Detail
	err:=beneficiary.GetAllBeneficiaryAccountsByUserId(&Beneficiary,userId)
	if err!=nil{
		c.AbortWithStatus(http.StatusBadRequest)
	} else{
		c.JSON(http.StatusOK,Beneficiary)
	}
}
