package wallet

import (
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/user_management/utilities"
)

//CreateWallet will create a new wallet
func CreateWallet(userId string)(err error){
	var wallet Detail
	wallet.ID=utilities.CreateID(Prefix,14)
	wallet.UserId=userId
	if err=config.DB.Create(&wallet).Error; err!=nil{
		return err
	}
	return nil
}

//GetWalletById will return the wallet by using the id
func GetWalletById(wallet *Detail,id string)(err error){
	if err = config.DB.Where("ID = ?",id).First(wallet).Error; err != nil {
		return err
	}
	return nil
}

func GetWalletByUserId(wallet *Detail,userId string)(err error){
	if err=config.DB.Table("wallet").Where("user_id = ?",userId).First(wallet).Error;err!=nil{
		return err
	}
	return nil
}

func GetWalletBalanceByUserId(wallet *Detail,userId string)(err error){
	if err=config.DB.Where("user_id = ?",userId).First(wallet).Error;err!=nil{
		return err
	}
	wallet.UserId=""
	wallet.ID=""
	return nil
}

func UpdateWalletBalance(walletID string,amount int64){
	var Wallet Detail
	err := GetWalletById(&Wallet, walletID)
	if err != nil {
		return
	}
	Wallet.Balance+=int(amount)
	config.DB.Table("wallet").Save(Wallet)
}
