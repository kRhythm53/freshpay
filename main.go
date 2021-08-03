package main

import (
	"fmt"
	"github.com/kshitij-nawandar9/freshpay/internal/Config"
	"github.com/kshitij-nawandar9/freshpay/internal/entities/Admin"
	"github.com/kshitij-nawandar9/freshpay/internal/entities/Campaigns"
	"github.com/kshitij-nawandar9/freshpay/internal/entities/Complaints"
	"github.com/kshitij-nawandar9/freshpay/internal/entities/Payments"
	"github.com/kshitij-nawandar9/freshpay/internal/entities/UserManagement/Bank"
	"github.com/kshitij-nawandar9/freshpay/internal/entities/UserManagement/Beneficiary"
	"github.com/kshitij-nawandar9/freshpay/internal/entities/UserManagement/Session"
	"github.com/kshitij-nawandar9/freshpay/internal/entities/UserManagement/User"
	"github.com/kshitij-nawandar9/freshpay/internal/entities/UserManagement/Wallet"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var err error

func main() {
	//Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB, err = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}

	//defer Config.DB.Close()
	Config.DB.AutoMigrate(&Payments.Payments{},&Payments.Transactions{})
	Config.DB.AutoMigrate(&Campaigns.Campaign{},&Complaints.Complaint{})
	Config.DB.AutoMigrate(&Admin.Detail{},&Bank.Detail{},&User.Detail{},&Beneficiary.Detail{},&Session.Detail{},&Wallet.Detail{})
	//r := Routes.SetupRouter()
	////running
	//r.Run()

}
