package models

import (
	"errors"

	"nwaycn/fssbc/log"
)

type XMLProfile struct {
	DB DbBase
}

func (this *XMLProfile) GetDomains() string {
	fspath := beego.AppConfig.String("fspath")
	varspath := fspath + "conf/vars.xml"

}
