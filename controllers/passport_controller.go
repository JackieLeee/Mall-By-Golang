package controllers

import (
	"image/color"

	"github.com/beego/beego/v2/core/logs"
	"github.com/mojocn/base64Captcha"

	"mall/controllers/base"
	"mall/controllers/base/session"
	"mall/logics"
	"mall/messages"
)

type PassportController struct {
	base.BaseController
}

var (
	CaptchaStore = base64Captcha.DefaultMemStore
)

func (c *PassportController) UserRegister() {
	request := messages.UserRegisterReq{}

	if errResp := c.InputCheck(&request); errResp != nil {
		c.ServeResponse(errResp)
		return
	}

	//verify captcha
	if !c.verifyCaptcha(request.Captcha) {
		c.ServeResponse(ErrCaptchaInvalid)
		return
	}

	//check user exist
	exist, err := logics.IsUserExist(request.Username)
	if err != nil {
		c.ServeResponse(base.ErrDatabase)
		return
	}
	if exist {
		c.ServeResponse(ErrUserExist)
		return
	}

	//create user
	ok := logics.CreateNewUser(request.Username, request.Password)
	if !ok {
		c.ServeResponse(base.ErrDatabase)
		return
	}

	c.ServeResponse(base.ErrOK)
}

func (c *PassportController) NewCaptcha() {
	base64, err := c.genAndSetCaptcha()
	if err != nil {
		logs.Warn("generate captcha failed, err: %s", err.Error())
		c.ServeResponse(base.ErrSystem)
		return
	}

	c.ServeResponse(base.ErrOK, map[string]string{"captchaBase64": base64})
}

func (c *PassportController) genAndSetCaptcha() (string, error) {
	var driver base64Captcha.Driver
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor:         &color.RGBA{R: 3, G: 102, B: 214, A: 125},
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, CaptchaStore)

	id, base64, err := captcha.Generate()
	if err != nil {
		logs.Warn("generate captcha failed, err: %s", err.Error())
		return "", err
	}

	if err := c.SetSession(session.KSessCaptcha, id); err != nil {
		logs.Warn("set captcha to session failed, err: %s", err.Error())
		return "", err
	}

	return base64, nil
}

func (c *PassportController) verifyCaptcha(captcha string) bool {
	captchaId := c.GetSession(session.KSessCaptcha)
	return CaptchaStore.Verify(captchaId.(string), captcha, true)
}
