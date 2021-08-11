package complaints

import (
	"fmt"
	"github.com/freshpay/internal/entities/complaints"
	"github.com/gin-gonic/gin"
	"net/http"
)



func CreateComplaint(c *gin.Context) {
	var Complaint complaints.Complaint
	c.BindJSON(&Complaint)
	userId :=c.GetString("userId")
	println(userId)
	err := complaints.CreateComplaint(&Complaint,userId)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Complaint)
	}
}


func UpdateComplaintById(c *gin.Context) {
	var Complaint complaints.Complaint
	id := c.Params.ByName("complaint_id")
	adminId := c.GetString("adminId")
	refund := c.GetString("refund")
	err := complaints.GetComplaintByID(&Complaint, id)
	if err != nil {
		c.JSON(http.StatusNotFound, Complaint)
	} else{
		c.BindJSON(&Complaint)
		err = complaints.UpdateComplaint(&Complaint, id, refund,adminId)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, Complaint)
		}
	}
}


func GetComplaints(c *gin.Context){
	var Complaint []complaints.Complaint
	err := complaints.GetAllComplaints(&Complaint)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Complaint)
	}
}

func GetActiveComplaints(c *gin.Context){
	var Complaint []complaints.Complaint
	err := complaints.GetAllActiveComplaints(&Complaint)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Complaint)
	}
}

func GetComplaintById(c *gin.Context) {
	id := c.Params.ByName("complaint_id")
	var Complaint complaints.Complaint
	err := complaints.GetComplaintByID(&Complaint, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Complaint)
	}
}