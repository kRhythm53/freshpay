package routes

import (
	"github.com/freshpay/internal/controllers/admin_management"
	"github.com/freshpay/internal/controllers/payments"
	"github.com/freshpay/internal/controllers/user_management"
	"github.com/freshpay/middleware"
	"github.com/gin-gonic/gin"
)
//SetupRouter ... Configure routes

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Authenticate)
	grp1 := r.Group("/payments")
	{
		grp1.POST("", payments.AddPayments)
	}

	grp2:= r.Group("/users")
	{
		grp2.POST("signup",user_management.SignUp)
		grp2.POST("signin",user_management.LoginByPassword)
		grp2.POST("bankaccount",user_management.AddBankAccount)
		grp2.GET("bankaccounts",user_management.GetAllBankAccountByUserId)

		grp2.POST("beneficiary",user_management.AddBeneficiary)
		grp2.GET("beneficiary",user_management.GetAllBeneficiaryByUserId)

		grp2.GET("balance",user_management.GetWalletBalance)
	}

	grp3:= r.Group("/admin")
	{
		grp3.POST("signup",admin_management.SignUp)
		grp3.POST("signin",admin_management.LoginByPassword)
	}
	return r
}
