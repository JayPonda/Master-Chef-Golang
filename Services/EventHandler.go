package services

import (
	cook "main/Cook"
	static "main/Static"
	"time"
)

func EventHandler(eventHandlerChan chan cook.PostStruct, loggerChan chan LogMessaage) {

	// make channel to communicate between eventHandler and eventLooper
	eventLooperChan := make(chan cook.PostStruct)

	loggerChan <- LogMessaage{time.Now(), -102, 0, static.Start, static.InitialEventLooper}
	go EventLooper(eventLooperChan)

	// confirm communication
	if looperRes := <- eventLooperChan; looperRes.MessageType == static.StartAck{
		loggerChan <- LogMessaage{time.Now(), -102, 1, static.StartAck, static.InitialEventLooper}
		eventHandlerChan <- cook.PostStruct{MessageType: static.StartAck, AttachedStap: nil, Channel: nil}
	}

}