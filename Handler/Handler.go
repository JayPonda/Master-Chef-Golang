package handler

import (
	"fmt"
	cook "main/Cook"
	madeAnOrder "main/RecipyStart"
	services "main/Services"
	static "main/Static"
	"math/rand"
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

	rand.Shuffle(totalRecipy, func(i, j int) {
		static.RecipyList[i], static.RecipyList[j] = static.RecipyList[j], static.RecipyList[i]
	})

	wg := sync.WaitGroup{}

	wg.Add(totalRecipy)

	// start Contest

	startContest(&wg, eventHandlerChan, loggerChan)

	wg.Wait()

	fmt.Print("----------------------------------\n")

	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: static.SystemService, ChildPsId: 0, TypeOfInfo: static.End, Additional: "contest end"}

	// end services
	terminationCall(loggerChan, eventHandlerChan)

	return err
}

func startContest(wg *sync.WaitGroup, eventHandlerChan chan cook.PostStruct, loggerChan chan services.LogMessaage) {

	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: static.SystemService, ChildPsId: 0, TypeOfInfo: static.Start, Additional: "contest start in 5 seconds"}
	fmt.Print("\r--- contest start in 5 seconds ---")
	time.Sleep(1 * time.Second) // 1
	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: static.SystemService, ChildPsId: 0, TypeOfInfo: static.Start, Additional: "contest start in 4 seconds"}
	fmt.Print("\r----------- start in 4 seconds ---")
	time.Sleep(1 * time.Second) // 2
	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: static.SystemService, ChildPsId: 0, TypeOfInfo: static.Start, Additional: "contest start in 3 seconds"}
	fmt.Print("\r----------------- in 3 seconds ---")
	time.Sleep(1 * time.Second) // 3
	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: static.SystemService, ChildPsId: 0, TypeOfInfo: static.Start, Additional: "contest start in 2 seconds"}
	fmt.Print("\r-------------------- 2 seconds ---")
	time.Sleep(1 * time.Second) // 4
	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: static.SystemService, ChildPsId: 0, TypeOfInfo: static.Start, Additional: "contest start in 1 seconds"}
	fmt.Print("\r---------------------- seconds ---")
	time.Sleep(1 * time.Second) // 5
	loggerChan <- services.LogMessaage{TimeStamp: time.Now(), ParentPsId: static.SystemService, ChildPsId: 0, TypeOfInfo: static.Start, Additional: "contest start in 0 seconds"}
	fmt.Print("\r----------------------------------\n")

	for _, recipy := range static.RecipyList {

		go madeAnOrder.ContestRecipy(recipy, wg, eventHandlerChan, loggerChan)

	}
}
