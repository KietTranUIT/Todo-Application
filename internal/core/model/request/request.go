package request

type RequestSignUp struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type RequestSendVerificationEmail struct {
	Email string `json:"email"`
	Type  string `json:"type"`
}
