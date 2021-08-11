package bank

import (
	"errors"
	"fmt"
	"github.com/freshpay/internal/config"
	utilities2 "github.com/freshpay/utilities"
)

//CreateBank will create a new bank
func CreateBank(bank *Detail, userId string) (err error) {
	fmt.Println("accountNumber: ",bank.AccountNumber);
	if len(bank.AccountNumber)<9 || len(bank.AccountNumber)>18{
		fmt.Println("x: " ,len(bank.AccountNumber))
		err=errors.New("Number of characters in account number should be b/w 9 and 18")
		return err
	}
	if len(bank.IFSCCode) !=11{
		err=errors.New("Number of characters in IFSCCode should be 11")
		return err
	}
	bank.ID = utilities2.CreateID(Prefix, IDLengthExcludingPrefix)
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
