package payments

import (
	"fmt"
	"github.com/freshpay/internal/entities/payments/payments"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func AddPayments(c *gin.Context) {
	var payment payments.Payments
	err := c.BindJSON(&payment)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err2 := payments.AddPayments(&payment)
	if err2 != nil {
		c.String(http.StatusOK,"Payment failed")
	} else {
		c.JSON(http.StatusOK, payment)
	}
}

func GetPaymentByID(c *gin.Context) {
	var payment payments.Payments
	id := c.Params.ByName("payments_id")
	err := payments.GetPaymentByID(&payment, id)
	if err != nil {
		c.String(http.StatusOK,"Record not found")
	} else {
		c.JSON(http.StatusOK, payment)
	}
}

func GetPaymentsByTime(c *gin.Context) {
	var payment []payments.Payments
	from := c.Query("from")
	to := c.Query("to")
	var startTime,endTime int64
	var err error
	if from==""{
		startTime=time.Now().Unix()
	}else{
		startTime,err=strconv.ParseInt(from,10,64)
		if err!=nil{
			return
		}
	}
	if to==""{
		endTime=time.Now().Unix()
	}else{
		endTime,err=strconv.ParseInt(to,10,64)
	}
	fmt.Println(startTime,endTime)
	err2 := payments.GetPaymentsByTime(&payment, startTime, endTime)
	if err2 != nil {
		c.String(http.StatusOK,"Request failed.")
	} else {
		c.JSON(http.StatusOK, payment)
	}
}
