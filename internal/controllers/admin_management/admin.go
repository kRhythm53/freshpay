package admin_management/*
	CreateProduct will add a new product to the database
*/import (
	"github.com/freshpay/internal/entities/Error"
	"github.com/freshpay/internal/entities/admin"
	"github.com/freshpay/internal/entities/admin/admin_session"
	"github.com/gin-gonic/gin"
	"net/http"
)


func SignUp(c *gin.Context) {
	var Admin admin.Detail
	c.BindJSON(&Admin)
	err := admin.SignUp(&Admin)
	if err != nil {
		c.JSON(http.StatusBadRequest,Error.Detail{
			"BAD_REQUEST_ERROR","Failed",err.Error(),"buisness",
			"BAD REQUEST","NA","{}",
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"Entity":admin.EntityName,
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
		c.JSON(401,Error.Detail{
			"UnAuthorized","Failed",err.Error(),"buisness",
			"Wrong Login Details","NA","{}",
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

