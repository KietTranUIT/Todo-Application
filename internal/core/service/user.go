package service

import (
	"errors"
	"fmt"
	"time"
	"user-service/internal/common"
	"user-service/internal/core/dto"
	"user-service/internal/core/entity"
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"
	"user-service/internal/core/port/repository"
	"user-service/internal/core/port/service"
)

type UserService struct {
	service repository.UserRepository
	mail    service.MailService
}

var (
	notExistVerifyData = errors.New("Not Verify Data")
	CodeExpired        = errors.New("Code Expired")
)

func NewUserService(userRepo repository.UserRepository, mail service.MailService) service.UserService {
	return UserService{
		service: userRepo,
		mail:    mail,
	}
}

func (u UserService) Signin(req request.RequestSignin) *response.Response {
	user, err := u.service.GetUserWithUsername(req.Username)

	if err != nil {
		if err.Error() == "Not exist User" {
			return CreateFailResponse(entity.NotExistUserMsg, entity.NotExistUserError)
		}
		return CreateFailResponse(entity.InternalErrorMsg, entity.InternalErrorCode)
	}

	if user.Password != req.Password {
		return CreateFailResponse(entity.WrongPasswordMsg, entity.WrongPasswordError)

	}
	return CreateSuccessResponse(entity.SuccessSingInMsg, entity.SuccessSignInError)
}

func (u UserService) SignUp(req request.RequestSignUp) *response.Response {
	if user, _ := u.service.GetUserWithEmail(req.Email); user != nil {
		return CreateFailResponse(entity.DuplicateUserMsg, entity.DuplicateErrorCode)
	}

	if verify, err := u.VerifyEmail(req.Email, req.Code); verify == false {
		if err == nil {
			return CreateFailResponse(entity.WrongCodeMsg, entity.WrongCode)
		}

		if err.Error() == notExistVerifyData.Error() {
			return CreateFailResponse(entity.NotExistVerifyDataMsg, entity.NotExistVerifyDataError)
		} else if err.Error() == CodeExpired.Error() {
			return CreateFailResponse(entity.ExpiredCodeMsg, entity.ExpiredCode)
		}
		return CreateFailResponse(entity.VerifyFailErrorMsg, entity.VerifyFailErrorCode)
	}

	user := dto.UserDTO{
		Email:      req.Email,
		Username:   req.Username,
		Password:   req.Password,
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
	}

	err := u.service.InsertUser(user)

	if err != nil {
		fmt.Println("IK")
		return CreateFailResponse(entity.InternalErrorMsg, entity.InternalErrorCode)
	}

	if err := u.service.DeleteVerifyData(user.Email); err != nil {
		fmt.Println("ok")
		return CreateFailResponse(entity.InternalErrorMsg, entity.InternalErrorCode)
	}

	return CreateSuccessResponse(entity.SuccessInsertUserMsg, entity.SuccessInsertUser)
}

func (u UserService) VerifyEmail(email string, code string) (bool, error) {
	verifyData, err := u.service.GetVerifyEmailData(email)

	if verifyData == nil {
		return false, notExistVerifyData
	}

	if verifyData.ExpireAt.Before(time.Now()) {
		return false, CodeExpired
	}

	if verifyData.Code == code {
		return true, nil
	}
	return false, err
}

func (u UserService) SendVerificationEmail(req request.RequestSendVerificationEmail) *response.Response {
	data, _ := u.service.GetUserWithEmail(req.Email)

	if data != nil {
		return CreateFailResponse(entity.DuplicateUserMsg, entity.DuplicateErrorCode)
	}

	code := common.GenerateRandomCode()

	if u.mail.SendMail(u.mail.CreateMail(req.Email, code)) != nil {
		return CreateFailResponse(entity.SendMailErrorMsg, entity.SendMailErrorCode)
	}

	verifyData := dto.VerifyEmailDTO{
		Email:      req.Email,
		Code:       code,
		Type:       req.Type,
		ExpireAt:   common.GetExpireTime(),
		IsVerified: 0,
	}

	err := u.service.InsertVerifyData(verifyData)

	if err != nil {
		if err.Error() == repository.DuplicateError.Error() {
			return CreateFailResponse(entity.DuplicateCodeMsg, entity.DuplicateErrorCode)
		}

		return CreateFailResponse(entity.InternalErrorMsg, entity.InternalErrorCode)
	}

	return CreateSuccessResponse(entity.SuccessInsertCodeMsg, entity.SuccessInsertCode)
}

func CreateFailResponse(message string, err entity.ErrorCode) *response.Response {
	return &response.Response{
		Status:   false,
		Err_code: err,
		Err_msg:  message,
	}
}

func CreateSuccessResponse(message string, err entity.ErrorCode) *response.Response {
	return &response.Response{
		Status:   true,
		Err_code: err,
		Err_msg:  message,
	}
}
