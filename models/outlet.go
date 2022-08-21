package models

import (
	"time"

	"github.com/ersa97/test-majoo/paginations"
	"github.com/jinzhu/gorm"
)

type MerchantOutlet struct {
	MerchantName string    `json:"merchant_name" gorm:"merchant_name"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at"`
	OutletName   string    `json:"outlet_name" gorm:"outlet_name"`
	BillTotal    float64   `json:"bill_total" gorm:"bill_total"`
}

type MerchantOutlets []MerchantOutlet

func (MerchantOutlet) TableName() string {
	return "merchantomzetoutlet"
}

func GetMerchantOutletOmzetByUserId(userid, limit, page int, DB *gorm.DB) (*paginations.Paginator, error) {

	var mo MerchantOutlets
	var GetMO, CountMO *gorm.DB

	CountMO = DB.Where("user_id = ?", userid)
	GetMO = DB.Raw("select merchant_name ,outlet_name ,created_at,bill_total from merchantomzetoutlet where user_id = ?", userid)
	if GetMO.Error != nil {
		return nil, GetMO.Error
	}

	paginator := paginations.Paging(&paginations.Param{
		DbSelect: GetMO,
		DbCount:  CountMO,
		Page:     page,
		Limit:    limit,
	}, &mo)

	return paginator, nil
}
