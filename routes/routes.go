package routes

import (
	campaigns "github.com/freshpay/internal/controllers/campaign"
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
		grp1.GET("/:payments_id", payments.GetPaymentByID)
		grp1.GET("/",payments.GetPaymentsByTime)
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
	grp5:= r.Group("/admin")
	{
		grp5.POST("signup",admin_management.SignUp)
		grp5.POST("signup/otp/verification",otp_verification.VerifyOTPAdmin)
		grp5.POST("signin",admin_management.LoginByPassword)
	}
	grp3:= r.Group("/campaigns")
	{
		grp3.POST("/",campaigns.CreateCampaign)
		grp3.GET("/",campaigns.GetCampaign)
		grp3.GET("/:campaign_id",campaigns.GetCampaignByID)
		grp3.PATCH("/:campaign_id",campaigns.UpdateCampaign)
	}
	return r
}
