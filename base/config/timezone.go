package config

import (
	"fmt"
	"github.com/iikmaulana/gateway/libs"
	"github.com/iikmaulana/gateway/libs/helper"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"time"
)

func (c *Config) InitTimezone() serror.SError {
	loc := helper.Env(libs.AppTimezone, "Asia/Jakarta")
	local, err := time.LoadLocation(loc)
	if err != nil {
		return serror.NewFromErrorc(err, fmt.Sprintf("failed load location %s", loc))
	}
	time.Local = local
	return nil
}
