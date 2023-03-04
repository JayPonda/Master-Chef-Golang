package handler

import (
	cook "main/Cook"
	services "main/Services"
	"time"
)

func Handle() error {

	loggerChan := make(chan services.LogMessaage)

	loggerFileName := time.Now().Format("20060102150405")

	eventHandlerChan := make(chan cook.PostStruct)

	err := initialCall(loggerFileName, loggerChan, eventHandlerChan)

	return err
}