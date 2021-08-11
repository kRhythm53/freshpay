package payments

import (
	"errors"
	"github.com/freshpay/internal/config"
)

func AddPaymentToDB(payment *Payments) (err error) {
	if err := config.DB.Table("payments").Create(payment).Error; err != nil {
		return errors.New("failed payment because of server issues")
	}
	return nil
}
func GetPaymentByIDFromDB(payment *Payments, id string) (err error) {
	if err = config.DB.Table("payments").Where("id = ?", id).First(payment).Error; err != nil {
		return errors.New("payment not found")
	}
	return nil
}

func GetPaymentByTimeFromDB(payments *[]Payments, startTime int64, endTime int64, TransactionType string, WalletID string) (err error) {
	if TransactionType == "debit" {
		if err = config.DB.Table("payments").Where("created_at > ? AND created_at < ? AND source_id = ?", startTime, endTime, WalletID).Find(payments).Error; err != nil {
			return errors.New("payment not found")
		}
	} else if TransactionType == "credit"{
		if err = config.DB.Table("payments").Where("created_at > ? AND created_at < ? AND destination_id = ?", startTime, endTime, WalletID).Find(payments).Error; err != nil {
			return errors.New("payment not found")
		}
	} else{
		if err = config.DB.Table("payments").Where("created_at > ? AND created_at < ? ", startTime, endTime).Find(payments).Error; err != nil {
			return errors.New("payment not found")
		}
	}
	return nil
}

func UpdatePaymentToDB(payment *Payments) (err error) {
	if err=config.DB.Table("payments").Save(payment).Error;err!=nil{
		return errors.New("failed to update payment")
	}
	return nil
}
