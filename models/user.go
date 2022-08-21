package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int       `json:"id" gorm:"id"`
	Name      string    `json:"name" gorm:"name"`
	Username  string    `json:"user_name" gorm:"user_name"`
	Password  string    `json:"password" gorm:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	CreatedBy int       `json:"created_by" gorm:"created_by"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	UpdatedBy int       `json:"updated_by" gorm:"updated_by"`
}

func GetUser(param interface{}, DB *gorm.DB) (User, error) {

	var user User

	GetUser := DB.Raw("select id, name , user_name , password , created_at, created_by , updated_at , updated_by from users where user_name = ? ", param).First(&user)
	if GetUser.Error != nil {
		return user, GetUser.Error
	}
	return user, nil
}
