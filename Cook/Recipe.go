package cook

type Recipy struct {
	Id                int
	Name              string
	Stapes            []Step
	CurrentStap       int
	CommunicationChan chan int
}
