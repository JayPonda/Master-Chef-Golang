package handler

import (
	"errors"
	"fmt"
	cook "main/Cook"
	services "main/Services"
	static "main/Static"
	"time"
)

func initialCall(loggerFileName string, loggerChan chan services.LogMessaage, eventHandlerChan chan cook.PostStruct) error {

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
		fmt.Println("logger start successfully")
	}

	// eventHandler start 

	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: -101, ChildPsId: 0, TypeOfInfo: static.Start, Additional: static.InitialEventHandler}

	go services.EventHandler(eventHandlerChan)

	if resFromEventHandler:= <-eventHandlerChan; resFromEventHandler.MessageType == static.StartAck{ // event-handler started sucessfully
		loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: -101, ChildPsId: 1, TypeOfInfo: static.StartAck, Additional: static.InitialEventHandler}
		fmt.Println("event-handler started successfully")
	}
			
	return nil
}