package request

type ParamUserRegister struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type ParamUserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
