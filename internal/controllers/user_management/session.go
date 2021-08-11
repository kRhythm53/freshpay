package user_management

import (
	"github.com/gin-gonic/gin"
	"github.com/freshpay/internal/entities/user_management/user_session"
	"net/http"
)

func GetSessionById(c* gin.Context){
	var Session user_session.Detail
	sessionId:=c.Params.ByName("entity_id")
	err:= user_session.GetSessionById(&Session,sessionId)
	if err!=nil{
		c.AbortWithStatus(http.StatusBadGateway)
	} else{
		c.JSON(http.StatusOK,Session)
	}
}


