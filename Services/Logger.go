package services

import (
	"log"
	static "main/Static"
	"os"
	"time"
)

type LogMessaage struct {
	TimeStamp  time.Time
	ParentPsId int
	ChildPsId  int
	TypeOfInfo string
	Additional string
}

func LogWritter(loggerChan chan LogMessaage, fileName string) {

	file, err := os.OpenFile("Log"+fileName+".txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		loggerChan <- LogMessaage{time.Now(), -100, 1, static.ErrorAck, err.Error()}
		return
	}

	log.SetOutput(file)
	log.SetFlags(0)

	loggerChan <- LogMessaage{time.Now(), static.LoggerService, static.StartServiceAck, static.StartAck, static.InitialtLogger}

	for reqToLog := range loggerChan {
		if reqToLog.TypeOfInfo == static.End && reqToLog.ParentPsId == static.LoggerService {
			log.Printf("%37s %4d %4d %q %s\n", reqToLog.TimeStamp.UTC(), reqToLog.ParentPsId, reqToLog.ChildPsId, reqToLog.TypeOfInfo, reqToLog.Additional)
			loggerChan <- LogMessaage{time.Now(), reqToLog.ParentPsId, static.EndServiceAck, static.EndAck, ""}
			return
		}
		log.Printf("%37s %4d %4d %q %s\n", reqToLog.TimeStamp.UTC(), reqToLog.ParentPsId, reqToLog.ChildPsId, reqToLog.TypeOfInfo, reqToLog.Additional)
	}

}
