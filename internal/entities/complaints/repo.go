package complaints

import (
	"fmt"
	"github.com/freshpay/internal/config"
)

func AddComplaintToDB(Complaint *Complaint) (err error){
	if err := config.DB.Create(Complaint).Error; err != nil {
		fmt.Println("found error")
		return err
	}
	return nil
}

func GetComplaintByIDFromDB(Complaint *Complaint,id string) (err error){
	if err = config.DB.Where("id = ?", id).First(Complaint).Error; err != nil {
		return err
	}
	return nil
}

func GetAllComplaintsFromDB(Complaints *[]Complaint)(err error){
	if err = config.DB.Find(Complaints).Error; err != nil {
		return err
	}
	return nil
}

func GetAllActiveComplaintsFromDB(Complaints *[]Complaint)(err error) {
	if err = config.DB.Where("status = ?" , "pending").Find(Complaints).Error; err!=nil{
		return err
	}
	return nil
}