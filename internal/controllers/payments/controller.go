package payments

import (
	"github.com/freshpay/internal/entities/payments/payments"
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct{
	payment payments.Payments
	entity string
}
func AddPayments(c *gin.Context) {
	var payment payments.Payments
	err := c.BindJSON(&payment)
	if err != nil {
		return
	}
	userId :=c.GetString("userId")
	err2 := payments.AddPayments(&payment,userId)
	if err2 != nil {
		c.String(http.StatusBadRequest, err2.Error())
	} else {
		c.JSON(http.StatusOK, response{payment: payment,entity: "payments"})
	}
}

func GetPaymentByID(c *gin.Context) {
	var payment payments.Payments
	id := c.Params.ByName("payments_id")
	err := payments.GetPaymentByID(&payment, id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, payment)
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
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, payment)
	}
}
