package payments

import (
	"errors"
	"github.com/freshpay/internal/constants"
	"github.com/freshpay/internal/entities/payments/utilities"
	"github.com/freshpay/internal/entities/user_management/bank"
	"github.com/freshpay/internal/entities/user_management/wallet"
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
	err=ValidityCheck(payment)
	if err!=nil{
		return err
	}
	InputPaymentsChannel <-payment

	return AddPaymentToDB(payment)
}

func GetPaymentByID(payment *Payments, id string) (err error) {
	return GetPaymentByIDFromDB(payment,id)
}

func GetPaymentsByTime(payments *[]Payments, from string,to string,userID string ) (err error) {
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
	return GetPaymentByTimeFromDB(payments,startTime,endTime,userID)
}

func UpdatePayment(payment *Payments) (err error) {
	Eligibility(payment *Payment,userID string)
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

func GetPaymentType(payment *Payments) string{
	if strings.HasPrefix(payment.SourceId, constants.WalletPrefix){
		if strings.HasPrefix(payment.DestinationId, constants.WalletPrefix){
			return constants.PaymentTypeWalletTransfer
		}else{
			return constants.PaymentTypeBankWithdrawal
		}
	}else{
		return constants.PaymentTypeAddToWallet
	}
}

func ValidityCheck(payment *Payments) (err error ){
	var balance int
	if strings.HasPrefix(payment.SourceId, constants.WalletPrefix){
		var Source *wallet.Detail
		err := wallet.GetWalletById(Source, payment.SourceId)
		if err != nil {
			return err
		}
		balance=Source.Balance
	} else{
		var Source *bank.Detail
		err := bank.GetBankById(Source, payment.SourceId)
		if err != nil {
			return err
		}
	}
	if strings.HasPrefix(payment.DestinationId, constants.WalletPrefix){
		var Destination *wallet.Detail
		err := wallet.GetWalletById(Destination, payment.SourceId)
		if err != nil {
			return err
		}
	} else{
		var Destination *bank.Detail
		err := bank.GetBankById(Destination, payment.SourceId)
		if err != nil {
			return err
		}
	}
	if balance<int(payment.Amount){
		return errors.New("low wallet balance")
	}
	return nil
}

func GenerateID() string{
	return utilities.RandomString(14,constants.IDPrefix)
}

func InitiateRefund(amount int64,paymentID string)(err error){

}
