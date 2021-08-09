package payments

import (
	"errors"
	"fmt"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/constants"
	"github.com/freshpay/internal/entities/campaigns"
	"github.com/freshpay/internal/entities/payments/utilities"
	"github.com/freshpay/internal/entities/user_management/bank"
	"github.com/freshpay/internal/entities/user_management/user"
	"github.com/freshpay/internal/entities/user_management/wallet"
	"strconv"
	"strings"
	"time"
)

var InputPaymentsChannel = make(chan *Payments, 1000)
var ResultsPaymentsChannel = make(chan *Payments, 1000)

func AddPayments(payment *Payments) (err error) {
	payment.Type = GetPaymentType(payment)
	payment.Status = "processing"
	payment.ID = GenerateID()
	err=ValidityCheck(payment)
	if err!=nil{
		return err
	}
	InputPaymentsChannel <- payment

	return AddPaymentToDB(payment)
}

func GetPaymentByID(payment *Payments, id string) (err error) {
	return GetPaymentByIDFromDB(payment, id)
}

func GetPaymentsByTime(payments *[]Payments, from string, to string, userID string) (err error) {
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
	return GetPaymentByTimeFromDB(payments, startTime, endTime, userID)
}

func UpdatePayment(payment *Payments) (err error) {
	UpdateTransactionCount(payment)
	if payment.Type != "Cashback" || payment.Type != "Refund" {
		err2 := InitiateCashback(payment)
		if err2 != nil {
			return err2
		}
	}
	return UpdatePaymentToDB(payment)
}

func PaymentReceiver() {
	for {
		select {
		case payment := <-ResultsPaymentsChannel:
			err := UpdatePayment(payment)
			if err != nil {
				return
			}
		}
	}
}

func GetPaymentType(payment *Payments) string {
	if strings.HasPrefix(payment.SourceId, constants.WalletPrefix) {
		if strings.HasPrefix(payment.DestinationId, constants.WalletPrefix) {
			return constants.PaymentTypeWalletTransfer
		} else {
			return constants.PaymentTypeBankWithdrawal
		}
	} else {
		return constants.PaymentTypeAddToWallet
	}
}

func ValidityCheck(payment *Payments) (err error) {
	var balance int
	if strings.HasPrefix(payment.SourceId, constants.WalletPrefix) {
		var Source wallet.Detail
		err := wallet.GetWalletById(&Source, payment.SourceId)
		if err != nil {
			return err
		}
		balance = Source.Balance
		if balance < int(payment.Amount) {
			fmt.Println("low balance")
			return errors.New("low wallet balance")
		}
	} else {
		var Source bank.Detail
		err := bank.GetBankById(&Source, payment.SourceId)
		if err != nil {
			return err
		}
	}
	if strings.HasPrefix(payment.DestinationId, constants.WalletPrefix) {
		var Destination wallet.Detail
		err := wallet.GetWalletById(&Destination, payment.DestinationId)
		if err != nil {
			return err
		}
	} else {
		var Destination bank.Detail
		err := bank.GetBankById(&Destination, payment.DestinationId)
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateID() string {
	return utilities.RandomString(14, constants.IDPrefix)
}

func InitiateRefund(paymentID string, UserID string) (RefundID string, err error) {
	var RefundPayment Payments

	var payment *Payments
	err2 := GetPaymentByID(payment, paymentID)
	if err2 != nil {
		return "", err2
	}

	var RefundWallet wallet.Detail
	err3 := wallet.GetWalletByUserId(&RefundWallet, UserID)
	if err3 != nil {
		return "", err3
	}

	RefundPayment.ID = GenerateID()
	RefundPayment.Amount = payment.Amount
	RefundPayment.Currency = "INR"
	RefundPayment.SourceId = "wal_Mh5gqYDWlNBYWq"
	RefundPayment.DestinationId = RefundWallet.ID
	RefundPayment.Type = "Refund"
	RefundPayment.Status = "processing"
	InputPaymentsChannel <- &RefundPayment
	return RefundPayment.ID, nil
}

//payments.InitiateRefund(Complaint.PaymentsId,Complaint.UserId)

func InitiateCashback(payment *Payments) (err error) {
	var userID string
	if strings.HasPrefix(payment.SourceId, constants.WalletPrefix) {
		var Source wallet.Detail
		err := wallet.GetWalletById(&Source, payment.SourceId)
		if err != nil {
			return err
		}
		userID = Source.UserId
	} else {
		var Source bank.Detail
		err := bank.GetBankById(&Source, payment.SourceId)
		if err != nil {
			return err
		}
		userID = Source.UserId
	}
	Cashback := campaigns.Eligibility(payment.CreatedAt, payment.Amount, userID)
	if Cashback > 0 {
		fmt.Println("initiating cashback :",Cashback)
		var CashbackPayment Payments
		CashbackPayment.ID = GenerateID()
		CashbackPayment.Amount = int64(Cashback)
		CashbackPayment.Currency = "INR"
		CashbackPayment.SourceId = "wal_Mh5gqYDWlNBYWq"
		CashbackPayment.DestinationId = payment.SourceId
		CashbackPayment.Type = "Cashback"
		CashbackPayment.Status = "processing"
		InputPaymentsChannel <- &CashbackPayment
		return AddPaymentToDB(&CashbackPayment)
	}
	return nil
}

func UpdateTransactionCount(payment *Payments) (err error) {
	var userID string
	if strings.HasPrefix(payment.SourceId, constants.WalletPrefix) {
		var Source wallet.Detail
		err := wallet.GetWalletById(&Source, payment.SourceId)
		if err != nil {
			return err
		}
		userID = Source.UserId
	} else {
		var Source bank.Detail
		err := bank.GetBankById(&Source, payment.SourceId)
		if err != nil {
			return err
		}
		userID = Source.UserId
	}
	var x user.Detail
	user.GetUserById(&x,userID)
	x.NumberOfTransactions=x.NumberOfTransactions+1
	config.DB.Table("user").Save(&x)
	return nil
}
