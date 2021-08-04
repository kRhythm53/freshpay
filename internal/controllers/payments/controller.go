package payments

import (
	"fmt"
	"github.com/gin-gonic/gin"
	payments2 "github.com/kshitij-nawandar9/freshpay/internal/entities/payments/payments"
	"net/http"
)

func AddPayment(c *gin.Context) {
	var payment payments2.Payments
	err := c.BindJSON(&payment)
	if err != nil {
		return
	}
	err2 := payments2.AddPayments(&payment)
	if err2 != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, payment)
	}
}