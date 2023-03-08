package static

const (

	// status
	Requested      = "Requested"
	RequestReturn  = "Request Return"
	ReRequest      = "Re-request"
	AccessGranteds = "access granted"
	Wait           = "Wait"
	CookStart      = "Cook start"
	CookEnd        = "Cook end"
	Return         = "Return"
	ErrorReturn    = "Return Error"
	Start          = "start"
	StartAck       = "ack start"
	End            = "end"
	EndAck         = "ack end"
	QueryGet       = "get res"
	QueryTake      = "take res"
	ImmidateCheck  = "imd check"
	ErrorAck       = "Error Ack"

	// description
	InitialtLogger         = "Initial logger"
	InitialEventHandler    = "Initial event-handler"
	InitialEventLooper     = "Initial event-Looper"
	StartStap              = "Stap start"
	Check                  = "Check"
	EndStap                = "Stap end"
	StartRecipy            = "Recipy start"
	EndRecipy              = "Recipy end"
	TerminatetLogger       = "Terminat logger"
	TerminateEventListener = "Terminat eListener"
	TerminateEventLooper   = "Terminat eLooper"

	// response code
	Startservice           = 0
	StartServiceAck        = 1
	EndService             = 2
	EndServiceAck          = 3
	AccessGranted          = 200
	NoAssetsAvalable       = 300
	DelaiedAccessGranted   = 400
	AssetReturnSucessfully = 500
	ErrorOnReturn          = 600

	// service code
	SystemService       = -51
	RecipyEnd           = -99
	RecipyStart         = -100
	LoggerService       = -101
	EventHandlerService = -102
)
