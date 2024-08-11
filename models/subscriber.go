package models

import (
	"fmt"
	"time"

	"gorm.io/gorm/clause"
)

type Subscriber struct {
	ID                  int `gorm:"primary_key"`
	CustomerId          string
	Formno              string
	LocationId          int
	SublocationId       int
	OperatorId          int
	Fname               string
	Mname               string
	Lname               string
	Gender              int
	BillingAddress      string
	CustomerType        int
	Pincode             string
	Flat                string
	Floor               string
	Wing                string
	InstallationAddress string
	Dob                 time.Time
	Email               string
	MobileNo            string
	PhoneNo             string
	UId                 string
	Minid               string
	Remark              string
	LcoForm             string
	UploadFlag          string
	MetaData            string
	ExtraEmails         string
	IsVerified          int
	Status              int
	CreatedAt           time.Time
	UpdatedAt           time.Time
	CreatedBy           int
	UpdatedBy           int
	SubscriberAccount   []Account   `gorm:"foreignKey:SubscriberId"`
	Location            Location    `gorm:"foreignKey:LocationId"`
	Sublocation         Sublocation `gorm:"foreignKey:SublocationId;"`
	Operator            Operator    `gorm:"foreignKey:OperatorId"`
}

func (s *Subscriber) TableName() string {
	return "subscriber"
}

func (s *Subscriber) GetSubscriber(id int) (Subscriber, error) {
	var subscriber Subscriber
	sql := DB.Table(s.TableName()).
		Preload("SubscriberAccount").Preload("SubscriberAccount.SubscriberBouque").Preload("SubscriberAccount.SubscriberBouque.Bouque").
		Preload("Operator.Mso").Preload("Operator.Branch").Preload("Operator.Distributor").Preload("subscriber.Operator.District").
		Preload("Location").Preload("Sublocation").Preload("Operator").Preload("Operator.City").
		Preload(clause.Associations).Where("subscriber.id=?", id)

	err := sql.Find(&subscriber).Error
	fmt.Println("errr of the query", err)
	return subscriber, err
}
