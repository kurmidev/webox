package models

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm/clause"
)

type SubscriberBouque struct {
	ID               int `gorm:"primary_key"`
	BouqueId         int
	BouqueType       int
	AccountId        int
	CasId            int
	SubscriberId     int
	OperatorId       int
	ActivationDate   time.Time
	DeactivationDate time.Time
	Status           int
	Amount           float64
	Tax              float64
	Mrp              float64
	RefundAmount     float64
	RefundTax        float64
	RefundTds        float64
	RefundMrp        float64
	RefundDate       time.Time
	Remark           string
	CasStatus        int
	IsRefundable     int
	MrpTax           float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CreatedBy        int
	UpdatedBy        int
	Bouque           Bouquet
}

func (sb *SubscriberBouque) TableName() string {
	return "subscriber_bouque"
}

type Bouquet struct {
	ID               int `gorm:"primary_key"`
	Name             string
	Code             string
	IsHd             int
	Type             int
	Description      string
	MrpData          string
	Mrp              float64
	AdditionalRates  string
	IsExclusive      int
	SortBy           int
	IsPromotional    int
	IsOnlineApp      int
	ExpiryDate       time.Time
	LastSyncDetails  string
	FamilyIds        string
	Remark           string
	Status           int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CreatedBy        int
	UpdatedBy        int
	BouqueAssetAssoc []BouqueAssetAssoc `gorm:"foreignKey:ID"`
}

func (b *Bouquet) TableName() string {
	return "bouque"
}

type BouqueAssetAssoc struct {
	ID        int `gorm:"primary_key"`
	BouqueId  int
	PackageId int
	ChannelId int
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy int
	UpdatedBy int
	Package   Package `gorm:"foreignKey:PackageId"`
	Channel   Channel `gorm:"foreignKey:ChannelId"`
}

func (bac *BouqueAssetAssoc) TableName() string {
	return "bouque_asset_assoc"
}

func (b *Bouquet) GetBouques() ([]Bouquet, error) {
	var bouque []Bouquet
	sql := DB.Table(b.TableName()).
		Preload("BouqueAssetAssoc").
		Preload("BouqueAssetAssoc.Package").
		Preload("BouqueAssetAssoc.Channel.Broadcaster").
		Preload("BouqueAssetAssoc.Channel").
		Preload("BouqueAssetAssoc.Package.ChannelPackageAssoc").
		Preload(clause.Associations).Where("bouque.is_online_app=1")

	err := sql.Find(&bouque).Error
	fmt.Println("errr of the query", err)
	return bouque, err
}

func (b *Bouquet) GetBouque(id int) (Bouquet, error) {
	var bouque Bouquet
	sql := DB.Table(b.TableName()).
		Preload("BouqueAssetAssoc").
		Preload("BouqueAssetAssoc.Package").
		Preload("BouqueAssetAssoc.Channel.Broadcaster").
		Preload("BouqueAssetAssoc.Channel").
		Preload("BouqueAssetAssoc.Channel.Genre").
		Preload("BouqueAssetAssoc.Package.ChannelPackageAssoc").
		Preload(clause.Associations).Where("bouque.is_online_app=1").Where("id=?", id)

	err := sql.Find(&bouque).Error
	fmt.Println("errr of the query", err)
	return bouque, err
}

func (b *Bouquet) GetIsCutomNcf() bool {
	var meta = make(map[string]interface{})
	err := json.Unmarshal([]byte(b.MrpData), &meta)

	if err != nil {
		return false
	}

	if len(meta) > 0 {
		if _, ok := meta["ncf"]; ok {
			return true
		}
	}
	return false
}
