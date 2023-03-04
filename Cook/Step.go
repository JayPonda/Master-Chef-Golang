package cook

type Step struct {
	Id            int
	ParentRecp    int
	Action        string
	RequiredRes   string
	TimeInSec     int
	CurrentStatus int
}