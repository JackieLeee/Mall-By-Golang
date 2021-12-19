package i18n

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/i18n"
)

const (
	LangEnglish = "en_US"
	LangChinese = "zh_CN"
)

const (
	defaultLang = LangChinese
	defaultPath = "conf/locale_zh-CN.ini"
)

var (
	appLang = defaultLang
)

func init() {
	if err := i18n.SetMessage(defaultLang, defaultPath); err != nil {
		logs.Warn("set i18n message failed")
	}
}

func Tr(key string, args ...interface{}) string {
	return i18n.Tr(appLang, key, args...)
}
