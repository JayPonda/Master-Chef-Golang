package static

// import "sync"

// var RecipyList = []Recipy{
// 	{Id: 0, Name: "Recipy 1", Stapes: []Stape{
// 		{Id: 0, ParentRecp: 0, Action: "cutting vegies", RequiredRes: "knif", TimeInSec: 3, CurrentStatus: -1},
// 		{Id: 1, ParentRecp: 0, Action: "cook vegis", RequiredRes: "stove", TimeInSec: 4, CurrentStatus: -1},
// 		{Id: 2, ParentRecp: 0, Action: "cutting coriender", RequiredRes: "knif", TimeInSec: 2, CurrentStatus: -1},
// 	}, CurrentStap: -1, CommunicationChan: make(chan int)},

// 	{Id: 1, Name: "Recipy 2", Stapes: []Stape{
// 		{Id: 0, ParentRecp: 1, Action: "cutting vegies", RequiredRes: "knif", TimeInSec: 5, CurrentStatus: -1},
// 		{Id: 1, ParentRecp: 1, Action: "cook vegis", RequiredRes: "stove", TimeInSec: 2, CurrentStatus: -1},
// 		{Id: 2, ParentRecp: 1, Action: "cutting coriender", RequiredRes: "knif", TimeInSec: 3, CurrentStatus: -1},
// 		{Id: 3, ParentRecp: 1, Action: "mix coriender", RequiredRes: "stove", TimeInSec: 1, CurrentStatus: -1},
// 	}, CurrentStap: -1, CommunicationChan: make(chan int)},

// 	{Id: 2, Name: "Recipy 3", Stapes: []Stape{
// 		{Id: 0, ParentRecp: 2, Action: "cutting vegies", RequiredRes: "knif", TimeInSec: 3, CurrentStatus: -1},
// 		{Id: 1, ParentRecp: 2, Action: "griend", RequiredRes: "mixer", TimeInSec: 4, CurrentStatus: -1},
// 		{Id: 2, ParentRecp: 2, Action: "cook", RequiredRes: "stove", TimeInSec: 4, CurrentStatus: -1},
// 	}, CurrentStap: -1, CommunicationChan: make(chan int)},

// 	{Id: 3, Name: "Recipy 4", Stapes: []Stape{
// 		{Id: 0, ParentRecp: 3, Action: "cutting fruits", RequiredRes: "knif", TimeInSec: 2, CurrentStatus: -1},
// 		{Id: 1, ParentRecp: 3, Action: "make shake", RequiredRes: "mixer", TimeInSec: 4, CurrentStatus: -1},
// 		{Id: 2, ParentRecp: 3, Action: "cool it", RequiredRes: "icebox", TimeInSec: 2, CurrentStatus: -1},
// 	}, CurrentStap: -1, CommunicationChan: make(chan int)},
// }

// var ResList = struct {
// 	sync.Mutex
// 	RList  map[string]*Resourice
// 	OStack []*Stape
// }{RList: map[string]*Resourice{
// 	"knif":   {Id: 0, Name: "knif", Quentaty: 1, ItemInUse: 0},
// 	"stove":  {Id: 1, Name: "stove", Quentaty: 1, ItemInUse: 0},
// 	"mixer":  {Id: 2, Name: "mixer", Quentaty: 1, ItemInUse: 0},
// 	"oven":   {Id: 3, Name: "oven", Quentaty: 1, ItemInUse: 0},
// 	"icebox": {Id: 4, Name: "icebox", Quentaty: 1, ItemInUse: 0},
// }, OStack: []*Stape{},
// }