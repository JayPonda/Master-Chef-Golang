package services

import (
	cook "main/Cook"
	static "main/Static"
)

func EventLooper(eventLooperChan chan cook.PostStruct) {

	// sending ack
	eventLooperChan <- cook.PostStruct{MessageType: static.StartAck, AttachedStap: nil, Channel: nil}

}