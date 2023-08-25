package service

import (
	"user-service/internal/core/dto"
)

type MailService interface {
	CreateMail(string, string) dto.MailData
	SendMail(dto.MailData) error
	
}
