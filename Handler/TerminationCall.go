package handler

import (
	"fmt"
	cook "main/Cook"
	services "main/Services"
	static "main/Static"
	"sync"
	"time"
)

func terminationCall(loggerChan chan services.LogMessaage, eventHandlerChan chan cook.PostStruct) error {

	var wg sync.WaitGroup
	wg.Add(1)

	// terminate eventHandler
	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: static.EventHandlerService, ChildPsId: static.EndService, TypeOfInfo: static.End, Additional: static.TerminateEventListener}
	eventHandlerChan <- cook.PostStruct{MessageType: static.End, AttachedStap: nil, Channel: nil}

	if resFromEventHandler := <-eventHandlerChan; resFromEventHandler.MessageType == static.EndAck { // eventhandler terminated
		loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: static.EventHandlerService, ChildPsId: static.EndService, TypeOfInfo: static.EndAck, Additional: static.TerminateEventListener}
		fmt.Println("event-handler terminated successfully")
	}

	// termination logger

	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: static.LoggerService, ChildPsId: static.EndService, TypeOfInfo: static.End, Additional: static.TerminatetLogger}

	if resFromLogger := <-loggerChan; resFromLogger.TypeOfInfo == static.EndAck {
		fmt.Println("logger service terminated")
	}

	wg.Done()
	wg.Wait()

	close(eventHandlerChan)
	close(loggerChan)

	return nil
}
