package utils

import (
	slog "github.com/cihub/seelog"
)

var Logger slog.LoggerInterface

func init() {
	var err error
	Logger, err = slog.LoggerFromConfigAsFile("config/slog.xml")
	if err != nil {
		panic(err)
	}
}
