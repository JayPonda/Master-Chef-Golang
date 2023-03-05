package distributions

import (
	"errors"
	"fmt"
	cook "main/Cook"
)

func QueryToGetRes(p cook.PostStruct) bool {
	AssetList.Lock()
	defer AssetList.Unlock()
	res := AssetList.RList[p.AttachedStap.RequiredRes]
	// fmt.Println(res.Quentaty, res.ItemInUse, res.Name, res.ItemInUse < res.Quentaty, p.AttachedStap.ParentRecp, p.AttachedStap.Id)
	if res.ItemInUse < res.Quentaty {
		AssetList.OStack = append(AssetList.OStack, p.AttachedStap)
		res.ItemInUse++
		return true
	}
	return false
}

func QueryToTakeRes(p cook.PostStruct) error {
	AssetList.Lock()
	defer AssetList.Unlock()
	res := AssetList.RList[p.AttachedStap.RequiredRes]
	ostack := AssetList.OStack
	for ind, stap := range ostack {
		if stap == (p.AttachedStap) {
			fmt.Println(ostack)
			fmt.Println(p.AttachedStap.RequiredRes)		
			AssetList.OStack = append(ostack[:ind], ostack[ind+1:]...)
			res.ItemInUse--
			fmt.Println(ostack)
			fmt.Println(p.AttachedStap.RequiredRes)
		
			return nil
		}
	}
	return errors.New("Asset hasn't aquired by requester")
}
