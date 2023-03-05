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

	loggerChan <- services.LogMessaage{time.Now(), -102, 0, static.End, "contest end"}

	// end services

	return err
}

func Cook(stap cook.Step) {
	time.Sleep(time.Duration(stap.TimeInSec * int(time.Second)))
}

func startContest(wg *sync.WaitGroup, eventHandlerChan chan cook.PostStruct, loggerChan chan services.LogMessaage){

	loggerChan <- services.LogMessaage{time.Now(), -102, 0, static.Start, "contest start in three seconds"}
	time.Sleep(1 * time.Second)
	loggerChan <- services.LogMessaage{time.Now(), -102, 1, static.Start, "--3--"}
	fmt.Println("--3--")
	time.Sleep(1 * time.Second)
	loggerChan <- services.LogMessaage{time.Now(), -102, 1, static.Start, "--2--"}
	fmt.Println("--2--")
	time.Sleep(1 * time.Second)
	loggerChan <- services.LogMessaage{time.Now(), -102, 1, static.Start, "--1--"}
	fmt.Println("--1--")

	for _, recipy := range static.RecipyList{

		go ContestRecipy(recipy, wg, eventHandlerChan, loggerChan)

	}

}

func ContestRecipy(racipy cook.Recipy, wg *sync.WaitGroup, eventHandlerChan chan cook.PostStruct, loggerChan chan services.LogMessaage){
	// get racipy
	loggerChan <- services.LogMessaage{time.Now(), racipy.Id, -1, static.StartRecipy, racipy.Name}

	// follow each stap in order
	for _, stap := range racipy.Stapes {

		fmt.Println(racipy.Id, stap.Id)

		// send request for resourice
		eventHandlerChan <- cook.PostStruct{MessageType: static.QueryGet, AttachedStap: &stap, Channel: racipy.CommunicationChan}
		loggerChan <- services.LogMessaage{time.Now(), racipy.Id, stap.Id, static.Requested, stap.RequiredRes}
		racipy.CurrentStap = stap.Id

		// wait for response
		attempt := 1
		for chRes := range racipy.CommunicationChan {
			if chRes == static.AccessGranted {
				loggerChan <- services.LogMessaage{time.Now(), racipy.Id, stap.Id, static.AccessGranteds, "attempt : " + fmt.Sprint(attempt)  + ", res:"+ stap.RequiredRes }
				break
			} else if chRes == static.NoAssetsAvalable {
				loggerChan <- services.LogMessaage{time.Now(), racipy.Id, stap.Id, static.Wait, "attempt : " + fmt.Sprint(attempt) + ", res:"+ stap.RequiredRes }
			} else if chRes == static.DelaiedAccessGranted {
				loggerChan <- services.LogMessaage{time.Now(), racipy.Id, stap.Id, static.AccessGranteds, "attempt : " + fmt.Sprint(attempt) + ", res:"+ stap.RequiredRes }
			}
			attempt++
		}

		// cook time
		loggerChan <- services.LogMessaage{time.Now(), racipy.Id, stap.Id, static.CookStart, "action: " + stap.Action }
		Cook(stap)
		loggerChan <- services.LogMessaage{time.Now(), racipy.Id, stap.Id, static.CookEnd, "action: " + stap.Action }

		// send request to leave resourice
		eventHandlerChan <- cook.PostStruct{MessageType: static.QueryTake, AttachedStap: &stap, Channel: racipy.CommunicationChan}
		loggerChan <- services.LogMessaage{time.Now(), racipy.Id, stap.Id, static.CookEnd, "action: " + stap.Action }

		// resourice taken
		if chRes := <-racipy.CommunicationChan; chRes == static.AssetReturnSucessfully {
			loggerChan <- services.LogMessaage{time.Now(), racipy.Id, stap.Id, static.Return, "res: "+ stap.RequiredRes }
		} else if chRes == static.ErrorOnReturn {
			loggerChan <- services.LogMessaage{time.Now(), racipy.Id, stap.Id, static.ErrorReturn, "res: " + stap.RequiredRes  }
		}
	}

	loggerChan <- services.LogMessaage{time.Now(), racipy.Id, -2, static.EndRecipy, racipy.Name + " complated!!!"}
	
	wg.Done()
}