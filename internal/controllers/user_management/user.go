package user_management

import (
	"fmt"
	"github.com/freshpay/internal/entities/user_management/user"
	"github.com/freshpay/internal/entities/user_management/user_session"
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
		c.JSON(500,gin.H{
			"Code": "BAD_REQUEST_ERROR",
			"Description":err.Error(),
			"Source": "business",
			"Reason": "input_validation_failed",
			"Step": "NA",
			"Metadata":"{}",
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"Entity":user.EntityName,
			"Status":"success",
			"ID":User.ID,
			"Name":User.Name,
			"PhoneNumber":User.PhoneNumber,
			"Email":User.Email,
		})
	}
}


//LoginByPassword will login the user by Password
func LoginByPassword(c *gin.Context){
	var loginInfo user.Detail
	c.BindJSON(&loginInfo)
	fmt.Println(loginInfo)
	var Session user_session.Detail
	var User user.Detail
	err:=user.LoginByPassword(loginInfo.PhoneNumber,loginInfo.Password,&Session,&User)
	if err!=nil{
		c.JSON(401,gin.H{
			"Code": "Unauthorized",
			"Description":err.Error(),
			"Source": "business",
			"Reason": "Wrong Login Details",
			"Step": "NA",
			"Metadata":"{}",
		})
	} else{
		c.Writer.Header().Set("session_id",Session.ID)
		c.JSON(http.StatusOK,gin.H{
			"Entity": user.EntityName,
			"status": gin.H{
				"type": "success",
				"message": "Success",
				"code": 200,
				"error": false,
			},
			"Data": gin.H{
				"status": "Authenticated",
				"User": gin.H{
					"ID":User.ID,
					"Name":User.Name,
					"PhoneNumber":User.PhoneNumber,
					"Email":User.Email,
				},
			},
		})
	}
}
