package accounts

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name        string         `json:"name,omitempty"`
	Email       string         `json:"email,omitempty"`
	Password    string         `json:"password,omitempty"`

	Account     Account        `json:"profile,omitempty"`
	AccountId   uint
}

type Account struct {
	gorm.Model
}


type SignupCode struct {
	gorm.Model

	Code string
	MaxUses int
	Expiry time.Time
	Email string
	Notes string
	Sent time.Time
	Created time.Time
	UseCount int

	Inviter User `json:"inviter,omitempty" gorm:"ForeignKey:UserID"`
	UserID  uint
}


type SignupCodeResult struct {
	gorm.Model

	Code SignupCode
	SignupCodeID uint

	User User
	UserID uint

	Timestamp time.Time
}


