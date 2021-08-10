package routes

import (
	"github.com/freshpay/internal/controllers/admin_management"
	campaigns "github.com/freshpay/internal/controllers/campaign"
	"github.com/freshpay/internal/controllers/complaints"
	"github.com/freshpay/internal/controllers/otp_verification"
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
		grp1.GET("/", payments.GetPaymentsByTime)
	}

	grp2 := r.Group("/users")
	{
		grp2.POST("signup", user_management.SignUp)
		grp2.POST("signup/otp/verification", otp_verification.VerifyOTPUser)
		grp2.POST("signin", user_management.LoginByPassword)
		grp2.POST("bankaccount", user_management.AddBankAccount)
		grp2.GET("bankaccounts", user_management.GetAllBankAccountByUserId)

		grp2.POST("beneficiary", user_management.AddBeneficiary)
		grp2.GET("beneficiary", user_management.GetAllBeneficiaryByUserId)

		grp2.GET("balance", user_management.GetWalletBalance)
		grp2.POST("complaint", complaints.CreateComplaint)
	}

	grp3 := r.Group("/campaigns")
	{
		grp3.POST("/", campaigns.CreateCampaign)
		grp3.GET("/", campaigns.GetCampaign)
		grp3.GET("/:campaign_id", campaigns.GetCampaignByID)
		grp3.PATCH("/:campaign_id", campaigns.UpdateCampaign)
	}

	grp5 := r.Group("/admin")
	{
		grp5.POST("signup",admin_management.SignUp)
		grp5.POST("signup/otp/verification",otp_verification.VerifyOTPAdmin)
		grp5.POST("signin",admin_management.LoginByPassword)
		grp5.GET("complaints",complaints.GetComplaints)
		grp5.GET("active_complaints",complaints.GetActiveComplaints)
		grp5.GET("complaint/:complaint_id",complaints.GetComplaintById)
		grp5.PATCH("complaint/:complaint_id",complaints.UpdateComplaintById)
	}
	return r
}
