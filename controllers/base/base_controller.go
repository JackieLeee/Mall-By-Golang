package base

import (
	"encoding/json"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"

	"mall/controllers/base/session"
)

type BaseController struct {
	web.Controller

	IsAuth   bool
	SourceIp string
}

type GeneralResponse struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (c *BaseController) Prepare() {
	c.SourceIp = c.Ctx.Input.IP()
}

func (c *BaseController) ServeResponse(errMsg *GeneralResponse, args ...interface{}) {
	if nil == errMsg {
		return
	}

	data := errMsg.Data
	if len(args) >= 1 {
		if nil != args[0] {
			data = args[0]
		}
	}

	result := &GeneralResponse{
		Code:    errMsg.Code,
		Message: errMsg.Message,
		Data:    data,
	}

	c.Data["json"] = result

	if err := c.ServeJSON(); err != nil {
		logs.Warn("base controller sends a json response failed, err: %s", err.Error())
	}
}

func (c *BaseController) MyGetSession(name string) (interface{}, bool) {
	sessItf := c.GetSession(name)
	sess, ok := sessItf.(string)
	if !ok {
		return nil, false
	}

	switch name {
	case session.KSessKeyUser:
		result := session.User{}
		err := json.Unmarshal([]byte(sess), &result)
		if err != nil {
			logs.Debug("json unmarshal session [%s] err: %s", name, err)
			return nil, false
		}
		return result, true
	}
	return nil, false
}
