package user_management

import (
	"github.com/freshpay/internal/entities/user_management/session"
	"github.com/freshpay/internal/entities/user_management/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	SignUp will signup the user
*/
func SignUp(c *gin.Context) {
	var User user.Detail
	c.BindJSON(&User)
	err := user.SignUp(&User)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest,err)
	} else {
		c.JSON(http.StatusOK, User)
	}
}


//LoginByPassword will login the user by Password
func LoginByPassword(c *gin.Context){
	var loginInfo user.Detail
	c.BindJSON(&loginInfo)
	var Session session.Detail
	err:=user.LoginByPassword(loginInfo.PhoneNumber,loginInfo.Password,&Session)
	if err!=nil{
		c.AbortWithError(http.StatusNotFound,err)
	} else{
		c.JSON(http.StatusOK,gin.H{
			"session_id":Session.ID,
		})
	}
}
