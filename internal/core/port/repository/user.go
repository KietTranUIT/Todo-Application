package repository

import (
	"errors"
	"user-service/internal/core/dto"
)

var DuplicateError = errors.New("Duplicate Verificate Email")

type UserRepository interface {
	InsertUser(dto.UserDTO) error
	GetUserWithEmail(string) (*dto.UserDTO, error)
	GetUserWithUsername(string) (*dto.UserDTO, error)
	InsertVerifyData(dto.VerifyEmailDTO) error
	DeleteVerifyData(email string) error
	GetVerifyEmailData(email string) (*dto.VerifyEmailDTO, error)
}
