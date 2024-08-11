package models

import "time"

type Operator struct {
	ID            int `gorm:"primary_key"`
	Name          string
	Code          string
	ContactPerson string
	Email         string
	MobileNo      string
	PhoneNo       string
	Addr          string
	Type          int
	MsoId         int
	BranchId      int
	DistributorId int
	UserId        int
	CityId        int
	Status        int
	DistrictId    int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreatedBy     int
	UpdatedBy     int
	Balance       OperatorBalance         `gorm:"foreignKey:ID"`
	City          AdministrativeDivision  `gorm:"foreignKey:CityId"`
	Mso           *Operator               `gorm:"foreignKey:MsoId"`
	Branch        *Operator               `gorm:"foreignKey:BranchId"`
	Distributor   *Operator               `gorm:"foreignKey:DistributorId"`
	District      *AdministrativeDivision `gorm:"foreignKey:DistrictId"`
}

func (op *Operator) TableName() string {
	return "operator"
}
