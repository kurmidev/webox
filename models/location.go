package models

import "time"

type Sublocation struct {
	ID         int `gorm:"primary_key"`
	Name       string
	Code       string
	LocationId int
	OperatorId int
	Status     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CreatedBy  int
	UpdatedBy  int
}

func (sl *Sublocation) TableName() string {
	return "sublocation"
}

type Location struct {
	Id         int `gorm:"primary_key"`
	Name       string
	Code       string
	OperatorId int
	Status     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CreatedBy  int
}

func (loc *Location) TableName() string {
	return "location"
}
