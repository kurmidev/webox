package models

import "time"

type AdministrativeDivision struct {
	ID          int `gorm:"primary_key"`
	Name        string
	Code        string
	Status      int
	Description string
	StateId     int
	DistrictId  int
	Type        int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   int
	UpdatedBy   int
}

func (adv *AdministrativeDivision) TableName() string {
	return "administrative_division"
}
