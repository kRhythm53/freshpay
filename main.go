package main

import (
	"fmt"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/admin"
	"github.com/freshpay/internal/entities/campaigns"
	"github.com/freshpay/internal/entities/complaints"
	payments2 "github.com/freshpay/internal/entities/payments/payments"
	"github.com/freshpay/internal/entities/payments/transactions"
	"github.com/freshpay/internal/entities/user_management/bank"
	"github.com/freshpay/internal/entities/user_management/beneficiary"
	"github.com/freshpay/internal/entities/user_management/session"
	"github.com/freshpay/internal/entities/user_management/user"
	"github.com/freshpay/internal/entities/user_management/wallet"
	"github.com/freshpay/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var err error

func main() {
	//config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}

	go transactions.InitiateTransaction()
	go payments2.PaymentReceiver()

	//defer config.DB.Close()
	config.DB.AutoMigrate(&payments2.Payments{},&transactions.Transactions{})
	config.DB.AutoMigrate(&campaigns.Campaign{},&complaints.Complaint{})
	config.DB.AutoMigrate(&admin.Detail{},&bank.Detail{},&user.Detail{},&beneficiary.Detail{},&session.Detail{},&wallet.Detail{})
	r:=routes.SetupRouter()
	//////running
	//r.Run()

}
