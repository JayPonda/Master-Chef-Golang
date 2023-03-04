package cook

type PostStruct struct {
	MessageType  string
	AttachedStap *Step
	Channel      chan int
}