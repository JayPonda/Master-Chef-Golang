package static

import (
	cook "main/Cook"
)

var RecipyList = []cook.Recipy{
	{Id: 0, Name: "Recipy 1", Stapes: []cook.Step{
		{Id: 0, ParentRecp: 0, Action: "cutting vegies", RequiredRes: "knif", TimeInSec: 4},
		{Id: 1, ParentRecp: 0, Action: "cook vegis", RequiredRes: "stove", TimeInSec: 3},
		{Id: 2, ParentRecp: 0, Action: "cutting coriender", RequiredRes: "knif", TimeInSec: 4},
	}, CurrentStap: -1, CommunicationChan: make(chan int)}, // 11 sec

	{Id: 1, Name: "Recipy 2", Stapes: []cook.Step{
		{Id: 0, ParentRecp: 1, Action: "cutting vegies", RequiredRes: "knif", TimeInSec: 5},
		{Id: 1, ParentRecp: 1, Action: "cook vegis", RequiredRes: "stove", TimeInSec: 2},
		{Id: 2, ParentRecp: 1, Action: "cutting coriender", RequiredRes: "knif", TimeInSec: 3},
		{Id: 3, ParentRecp: 1, Action: "mix coriender", RequiredRes: "stove", TimeInSec: 1},
	}, CurrentStap: -1, CommunicationChan: make(chan int)}, // 11 sec

	{Id: 2, Name: "Recipy 3", Stapes: []cook.Step{
		{Id: 0, ParentRecp: 2, Action: "cutting vegies", RequiredRes: "knif", TimeInSec: 7},
		{Id: 1, ParentRecp: 2, Action: "griend", RequiredRes: "mixer", TimeInSec: 2},
		{Id: 2, ParentRecp: 2, Action: "cook", RequiredRes: "stove", TimeInSec: 2},
	}, CurrentStap: -1, CommunicationChan: make(chan int)}, // 11 sec

	{Id: 3, Name: "Recipy 4", Stapes: []cook.Step{
		{Id: 0, ParentRecp: 3, Action: "cutting fruits", RequiredRes: "knif", TimeInSec: 3},
		{Id: 1, ParentRecp: 3, Action: "make shake", RequiredRes: "mixer", TimeInSec: 4},
		{Id: 2, ParentRecp: 3, Action: "cool it", RequiredRes: "icebox", TimeInSec: 4},
	}, CurrentStap: -1, CommunicationChan: make(chan int)}, // 11 sec
}
