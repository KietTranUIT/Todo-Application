package dto

import (
	"time"
)

type UserDTO struct {
	Email      string
	Username   string
	Password   string
	CreateDate time.Time
	UpdateDate time.Time
}

type VerifyEmailDTO struct {
	Email      string
	Code       string
	Type       string
	ExpireAt   time.Time
	IsVerified uint
}

type MailData struct {
	From    string
	To      []string
	Subject string
	Code    string
}
