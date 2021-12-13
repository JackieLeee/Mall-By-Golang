package base

import (
	"mall/libs/i18n"
)

var (
	ErrOK  = &GeneralResponse{Code: 0, Message: i18n.Tr("ErrOK")}
	Err302 = &GeneralResponse{Code: 302, Message: i18n.Tr("Err302")}
	Err404 = &GeneralResponse{Code: 404, Message: i18n.Tr("Err404")}
	Err500 = &GeneralResponse{Code: 500, Message: i18n.Tr("Err500")}

	ErrInputData = &GeneralResponse{Code: 1001, Message: i18n.Tr("ErrInputData")}
	ErrSystem    = &GeneralResponse{Code: 1002, Message: i18n.Tr("ErrSystem")}
	ErrDatabase  = &GeneralResponse{Code: 1003, Message: i18n.Tr("ErrDatabase")}
	ErrNeedLogin = &GeneralResponse{Code: 1004, Message: i18n.Tr("ErrNeedLogin")}
	ErrCSRFToken = &GeneralResponse{Code: 1005, Message: i18n.Tr("ErrCSRFToken")}
)
