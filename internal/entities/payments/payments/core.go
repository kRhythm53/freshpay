package payments

import (
	"github.com/freshpay/internal/config"
)

func AddPayments(payments *Payments) (err error) {
	if err = config.DB.Table("payments").Create(payments).Error; err != nil {
		return err
	}
	return nil
}
