package repository

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
	"user-service/internal/core/dto"
	"user-service/internal/core/port/repository"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepo struct {
	db repository.Database
}

var (
	InsertUserStatement          = "INSERT INTO USER_ACCOUNT VALUES ('%s', '%s', '%s', '%s', '%s')"
	GetUserWithEmailStatement    = "SELECT * FROM USER_ACCOUNT WHERE Email='%s'"
	GetUserWithUsernameStatement = "SELECT * FROM USER_ACCOUNT WHERE Username='%s'"

	InsertVerificationEmailStatement = "INSERT INTO VERIFICATION_EMAIL VALUES ('%s', '%s', '%s', '%s', '%d')"
	GetVerificationEmailStatement    = "SELECT * FROM VERIFICATION_EMAIL WHERE Email='%s'"
	DeleteVerifyDataStatement        = "DELETE FROM VERIFICATION_EMAIL WHERE Email='%s'"
)

const (
	duplicateEntry = "Duplicate entry"
)

var (
	duplicateError     = errors.New("Duplicate Verificate Email")
	internalError      = errors.New("Internal Error")
	rowAffectedError   = errors.New("Row Affected Error")
	notExistVerifyData = errors.New("Not Verify Data")
	notExistUserData   = errors.New("Not exist User")
)

func NewUserRepo(db repository.Database) repository.UserRepository {
	return UserRepo{
		db: db,
	}
}

func (u UserRepo) InsertUser(user dto.UserDTO) error {
	result, err := u.db.GetDB().Exec(fmt.Sprintf(InsertUserStatement,
		user.Email,
		user.Username,
		user.Password,
		user.CreateDate.Format("2006-01-02 15:04:05"),
		user.UpdateDate.Format("2006-01-02 15:04:05"),
	))

	if err != nil {
		if strings.Contains(err.Error(), duplicateEntry) {
			return duplicateError
		}
		return internalError
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return internalError
	}

	if rowAffected != 1 {
		return rowAffectedError
	}

	return nil
}

func (u UserRepo) InsertVerifyData(data dto.VerifyEmailDTO) error {
	result, err := u.db.GetDB().Exec(fmt.Sprintf(InsertVerificationEmailStatement,
		data.Email,
		data.Code,
		data.Type,
		data.ExpireAt.Format("2006-01-02 15:04:05"),
		data.IsVerified,
	))

	if err != nil {

		if strings.Contains(err.Error(), duplicateEntry) {
			verifyData, err := u.GetVerifyEmailData(data.Email)

			if err != nil {
				log.Println(err.Error())
				return err
			}

			if verifyData.IsVerified == 0 && verifyData.ExpireAt.Before(time.Now()) {
				err := u.DeleteVerifyData(verifyData.Email)

				if err != nil {
					return err
				}
				return u.InsertVerifyData(data)
			}
			fmt.Println("ok")
			return duplicateError
		}
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if rowAffected != 1 {
		return rowAffectedError
	}
	return nil
}

func (u UserRepo) GetVerifyEmailData(email string) (*dto.VerifyEmailDTO, error) {
	row := u.db.GetDB().QueryRow(fmt.Sprintf(GetVerificationEmailStatement,
		email,
	))

	if row == nil {
		return nil, notExistVerifyData
	}

	var verifyData dto.VerifyEmailDTO
	var expire string
	err := row.Scan(
		&verifyData.Email,
		&verifyData.Code,
		&verifyData.Type,
		&expire,
		&verifyData.IsVerified,
	)

	verifyData.ExpireAt, _ = time.ParseInLocation("2006-01-02 15:04:05", expire, time.Now().Location())

	if err != nil {
		return nil, err
	}

	return &verifyData, nil
}

func (u UserRepo) GetUserWithUsername(username string) (*dto.UserDTO, error) {
	row := u.db.GetDB().QueryRow(fmt.Sprintf(GetUserWithUsernameStatement, username))

	var user dto.UserDTO
	var createDate, updateDate string

	err := row.Scan(
		&user.Email,
		&user.Username,
		&user.Password,
		&createDate,
		&updateDate,
	)
	if err != nil {
		return nil, notExistUserData
	}

	user.CreateDate, _ = time.Parse("2006-01-02 15:04:05", createDate)
	user.UpdateDate, _ = time.Parse("2006-01-02 15:04:05", updateDate)

	return &user, nil
}

func (u UserRepo) GetUserWithEmail(email string) (*dto.UserDTO, error) {
	row := u.db.GetDB().QueryRow(fmt.Sprintf(GetUserWithEmailStatement, email))

	if row == nil {
		return nil, notExistUserData
	}

	var user dto.UserDTO
	var createDate, updateDate string

	err := row.Scan(
		&user.Email,
		&user.Username,
		&user.Password,
		&createDate,
		&updateDate,
	)

	if err != nil {
		return nil, err
	}

	user.CreateDate, _ = time.Parse("2006-01-02 15:04:05", createDate)
	user.UpdateDate, _ = time.Parse("2006-01-02 15:04:05", updateDate)

	return &user, nil
}

func (u UserRepo) DeleteVerifyData(email string) error {
	result, err := u.db.GetDB().Exec(fmt.Sprintf(DeleteVerifyDataStatement, email))

	if err != nil {
		log.Println(err.Error())
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if rowAffected != 1 {
		log.Println(err.Error())
		return err
	}
	return nil
}
