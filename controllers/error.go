package controllers

import (
	"mall/controllers/base"
	"mall/libs/i18n"
)

var (
	ErrUserExist      = &base.GeneralResponse{Code: 1006, Message: i18n.Tr("ErrUserExist")}
	ErrCaptchaInvalid = &base.GeneralResponse{Code: 1007, Message: i18n.Tr("ErrCaptchaInvalid")}
)
