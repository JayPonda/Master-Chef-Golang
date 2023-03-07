package distributions

import (
	cook "main/Cook"
	"sync"
)

type Resourice struct {
	Id        int
	Name      string
	Quentaty  int
	ItemInUse int
}

var muxOccupiedResouriceArray sync.Mutex
var occupiedResouriceArray = []*cook.Step{}

var AssetList = struct {
	sync.Mutex
	RList map[string]*Resourice
}{RList: map[string]*Resourice{
	"knif":   {Id: 0, Name: "knif", Quentaty: 2, ItemInUse: 0},
	"stove":  {Id: 1, Name: "stove", Quentaty: 1, ItemInUse: 0},
	"mixer":  {Id: 2, Name: "mixer", Quentaty: 2, ItemInUse: 0},
	"oven":   {Id: 3, Name: "oven", Quentaty: 1, ItemInUse: 0},
	"icebox": {Id: 4, Name: "icebox", Quentaty: 2, ItemInUse: 0},
},
}
