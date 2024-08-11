package models

import "time"

type Brand struct {
	ID        int `gorm:"primary_key"`
	Name      string
	Code      string
	Type      string
	CasId     string
	Length    string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy int
	UpdatedBy int
}

func (cas *Brand) TableName() string {
	return "brand"
}
