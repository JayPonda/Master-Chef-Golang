package services

import (
	"log"
	static "main/Static"
	"os"
	"time"
)

type LogMessaage struct {
	TimeStamp  time.Time
	ParentPsId   int
	ChildPsId     int
	TypeOfInfo string
	Additional string
}

func LogWritter(loggerChan chan LogMessaage, fileName string) {

	file, err := os.OpenFile("Log" + fileName + ".txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		loggerChan <- LogMessaage{time.Now(), -100, 1, static.ErrorAck, err.Error()}
		return
	}

	log.SetOutput(file)
	log.SetFlags(0)
	
	loggerChan <- LogMessaage{time.Now(), -100, 1, static.StartAck, static.InitialtLogger}
	
	for reqToLog := range loggerChan{
		log.Printf("%s %-4d %-4d %q %s\n", reqToLog.TimeStamp, reqToLog.ParentPsId, reqToLog.ChildPsId, reqToLog.TypeOfInfo, reqToLog.Additional)
	}

}
