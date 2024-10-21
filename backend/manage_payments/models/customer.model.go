package models

import (
	"time"
)

type Customer struct {
	// gorm.Model
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	// DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Email          string `gorm:"not null;uniqueIndex" json:"email" validate:"required,email,lte=255"`
	TelphoneNumber string `gorm:"column:telephonenumber;not null;uniqueIndex" json:"telphoneNumber" validate:"required,numeric,gte=10"`
	Name           string `gorm:"column:name" json:"name" validate:"required,lte=255"`
	Sex            string `json:"sex" validate:"required,lte=10"`
	Address        string `json:"address" validate:"required,lte=255"`
	Tumbon         string `json:"tumbon" validate:"required,lte=255"`
	District       string `json:"district" validate:"required,lte=255"`
	Province       string `json:"province" validate:"required,lte=255"`
	StatusActive   bool   `gorm:"default:true" json:"statusActive"`
}
