package entities

import (
	"resto-backend-go-1/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null;unique" valid:"required~Your Username is required"`
	Email    string `gorm:"not null;unique" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Role     string `gorm:"not null" valid:"required~Your role is required"`
	Birthday string `gorm:"not null" valid:"required~Your Birthday is required,date~Invalid Date Format"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
