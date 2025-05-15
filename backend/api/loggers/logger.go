package loggers

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

func Start() {
	timeNow := time.Now()

	folder := fmt.Sprintf("%d-%d-%d", timeNow.Year(), timeNow.Month(), timeNow.Day())
	fileName := fmt.Sprintf("%d-%d", timeNow.Hour(), timeNow.Minute())

	err := logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"logs/%s/%s.log","maxdays":30}`, folder, fileName))
	if err != nil {
		logs.Info(err.Error())
	}

}
