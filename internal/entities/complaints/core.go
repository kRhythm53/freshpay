package complaints

import (
	"fmt"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/payments/payments"
)


func GetAllComplaints(Complaints *[]Complaint) (err error) {
	return GetAllComplaintsFromDB(Complaints)
}
func GetAllActiveComplaints(Complaints *[]Complaint) (err error){
	return GetAllActiveComplaintsFromDB(Complaints)
}
func CreateComplaint(Complaint *Complaint,user_id string) (err error) {
	Complaint.ID = "cmplt_"+RandomString(14)
	return AddComplaintToDB(Complaint)
}

func GetComplaintByID(Complaint *Complaint, id string) (err error) {
	return GetComplaintByIDFromDB(Complaint,id)
}

func UpdateComplaint(Complaint *Complaint, id string, refund string) (err error) {
	config.DB.Where("id = ?", id).Delete(&Complaint)
	if Complaint.ComplaintType == "other"{
		Complaint.Status = "processed"
		fmt.Println("Complaint Resloved")
	}else if Complaint.ComplaintType == "refund"{
		if refund == "eligible"{
			Complaint.Status = "processed"
			payments.InitiateRefund(Complaint.PaymentsId)
		}else{
			Complaint.Status = "processed"
		}
	}
	config.DB.Save(Complaint)
	return nil
}
