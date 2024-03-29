package request

type RequestSignUp struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type RequestSignin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RequestSendVerificationEmail struct {
	Email string `json:"email"`
	Type  string `json:"type"`
}

type RequestCreateCategory struct {
	Owner       string `json:"owner"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
