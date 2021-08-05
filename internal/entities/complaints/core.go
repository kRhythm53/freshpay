package complaints

import (
	"github.com/freshpay/internal/config"
)


//GetAllComplaintss Fetch all user data
func GetAllComplaints(Complaints *[]Complaint) (err error) {
	if err = config.DB.Find(Complaints).Error; err != nil {
		return err
	}
	return nil
}

// CreateComplaints  ... Insert New data
func CreateComplaint(Complaint *Complaint) (err error) {
	Complaint.ID = "cmplt_"+RandomString(14)
	if err = config.DB.Create(Complaint).Error; err != nil {
		return err
	}
	return nil
}

//GetComplaintsByID ... Fetch only one user by Id
func GetComplaintByID(Complaint *Complaint, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(Complaint).Error; err != nil {
		return err
	}
	return nil
}

//UpdateComplaints ... Update user
func UpdateComplaint(Complaint *Complaint) (err error) {
	config.DB.Save(Complaint)
	return nil
}
