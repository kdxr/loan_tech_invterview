package models

import (
	"time"
)

type Payment struct {
	// gorm.Model
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`

	LoanId          uint    `gorm:"column:loan_id;not null" json:"loanId" validate:"required,numeric"`
	PayAmount       float64 `gorm:"column:pay_amount;not null" json:"payAmount" validate:"required,number"`
	PrincipleAmount float64 `gorm:"column:principle_amount;not null" json:"principleAmount" validate:"required,number"`
	InterestAmount  float64 `gorm:"column:interest_amount;not null" json:"interestAmount" validate:"required,number"`
}
