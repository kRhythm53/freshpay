package payments

import (
	"github.com/freshpay/internal/entities/payments/payments"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct{
	Entity string
	Payment payments.Payments
}
type Error struct{
	Code string
	Description string
	Source string
	Reason string
	Step string
	Metadata string
}
type Failure struct{
	Error Error
}

func AddPayments(c *gin.Context) {
	var payment payments.Payments
	err := c.BindJSON(&payment)
	if err != nil {
		return
	}
	userId :=c.GetString("userId")
	err = payments.AddPayments(&payment,userId)
	if err != nil {
		//c.String(http.StatusBadRequest, err2.Error())
		c.JSON(http.StatusBadRequest,Failure{Error: Error{"Bad request error",err.Error(),"business","validation failed","NA",""}})
	} else {
		//fmt.Println(resp)
		//fmt.Println(payment)
		resp:= Response{Entity: "payments",Payment: payment}
		c.JSON(http.StatusOK, resp)
		//c.JSON(http.StatusOK,payment)
	}
}

func GetPaymentByID(c *gin.Context) {
	var payment payments.Payments
	id := c.Params.ByName("payments_id")
	err := payments.GetPaymentByID(&payment, id)
	if err != nil {
		//c.String(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest,Failure{Error: Error{"Bad request error",err.Error(),"business","validation failed","NA",""}})
	} else {
		resp:= Response{Entity: "payments",Payment: payment}
		c.JSON(http.StatusOK, resp)
		//c.JSON(http.StatusOK, payment)
	}
}

func GetPaymentsByTime(c *gin.Context) {
	var payment []payments.Payments
	from := c.Query("from")
	to := c.Query("to")
	TransactionType := c.Query("type")
	userId := c.GetString("userId")
	err := payments.GetPaymentsByTime(&payment, from, to, TransactionType, userId)
	if err != nil {
		//c.String(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest,Failure{Error: Error{"Bad request error",err.Error(),"business","validation failed","NA",""}})
	} else {
		resp := make([]Response, len(payment))
		for i,payment := range payment{
			resp[i]= Response{Entity: "Payments",Payment: payment}
		}
		c.JSON(http.StatusOK, resp)
		//c.JSON(http.StatusOK, payment)
	}
}
