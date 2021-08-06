package wallet

import (
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/user_management/utilities"
)

//CreateWallet will create a new wallet
func CreateWallet(userId string)(err error){
	var wallet Detail
	wallet.ID=utilities.CreateID(14)
	wallet.Balance=0
	wallet.Currency="INR"
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
	if err=config.DB.Where("user_id = ?",userId).First(wallet).Error;err!=nil{
		return err;
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
