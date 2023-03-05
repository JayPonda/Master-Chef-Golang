package static

const (

	// status
	Requested     = "Requested"
	ReRequest     = "Re-request"
	Wait          = "Wait"
	CookStart     = "Cook start"
	CookEnd       = "Cook end"
	Return        = "Return"
	Start         = "start"
	StartAck      = "ack start"
	End           = "end"
	EndAck        = "ack end"
	QueryGet      = "get res"
	QueryTake     = "take res"
	ImmidateCheck = "imd check"
	ErrorAck      = "Error Ack"

	// description
	InitialtLogger         = "Initial logger"
	InitialEventHandler    = "Initial event-handler"
	InitialEventLooper     = "Initial event-Looper"
	StartStap              = "Stap start"
	Check                  = "Check"
	EndStap                = "Stap end"
	StartRecipy            = "Recipy start"
	EndRecipy              = "Recipy end"
	TerminateStartLogger   = "Terminat logger"
	TerminateEventListener = "Terminat eListener"
	TerminateEventLooper   = "Terminat eLooper"

	// response code
	AccessGranted        = 200
	NoAssetsAvalable     = 300
	DelaiedAccessGranted = 400
)