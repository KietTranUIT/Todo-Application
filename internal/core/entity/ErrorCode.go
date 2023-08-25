package entity

type ErrorCode string

const (
	InvalidErrorCode        ErrorCode = "INVALID"
	DuplicateErrorCode      ErrorCode = "DUPLICATE"
	InternalErrorCode       ErrorCode = "INTERNAL ERROR"
	SuccessInsertCode       ErrorCode = "SUCCESS INSERT CODE"
	SendMailErrorCode       ErrorCode = "SEND FAILED"
	VerifyFailErrorCode     ErrorCode = "VERIFY FAILED"
	VerifySuccess           ErrorCode = "VERIFY SUCCESS"
	NotExistVerifyDataError ErrorCode = "NOT EXIST VERIFY DATA"
	SuccessInsertUser       ErrorCode = "SUCCESS INSERT USER"
	WrongCode               ErrorCode = "WRONG CODE"
	ExpiredCode             ErrorCode = "EXPIRED CODE"
)

const (
	InvalidMessage        string = "invalid email address"
	DuplicateUserMsg      string = "duplicate user"
	DuplicateEmailMsg     string = "duplicate email user"
	DuplicateCodeMsg      string = "duplicate code"
	InternalErrorMsg      string = "internal error"
	SuccessInsertCodeMsg  string = "success insert code"
	SendMailErrorMsg      string = "send mail failed"
	VerifyFailErrorMsg    string = "verify failed"
	VerifySuccessMsg      string = "verify success"
	NotExistVerifyDataMsg string = "not exist verify data"
	SuccessInsertUserMsg  string = "success insert user"
	WrongCodeMsg          string = "wrong code"
	ExpiredCodeMsg        string = "code expired"
)
