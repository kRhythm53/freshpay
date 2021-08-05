package payments

import (
	"fmt"
	"github.com/freshpay/internal/config"
	"strings"
)

func AddPayments(payments *Payments) (err error) {
	if strings.HasPrefix(payments.SourceId,"wallt"){
		if strings.HasPrefix(payments.DestinationId,"wallt"){
			payments.Type="wallet transfer"
		}else{
			payments.Type="bank withdrawal"
		}
	}else{
		payments.Type="add to wallet"
	}
	payments.Status="processing"

	payments.ID="paymt_"+RandomString(14)
	fmt.Println(payments)
	if err = config.DB.Table("payments").Create(payments).Error; err != nil {
		return err
	}
	return nil
}

func GetPaymentByID(payment *Payments, id string) (err error) {
	if err = config.DB.Table("payments").Where("id = ?", id).First(payment).Error; err != nil {
		return err
	}
	return nil
}

func GetPaymentsByTime(payments *[]Payments, startTime int64,endTime int64 ) (err error) {

	if err = config.DB.Table("payments").Where("created_at > ? AND created_at < ?", startTime, endTime).Find(payments).Error; err != nil {
		return err
	}
	return nil
}
