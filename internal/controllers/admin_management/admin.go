package admin_management/*
	CreateProduct will add a new product to the database
*/import (
	"github.com/freshpay/internal/entities/admin"
	"github.com/freshpay/internal/entities/admin/admin_session"
	"github.com/freshpay/internal/entities/user_management/user"
	"github.com/gin-gonic/gin"
	"net/http"
)


func SignUp(c *gin.Context) {
	var Admin admin.Detail
	c.BindJSON(&Admin)
	err := admin.SignUp(&Admin)
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
			"ID":Admin.ID,
			"Name":Admin.Name,
			"PhoneNumber":Admin.PhoneNumber,
			"Email":Admin.Email,
		})
	}
}


//LoginByPassword will login the admin by Password
func LoginByPassword(c *gin.Context){
	var loginInfo admin.Detail
	c.BindJSON(&loginInfo)
	var Session admin_session.Detail
	var Admin admin.Detail
	err:=admin.LoginByPassword(loginInfo.PhoneNumber,loginInfo.Password,&Session,&Admin)
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
		c.JSON(200,gin.H{
			"Entity": admin.EntityName,
			"Status":"Success",
			"Message":"Login Successfully",
			"User": gin.H{
				"ID":Admin.ID,
				"Name":Admin.Name,
				"PhoneNumber":Admin.PhoneNumber,
				"Email":Admin.Email,
			},
		})
	}
}

