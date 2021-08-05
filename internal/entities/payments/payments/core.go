package payments

import (
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/payments/utilities"
	"strings"
)

var PaymentsChannel=make(chan *Payments,1000)

func AddPayments(payment *Payments) (err error) {
	if strings.HasPrefix(payment.SourceId,"wallt"){
		if strings.HasPrefix(payment.DestinationId,"wallt"){
			payment.Type="wallet transfer"
		}else{
			payment.Type="bank withdrawal"
		}
	}else{
		payment.Type="add to wallet"
	}
	payment.Status="processing"
	payment.ID="paymt_"+utilities.RandomString(14)
	PaymentsChannel<-payment
	if err = config.DB.Table("payments").Create(payment).Error; err != nil {
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
