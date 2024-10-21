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

type UpdateCustomer struct {
	ID             uint   `json:"id" validate:"required,numeric"`
	Email          string `json:"email" validate:"required,email,lte=255"`
	TelphoneNumber string `json:"telphoneNumber" validate:"required,numeric,gte=10"`
	Name           string `json:"name" validate:"required,lte=255"`
	Sex            string `json:"sex" validate:"required,lte=10"`
	Address        string `json:"address" validate:"required,lte=255"`
	Tumbon         string `json:"tumbon" validate:"required,lte=255"`
	District       string `json:"district" validate:"required,lte=255"`
	Province       string `json:"province" validate:"required,lte=255"`
	StatusActive   *bool  `json:"statusActive" validate:"required"`
}

// type Customer struct {
// 	gorm.Model
// 	// ID             int    `gorm:"column:customer_id;primaryKey;autoIncrement;not null"`
// 	Email          string `gorm:"not null;unique"`
// 	TelphoneNumber string `gorm:"column:telephonenumber;not null;unique"`
// 	Name           string
// 	Sex            string
// 	Address        string
// 	Tumbon         string
// 	District       string
// 	Province       string
// 	StatusActive   bool `gorm:"default:true"`
// }
