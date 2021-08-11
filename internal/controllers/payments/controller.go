package payments

import (
	"github.com/freshpay/internal/constants"
	"github.com/freshpay/internal/entities/payments/payments"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct{
	Entity string
	Payment payments.Payments
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
		c.JSON(http.StatusBadRequest,constants.Failure{Error: constants.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
	} else {
		resp:= Response{Entity: "payments",Payment: payment}
		c.JSON(http.StatusOK, resp)
	}
}

func GetPaymentByID(c *gin.Context) {
	var payment payments.Payments
	id := c.Params.ByName("payments_id")
	err := payments.GetPaymentByID(&payment, id)
	if err != nil {
		c.JSON(http.StatusBadRequest,constants.Failure{Error: constants.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
	} else {
		resp:= Response{Entity: "payments",Payment: payment}
		c.JSON(http.StatusOK, resp)
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
		c.JSON(http.StatusBadRequest,constants.Failure{Error: constants.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
	} else {
		resp := make([]Response, len(payment))
		for i,payment := range payment{
			resp[i]= Response{Entity: "Payments",Payment: payment}
		}
		c.JSON(http.StatusOK, resp)
	}
}
