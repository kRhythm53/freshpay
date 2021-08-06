package user_management

import (
	"errors"
	"github.com/freshpay/internal/entities/user_management/session"
	"github.com/freshpay/internal/entities/user_management/wallet"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetWalletById(c *gin.Context){
	/*
		Part of Middle Ware start
		check if session_id exist and not expire and get userId
	*/

	sessionId:= c.Request.Header["Session_id"][0]
	var Session session.Detail
	err1:=session.GetSessionById(&Session, sessionId)
	if err1!=nil{
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if Session.ExpireTime < uint64(time.Now().Unix()){
		c.AbortWithError(400,errors.New("Session has expired"))
		return
	}
	userId:=Session.UserId
	/*
	 Part of middleWare ends
	*/

	var Wallet wallet.Detail
	err:=wallet.GetWalletBalanceByUserId(&Wallet,userId)
	if err!=nil{
		c.AbortWithError(400,err)
	} else{
		c.JSON(200,Wallet)
	}
}
