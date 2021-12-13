package i18n

import (
	"github.com/beego/i18n"
)

const (
	LangEnglish = "en_US"
	LangChinese = "zh_CN"
)

const (
	defaultLang = LangChinese
)

var (
	appLang = defaultLang
)

func Tr(key string, args ...interface{}) string {
	return i18n.Tr(appLang, key, args...)
}
