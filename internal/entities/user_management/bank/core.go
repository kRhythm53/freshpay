package bank

import (
	"fmt"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/user_management/utilities"
)

//CreateBank will create a new bank
func CreateBank(bank *Detail, userId string) (err error) {
	bank.ID = utilities.CreateID(Prefix, 14)
	bank.UserId = userId
	fmt.Println("banK: ", bank)
	if err = config.DB.Create(bank).Error; err != nil {
		return err
	}
	return nil
}

//GetBankById will return the bank by using the id
func GetBankById(bank *Detail, id string) (err error) {
	if err = config.DB.Table("bank").Where("ID = ?", id).First(bank).Error; err != nil {
		return err
	}
	return nil
}

//GetAllBankAccountsByUserId will return all the bank accounts attached to a user
func GetAllBankAccountsByUserId(bank *[]Detail, user_id string) (err error) {
	if err = config.DB.Where("user_id = ?", user_id).Find(bank).Error; err != nil {
		return err
	}
	return nil
}
