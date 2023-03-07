package services

import (
	cook "main/Cook"
	distributions "main/Distributions"
	static "main/Static"
	"time"
)

func EventHandler(eventHandlerChan chan cook.PostStruct) {

	// confirm communication
	eventHandlerChan <- cook.PostStruct{MessageType: static.StartAck, AttachedStap: nil, Channel: nil}

	// recive request
	for query := range eventHandlerChan {

		// query to get
		if query.MessageType == static.QueryGet {
			if ans := distributions.QueryToGetRes(query); ans {
				query.Channel <- static.AccessGranted
				continue
			}
			go func(host cook.PostStruct) {
				for {
					if ans := distributions.QueryToGetRes(host); ans {
						host.Channel <- static.DelaiedAccessGranted
						return
					}

					host.Channel <- static.NoAssetsAvalable
					time.Sleep(500 * time.Millisecond)
				}
			}(query)
		} else if query.MessageType == static.QueryTake {
			err := distributions.QueryToTakeRes(query)
			if err != nil {
				query.Channel <- static.ErrorOnReturn
			} else {
				query.Channel <- static.AssetReturnSucessfully
			}
		} else if query.MessageType == static.End {
			eventHandlerChan <- cook.PostStruct{MessageType: static.EndAck, AttachedStap: nil, Channel: nil}
			return
		}
	}
}
