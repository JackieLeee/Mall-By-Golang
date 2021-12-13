package base

import (
	"reflect"

	"github.com/beego/beego/v2/core/logs"

	"mall/controllers/base/session"
)

type UserController struct {
	BaseController

	UserCtx *session.User

	IsAuth      bool
	disableCsrf bool
}

func (c *UserController) Prepare() {
	c.BaseController.Prepare()

	if !c.checkAuth() {
		c.ServeResponse(ErrNeedLogin)
		return
	}

	if !c.checkCSRF() {
		logs.Warn("CSRF Token invalid")
		c.ServeResponse(ErrCSRFToken)
		return
	}

	c.IsAuth = true
}

func (c *UserController) DisableCSRF() {
	c.disableCsrf = true
}

func (c *UserController) checkCSRF() bool {
	if c.disableCsrf {
		return true
	}
	token := c.Ctx.Request.Header.Get(session.KSessKeyCSRFToken)
	if c.UserCtx.CSRFToken == token {
		return true
	}
	logs.Warn("expect csrf_token is %s, actually is %s", c.UserCtx.CSRFToken, token)
	return false
}

func (c *UserController) checkAuth() bool {
	userSessItf, ok := c.MyGetSession(session.KSessKeyUser)
	if !ok {
		logs.Warn("not find session by user")
		return false
	}

	userSess, ok := userSessItf.(session.User)
	if !ok {
		logs.Warn("user session convert string failed, type: %v", reflect.TypeOf(userSessItf))
		return false
	}

	c.UserCtx = &userSess
	return userSess.IsAuth
}
