package transactions

import "github.com/freshpay/internal/config"

func AddTransactionToDB(transaction Transactions)(err error){
	if err := config.DB.Table("transactions").Create(transaction).Error; err != nil {
		return err
	}
	return nil
}
