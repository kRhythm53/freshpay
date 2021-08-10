package complaints

import (
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
	Complaint.UserId = user_id
	Complaint.Status = "pending"
	return AddComplaintToDB(Complaint)
}

func GetComplaintByID(Complaint *Complaint, id string) (err error) {
	return GetComplaintByIDFromDB(Complaint,id)
}

func UpdateComplaint(Complaint *Complaint, id string, refund string) (err error) {
	//config.DB.Where("id = ?", id).Delete(&Complaint)
	if Complaint.ComplaintType == "other"{
		Complaint.Status = "processed"
	}else if Complaint.ComplaintType == "refund"{
		if refund == "eligible"{
			Complaint.Status = "processed"
			var RefundId string
			RefundId,err = payments.InitiateRefund(Complaint.PaymentsId,Complaint.UserId)
			if err!=nil {
				return err
			}
			Complaint.RefundId=RefundId
		}else{
			Complaint.Status = "processed"
		}
	}
	config.DB.Save(Complaint)
	return nil
}