package complaints

import (
	"fmt"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/complaints"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetComplaints(c *gin.Context) {
	var Complaint []complaints.Complaint
	err := complaints.GetAllComplaints(&Complaint)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Complaint)
	}
}

func CreateComplaint(c *gin.Context) {
	var Complaint complaints.Complaint
	c.BindJSON(&Complaint)
	err := complaints.CreateComplaint(&Complaint)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Complaint)
	}
}

func GetComplaintByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var Complaint complaints.Complaint
	err := complaints.GetComplaintByID(&Complaint, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Complaint)
	}
}


func UpdateComplaint(c *gin.Context) {
	var Complaint complaints.Complaint
	id := c.Params.ByName("id")
	fmt.Println(id)
	err := complaints.GetComplaintByID(&Complaint, id)
	if err != nil {
		c.JSON(http.StatusNotFound, Complaint)
	}
	c.BindJSON(&Complaint)
	config.DB.Where("id = ?", id).Delete(&Complaint)
	err = complaints.UpdateComplaint(&Complaint)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Complaint)
	}
}

