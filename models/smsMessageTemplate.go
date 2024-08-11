package models

import "time"

type SmsMessageTemplate struct {
	ID          int `gorm:"primary_key"`
	CatId       int
	SubCatId    int
	Template    string
	Status      int
	ExtraConf   string
	PlaceHolder string
	TemplateId  string
	SenderId    string
	MetaData    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   int
	UpdatedBy   int
}

func (smt *SmsMessageTemplate) TableName() string {
	return "sms_message_template"
}

func (smt *SmsMessageTemplate) GetTemplate(catId int, subCatId int) (SmsMessageTemplate, error) {
	var template SmsMessageTemplate
	err := DB.Table(smt.TableName()).Where("cat_id=? and sub_cat_id=? and status=1", catId, subCatId).Scan(&template).Error
	return template, err
}
