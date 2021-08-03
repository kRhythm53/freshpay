package Beneficiary

import "github.com/kshitij-nawandar9/freshpay/internal/Config"

//CreateBeneficiary will create a new Beneficiary
func CreateBeneficiary(beneficiary *Detail)(err error){
	if err=Config.DB.Create(beneficiary).Error; err!=nil{
		return err
	}
	return nil
}

//GetBeneficiaryById will return the Beneficiary by using the id
func GetBeneficiaryByID(beneficiary *Detail, entity_id string)(err error){
	if err = Config.DB.Where("entity_id = ?", entity_id).First(beneficiary).Error; err != nil {
		return err
	}
	return nil
}