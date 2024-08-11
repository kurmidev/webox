package models

import (
	"fmt"
	"time"

	"gorm.io/gorm/clause"
)

type Account struct {
	ID                       int `gorm:"primary_key"`
	SubscriberId             int
	CustomerId               string
	LocationId               int
	SublocationId            int
	OperatorId               int
	PairingId                int
	Smartcardno              string
	Stbno                    string
	StbId                    int
	ScId                     int
	CasId                    int
	IsHd                     int
	IsEmbeded                int
	StbbrandId               int
	ActivationDate           time.Time
	DeactivationDate         time.Time
	Status                   int
	CreatedAt                time.Time
	UpdatedAt                time.Time
	CreatedBy                int
	UpdatedBy                int
	Location                 Location           `gorm:"foreignKey:LocationId"`
	Sublocation              Sublocation        `gorm:"foreignKey:SublocationId;"`
	Operator                 Operator           `gorm:"foreignKey:OperatorId"`
	Subscriber               Subscriber         `gorm:"foreignKey:SubscriberId"`
	Cas                      CasVendor          `gorm:"foreignKey:ID;references:CasId"`
	Brand                    Brand              `gorm:"foreignKey:ID;references:StbbrandId"`
	SubscriberBouque         []SubscriberBouque `gorm:"foreignKey:AccountId"`
	InActiveSubscriberBouque []SubscriberBouque `gorm:"foreignKey:AccountId"`
}

func (ac *Account) TableName() string {
	return "subscriber_account"
}

func (ac *Account) GetAccount(number string, types int) (Account, error) {
	var account Account
	sql := DB.Table(ac.TableName()).
		Preload("Subscriber").Preload("Location").Preload("Sublocation").Preload("Operator").Preload("Brand").
		Preload("Operator.City").
		Preload("SubscriberBouque").
		Preload("SubscriberBouque.Bouque").
		Preload("InActiveSubscriberBouque", " status!=? ", 1).
		Preload(clause.Associations)
	if types == 1 {
		sql.Where("subscriber_account.smartcardno=?", number)
	} else {
		sql.Where("subscriber_account.stbno=?", number)
	}
	err := sql.Find(&account).Error
	fmt.Println("errr of the query", err)
	return account, err
}
