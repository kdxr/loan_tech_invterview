package models

import (
	"time"
)

type Loan struct {
	// gorm.Model
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`

	CustomerId   uint      `gorm:"column:customer_id;not null" json:"customerId" validate:"required,numeric"`
	LoanAmount   float64   `gorm:"column:loan_amount;not null" json:"loanAmount" validate:"required,number"`
	InterestRate float64   `gorm:"column:interest_rate;not null" json:"interestRate" validate:"required,number"`
	// PayDate      time.Time `gorm:"column:pay_date;not null" json:"payDate" validate:"required"`
	StartDate    time.Time `gorm:"column:start_date;not null" json:"startDate" validate:"required"`
	EndDate      time.Time `gorm:"column:end_date;not null" json:"endDate" validate:"required"`
}

type LoanInformation struct {
	Loan
	Customer Customer `gorm:"foreignKey:CustomerId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"customer"`
}

type UpdateLoan struct {
	ID           uint      `json:"id" validate:"required,numeric"`
	LoanAmount   float64   `json:"loanAmount" validate:"required,number"`
	InterestRate float64   `json:"interestRate" validate:"required,number"`
	// PayDate      time.Time `json:"payDate" validate:"required"`
	StartDate    time.Time `json:"startDate" validate:"required"`
	EndDate      time.Time `json:"endDate" validate:"required"`
}
