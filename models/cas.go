package models

import "time"

type CasVendor struct {
	ID          int `gorm:"primary_key"`
	Name        string
	Code        string
	Setting     string
	Description string
	CasType     string
	IsCas       string
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   int
	UpdatedBy   int
}

func (cas *CasVendor) TableName() string {
	return "cas_vendor"
}
