package handler

import (
	"fmt"
	cook "main/Cook"
	services "main/Services"
	static "main/Static"
	"sync"
	"time"
)

func Handle() error {

	// start services
	// create channel + create filename + calculate waitgroup
	loggerChan := make(chan services.LogMessaage)

	loggerFileName := time.Now().Format("20060102150405")

	eventHandlerChan := make(chan cook.PostStruct)

	err := initialCall(loggerFileName, loggerChan, eventHandlerChan)

	totalRecipy := len(static.RecipyList)

	wg := sync.WaitGroup{}

	wg.Add(totalRecipy)

	// start Contest

	startContest(&wg, eventHandlerChan, loggerChan)

	wg.Wait()

	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: -51, ChildPsId: 0, TypeOfInfo: static.End, Additional: "contest end"}

	// end services
	terminationCall(loggerChan, eventHandlerChan)

	return err
}

func Cook(stap cook.Step) {
	time.Sleep(time.Duration(stap.TimeInSec * int(time.Second)))
}

func startContest(wg *sync.WaitGroup, eventHandlerChan chan cook.PostStruct, loggerChan chan services.LogMessaage) {

	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: -51, ChildPsId: 0, TypeOfInfo: static.Start, Additional: "contest start in three seconds"}
	time.Sleep(1 * time.Second)
	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: -51, ChildPsId: 1, TypeOfInfo: static.Start, Additional: "--3--"}
	fmt.Print("\r--- contest start in 3 seconds ---")
	time.Sleep(1 * time.Second)
	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: -51, ChildPsId: 1, TypeOfInfo: static.Start, Additional: "--2--"}
	fmt.Print("\r--- contest start in 2 seconds ---")
	time.Sleep(1 * time.Second)
	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: -51, ChildPsId: 1, TypeOfInfo: static.Start, Additional: "--1--"}
	fmt.Print("\r--- contest start in 1 seconds ---\n")

	for _, recipy := range static.RecipyList {

		go ContestRecipy(recipy, wg, eventHandlerChan, loggerChan)

	}

}

func ContestRecipy(racipy cook.Recipy, wg *sync.WaitGroup, eventHandlerChan chan cook.PostStruct, loggerChan chan services.LogMessaage) {
	// get racipy
	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: racipy.Id, ChildPsId: -1, TypeOfInfo: static.StartRecipy, Additional: racipy.Name}

	// follow each stap in order
	for _, stap := range racipy.Stapes {

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
		Cook(stap)
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

	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: racipy.Id, ChildPsId: -2, TypeOfInfo: static.EndRecipy, Additional: racipy.Name + " complated!!!"}
	fmt.Printf("Id: %d, Name: %s complated!!!\n", racipy.Id, racipy.Name)
	wg.Done()
}
