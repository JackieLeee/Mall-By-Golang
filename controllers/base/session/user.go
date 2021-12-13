package session

// User 用户信息存储格式
type User struct {
	Username  string `json:"user_name"`
	UserId    string `json:"user_id"`
	LastLogin int64  `json:"last_login"`
	IsAuth    bool   `json:"is_auth"`
	CSRFToken string `json:"csrf_token"`
}
