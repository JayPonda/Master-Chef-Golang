package madeAnOrder

import (
	"fmt"
	cook "main/Cook"
	services "main/Services"
	static "main/Static"
	"sync"
	"time"
)

func ContestRecipy(racipy cook.Recipy, wg *sync.WaitGroup, eventHandlerChan chan cook.PostStruct, loggerChan chan services.LogMessaage) {
	startAt := time.Now()
	expected := 0
	// get racipy
	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: racipy.Id, ChildPsId: static.RecipyStart, TypeOfInfo: static.StartRecipy, Additional: racipy.Name}
	// follow each stap in order
	for _, stap := range racipy.Stapes {

		expected += stap.TimeInSec

		// send request for resourice
		eventHandlerChan <- cook.PostStruct{MessageType: static.QueryGet, AttachedStap: &stap, Channel: racipy.CommunicationChan}
		loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: racipy.Id, ChildPsId: stap.Id, TypeOfInfo: static.Requested, Additional: stap.RequiredRes}
		racipy.CurrentStap = stap.Id

		// wait for response
		attempt := 1
		for chRes := range racipy.CommunicationChan {
			if chRes == static.AccessGranted {
				loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: racipy.Id, ChildPsId: stap.Id, TypeOfInfo: static.AccessGranteds, Additional: "attempt : " + fmt.Sprint(attempt) + ", res:" + stap.RequiredRes}
				break
			} else if chRes == static.NoAssetsAvalable {
				loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: racipy.Id, ChildPsId: stap.Id, TypeOfInfo: static.Wait, Additional: "attempt : " + fmt.Sprint(attempt) + ", res:" + stap.RequiredRes}
			} else if chRes == static.DelaiedAccessGranted {
				loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: racipy.Id, ChildPsId: stap.Id, TypeOfInfo: static.AccessGranteds, Additional: "attempt : " + fmt.Sprint(attempt) + ", res:" + stap.RequiredRes}
				break
			}
			attempt++
		}

		// cook time
		loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: racipy.Id, ChildPsId: stap.Id, TypeOfInfo: static.CookStart, Additional: "action: " + stap.Action}
		time.Sleep(time.Duration(stap.TimeInSec * int(time.Second)))
		loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: racipy.Id, ChildPsId: stap.Id, TypeOfInfo: static.CookEnd, Additional: "action: " + stap.Action}

		// send request to leave resourice
		eventHandlerChan <- cook.PostStruct{MessageType: static.QueryTake, AttachedStap: &stap, Channel: racipy.CommunicationChan}
		loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: racipy.Id, ChildPsId: stap.Id, TypeOfInfo: static.RequestReturn, Additional: "res: " + stap.RequiredRes}

		// resourice taken
		if chRes := <-racipy.CommunicationChan; chRes == static.AssetReturnSucessfully {
			loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: racipy.Id, ChildPsId: stap.Id, TypeOfInfo: static.Return, Additional: "res: " + stap.RequiredRes}
		} else if chRes == static.ErrorOnReturn {
			loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: racipy.Id, ChildPsId: stap.Id, TypeOfInfo: static.ErrorReturn, Additional: "res: " + stap.RequiredRes}
		}
	}

	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: racipy.Id, ChildPsId: static.RecipyEnd, TypeOfInfo: static.EndRecipy, Additional: racipy.Name + " complated!!!"}
	fmt.Printf("Id: %d, Name: %s complated in %d Sec vs %d Sec!!!\n", racipy.Id, racipy.Name, time.Since(startAt)/time.Second, expected)
	wg.Done()
}
