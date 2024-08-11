package models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type Models struct {
	DB *gorm.DB
	Account
	Subscriber
	CasVendor
	Location
	Sublocation
	Brand
	Operator
	OperatorBalance
	User
	Bouquet
	BouqueAssetAssoc
	Channel
	Package
	Broadcaster
	Genre
	Language
}

func New(databasePool *gorm.DB) *Models {
	DB = databasePool
	return &Models{
		DB:               databasePool,
		Account:          Account{},
		Subscriber:       Subscriber{},
		CasVendor:        CasVendor{},
		Location:         Location{},
		Sublocation:      Sublocation{},
		Brand:            Brand{},
		Operator:         Operator{},
		OperatorBalance:  OperatorBalance{},
		User:             User{},
		Bouquet:          Bouquet{},
		BouqueAssetAssoc: BouqueAssetAssoc{},
		Channel:          Channel{},
		Package:          Package{},
		Broadcaster:      Broadcaster{},
		Genre:            Genre{},
		Language:         Language{},
	}

}
