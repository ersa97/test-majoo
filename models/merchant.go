package models

import (
	"time"

	"github.com/ersa97/test-majoo/paginations"
	"github.com/jinzhu/gorm"
)

type Merchant struct {
	MerchantName string    `json:"merchant_name" gorm:"merchant_name"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at"`
	BillTotal    float64   `json:"bill_total" gorm:"bill_total"`
}

type Merchants []Merchant

func (Merchant) TableName() string {
	return "merchantomzet"
}

func GetMerchantOmzetByUserId(userid, limit, page int, DB *gorm.DB) (*paginations.Paginator, error) {

	var mo Merchants
	var GetMO, CountMO *gorm.DB
	CountMO = DB.Where("user_id =  ?", userid)
	GetMO = DB.Raw("select merchant_name, created_at, bill_total from merchantomzet where user_id =  ? ", userid)
	if GetMO.Error != nil {
		return nil, GetMO.Error
	}
	paginator := paginations.Paging(&paginations.Param{
		DbSelect: GetMO,
		DbCount:  CountMO,
		Page:     page,
		Limit:    limit,
		OrderBy:  []string{"user_id asc"},
	}, &mo)

	return paginator, nil
}
