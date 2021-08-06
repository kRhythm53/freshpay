package complaints

import (
	"github.com/freshpay/internal/config"
)


func GetAllComplaints(Complaints *[]Complaint) (err error) {
	if err = config.DB.Find(Complaints).Error; err != nil {
		return err
	}
	return nil
}

func CreateComplaint(Complaint *Complaint) (err error) {
	Complaint.ID = "cmplt_"+RandomString(14)
	if err = config.DB.Create(Complaint).Error; err != nil {
		return err
	}
	return nil
}

func GetComplaintByID(Complaint *Complaint, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(Complaint).Error; err != nil {
		return err
	}
	return nil
}

func UpdateComplaint(Complaint *Complaint) (err error) {
	config.DB.Save(Complaint)
	return nil
}
