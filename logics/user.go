package logics

import (
	"time"

	"github.com/beego/beego/v2/core/logs"

	"mall/libs/tools"
	"mall/models"
)

func CreateNewUser(username, password string) bool {
	salt := tools.GenerateSalt()
	encryptPwd := tools.EncryptPasswordWithSalt(password, salt)

	newUser := models.User{
		Username:   username,
		Password:   encryptPwd,
		Salt:       salt,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Status:     models.UserEnable,
		Balance:    0,
	}

	if _, err := models.InsertOneUser(&newUser); err != nil {
		logs.Warn("insert user failed, err: %s", err.Error())
		return false
	}

	logs.Info("create user, id[%s] username[%s]", newUser.Id, newUser.Username)

	return true
}

func IsUserExist(username string) (bool, error) {
	user, err := models.FindUserByName(username)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
