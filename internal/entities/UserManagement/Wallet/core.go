package Wallet

import "github.com/kshitij-nawandar9/freshpay/internal/Config"

//CreateWallet will create a new wallet
func CreateWallet(wallet *Detail)(err error){
	if err=Config.DB.Create(wallet).Error; err!=nil{
		return err
	}
	return nil
}

//GetWalletById will return the Wallet by using the id
func GetWalletByID(wallet *Detail, entity_id string)(err error){
	if err = Config.DB.Where("entity_id = ?", entity_id).First(wallet).Error; err != nil {
		return err
	}
	return nil
}
