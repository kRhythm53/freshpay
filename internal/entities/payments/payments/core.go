package payments

import (
	"errors"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/payments/utilities"
	"strings"
)

var InputPaymentsChannel =make(chan *Payments,1000)
var ResultsPaymentsChannel=make(chan *Payments,1000)

func AddPayments(payment *Payments) (err error) {
	payment.Type=GetPaymentType(payment)
	payment.Status="processing"
	payment.ID="paymt_"+utilities.RandomString(14)
	//err=ValidityCheck(payment)
	//if err!=nil{
	//	return err
	//}
	InputPaymentsChannel <-payment
	if err= config.DB.Table("payments").Create(payment).Error; err != nil {
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

func UpdatePayment(payment *Payments) (err error) {
	config.DB.Table("payments").Save(payment)
	return nil
}

func PaymentReceiver() {
	for {
		select {
		case payment:=<-ResultsPaymentsChannel:
			err := UpdatePayment(payment)
			if err != nil {
				return
			}
		}
	}
}

func GetUserID(searchID string) (err error,userID string) {
	var table string
	if strings.HasPrefix(searchID,"wallt"){
		table="wallet"
	}else{
		table="bank"
	}
	if err = config.DB.Table(table).Where("id = ?", searchID).First(userID).Error; err != nil {
		return err,""
	}
	return nil,userID
}

func GetPaymentType(payment *Payments) string{
	if strings.HasPrefix(payment.SourceId,"wallt"){
		if strings.HasPrefix(payment.DestinationId,"wallt"){
			return "wallet transfer"
		}else{
			return "bank withdrawal"
		}
	}else{
		return "add to wallet"
	}
}

func ValidityCheck(payment *Payments) (err error ){
	if err,_:=GetUserID(payment.SourceId);err!=nil{
		return errors.New("source not found")
	}else if err,_:=GetUserID(payment.DestinationId);err!=nil{
		return errors.New("destination not found")
	}
	return nil
}