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

var AssetList = struct {
	sync.Mutex
	RList  map[string]*Resourice
	OStack []*cook.Step
}{RList: map[string]*Resourice{
	"knif":   {Id: 0, Name: "knif", Quentaty: 1, ItemInUse: 0},
	"stove":  {Id: 1, Name: "stove", Quentaty: 1, ItemInUse: 0},
	"mixer":  {Id: 2, Name: "mixer", Quentaty: 1, ItemInUse: 0},
	"oven":   {Id: 3, Name: "oven", Quentaty: 1, ItemInUse: 0},
	"icebox": {Id: 4, Name: "icebox", Quentaty: 1, ItemInUse: 0},
}, OStack: []*cook.Step{},
}