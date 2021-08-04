package payments

import (
	"fmt"
	"github.com/freshpay/internal/entities/payments/payments"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddPayment(c *gin.Context) {
	var payment payments.Payments
	err := c.BindJSON(&payment)
	if err != nil {
		return
	}
	err2 := payments.AddPayments(&payment)
	if err2 != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, payment)
	}
}