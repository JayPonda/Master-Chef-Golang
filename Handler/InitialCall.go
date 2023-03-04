package handler

import (
	"errors"
	services "main/Services"
	static "main/Static"
	"time"
)

func initialCall(loggerFileName string, loggerChan chan services.LogMessaage) error {

	// logger start timestamp and start logWritter
	
	preLoggerStartMsg := services.LogMessaage{TimeStamp: time.Now(), ParentPsId: -100, ChildPsId: 0, TypeOfInfo: static.Start, Additional: static.InitialtLogger}
	
	go services.LogWritter(loggerChan, loggerFileName)

	resFromLogger := <- loggerChan
	
	if resFromLogger.TypeOfInfo == static.ErrorAck{ // if logger has an error
		
		return errors.New(resFromLogger.Additional)

	} else if resFromLogger.TypeOfInfo == static.StartAck { // logger is successfully started
		
		loggerChan <- preLoggerStartMsg
		resFromLogger.TimeStamp = time.Now()
		loggerChan <- resFromLogger
	
	}
			
	return nil
}