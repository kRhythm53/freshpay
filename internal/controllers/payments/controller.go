package payments

import (
	"github.com/freshpay/internal/entities/payments/payments"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddPayments(c *gin.Context) {
	var payment payments.Payments
	err := c.BindJSON(&payment)
	if err != nil {
		return
	}

	err2 := payments.AddPayments(&payment)
	if err2 != nil {
		c.String(http.StatusNotFound, "Payment failed")
	} else {
		c.JSON(http.StatusOK, payment)
	}
}

func GetPaymentByID(c *gin.Context) {
	var payment payments.Payments
	id := c.Params.ByName("payments_id")
	err := payments.GetPaymentByID(&payment, id)
	if err != nil {
		c.String(http.StatusNotFound, "Record not found")
	} else {
		c.JSON(http.StatusOK, payment)
	}
}

func GetPaymentsByTime(c *gin.Context) {
	var payment []payments.Payments
	from := c.Query("from")
	to := c.Query("to")

	err2 := payments.GetPaymentsByTime(&payment, from, to)
	if err2 != nil {
		c.String(http.StatusNotFound, "Request failed.")
	} else {
		c.JSON(http.StatusOK, payment)
	}
}
