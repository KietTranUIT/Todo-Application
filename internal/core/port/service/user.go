package service

import (
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"
)

type UserService interface {
	SignUp(req request.RequestSignUp) *response.Response
	SendVerificationEmail(req request.RequestSendVerificationEmail) *response.Response
	VerifyEmail(string, string) (bool, error)
}
