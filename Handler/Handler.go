package handler

import (
	services "main/Services"
	"time"
)

func Handle() error {

	loggerChan := make(chan services.LogMessaage)

	loggerFileName := time.Now().Format("20060102150405")

	err := initialCall(loggerFileName, loggerChan)

	return err
}