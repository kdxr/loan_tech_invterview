package models

import "time"

type Payment struct {
	// gorm.Model
	// ID              uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt       *time.Time `json:"updatedAt"`
	PayAmount       float64 `gorm:"column:pay_amount;not null" json:"payAmount" `
	PrincipleAmount float64 `gorm:"column:principle_amount;not null" json:"principleAmount" `
	InterestAmount  float64 `gorm:"column:interest_amount;not null" json:"interestAmount"`
}
