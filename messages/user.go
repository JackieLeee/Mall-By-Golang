package messages

type UserRegisterReq struct {
	Username string `json:"username" valid:"Required;Match(/^[a-zA-Z0-9_-]{4,16}$/)"`
	Password string `json:"password" valid:"Required;"`
	Captcha  string `json:"captcha" valid:"Required;Match(/^[a-zA-Z0-9]{4,16}$/)"`
}

type UserLoginReq struct {
	Username string `json:"username" valid:"Required;Match(/^[a-zA-Z0-9_-]{4,16}$/)"`
	Password string `json:"password" valid:"Required;"`
	Captcha  string `json:"captcha" valid:"Required;Match(/^[a-zA-Z0-9]{4,16}$/)"`
}
