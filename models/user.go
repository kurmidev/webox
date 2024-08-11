package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id                   int `gorm:"primary_key"`
	Name                 string
	Username             string
	AuthKey              string
	AccessTokenExpiredAt time.Time
	PasswordHash         string
	PasswordResetToken   sql.NullString
	Email                string
	ConfirmedAt          sql.NullTime
	RegistrationIp       sql.NullString
	LastLoginAt          time.Time
	LastLoginIp          sql.NullString
	BlockedAt            sql.NullTime
	Role                 int
	MobileNo             string
	OperatorId           int
	OperatorType         int
	Status               int
	DesignationId        int
	CreatedAt            time.Time
	UpdatedAt            time.Time
	CreatedBy            int
	UpdatedBy            int
	AccessPolicyId       int
	BlockMessage         sql.NullString
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) GetUser(username string) (User, error) {
	var userdetails User
	err := DB.Table(u.TableName()).Where("username=?", username).Scan(&userdetails).Error
	return userdetails, err
}

func (u *User) GetUserByMobile(mobileno string) (User, error) {
	var userdetails User
	err := DB.Table(u.TableName()).Where("mobileno=? and operator_type = 4", mobileno).Scan(&userdetails).Error
	return userdetails, err
}
