package main

import (
	"fmt"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/admin"
	"github.com/freshpay/internal/entities/admin/admin_session"
	"github.com/freshpay/internal/entities/campaigns"
	"github.com/freshpay/internal/entities/complaints"
	"github.com/freshpay/internal/entities/payments/payments"
	"github.com/freshpay/internal/entities/payments/transactions"
	"github.com/freshpay/internal/entities/user_management/bank"
	"github.com/freshpay/internal/entities/user_management/beneficiary"
	"github.com/freshpay/internal/entities/user_management/user"
	"github.com/freshpay/internal/entities/user_management/user_session"
	"github.com/freshpay/internal/entities/user_management/wallet"
	"github.com/freshpay/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var err error

func main() {
	config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}

	err = config.DB.AutoMigrate(&payments.Payments{}, &transactions.Transactions{})
	if err != nil {
		return
	}
	err = config.DB.AutoMigrate(&campaigns.Campaign{}, &complaints.Complaint{})
	if err != nil {
		return
	}
	err = config.DB.AutoMigrate(&admin.Detail{}, &bank.Detail{}, &user.Detail{}, &beneficiary.Detail{}, &user_session.Detail{},
		&wallet.Detail{}, &admin_session.Detail{})
	if err != nil {
		return
	}

	go transactions.InitiateTransaction()
	go payments.PaymentReceiver()
	err = payments.CreateRzpAccount()
	if err != nil {
		return

	}

	
	r:=routes.SetupRouter()
	////running
	r.Run()

}
