package payments

import (
	"errors"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/payments/utilities"
	"strconv"
	"strings"
	"time"
)

var InputPaymentsChannel =make(chan *Payments,1000)
var ResultsPaymentsChannel=make(chan *Payments,1000)

func AddPayments(payment *Payments) (err error) {
	payment.Type=GetPaymentType(payment)
	payment.Status="processing"
	payment.ID=GenerateID()
	//err=ValidityCheck(payment)
	//if err!=nil{
	//	return err
	//}
	InputPaymentsChannel <-payment

	return AddPaymentToDB(payment)
}

func GetPaymentByID(payment *Payments, id string) (err error) {
	return GetPaymentByIDFromDB(payment,id)
}

func GetPaymentsByTime(payments *[]Payments, from string,to string ) (err error) {
	var startTime, endTime int64
	if from == "" {
		startTime = time.Now().Unix()
	} else {
		startTime, err = strconv.ParseInt(from, 10, 64)
		if err != nil {
			return
		}
	}
	if to == "" {
		endTime = time.Now().Unix()
	} else {
		endTime, err = strconv.ParseInt(to, 10, 64)
	}
	return GetPaymentByTimeFromDB(payments,startTime,endTime)
}

func UpdatePayment(payment *Payments) (err error) {
	return UpdatePaymentToDB(payment)
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
	if strings.HasPrefix(searchID,WalletPrefix){
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
	if strings.HasPrefix(payment.SourceId,WalletPrefix){
		if strings.HasPrefix(payment.DestinationId,WalletPrefix){
			return PaymentTypeWalletTransfer
		}else{
			return PaymentTypeBankWithdrawal
		}
	}else{
		return PaymentTypeAddToWallet
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

func GenerateID() string{
	return IDPrefix+utilities.RandomString(14)
}