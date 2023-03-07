package distributions

import (
	"errors"
	cook "main/Cook"
)

func QueryToGetRes(p cook.PostStruct) bool {
	muxOccupiedResouriceArray.Lock()
	AssetList.Lock()
	defer muxOccupiedResouriceArray.Unlock()
	defer AssetList.Unlock()

	requestedRes := AssetList.RList[p.AttachedStap.RequiredRes]

	if requestedRes.Quentaty-requestedRes.ItemInUse <= 0 {
		return false
	}

	requestedRes.ItemInUse++
	occupiedResouriceArray = append(occupiedResouriceArray, p.AttachedStap)
	return true
}

func QueryToTakeRes(p cook.PostStruct) error {
	muxOccupiedResouriceArray.Lock()
	AssetList.Lock()
	defer muxOccupiedResouriceArray.Unlock()
	defer AssetList.Unlock()

	requestedRes := AssetList.RList[p.AttachedStap.RequiredRes]

	for index, occRes := range occupiedResouriceArray {
		if occRes == p.AttachedStap {
			requestedRes.ItemInUse--

			if index+1 < len(occupiedResouriceArray) {
				occupiedResouriceArray = append(occupiedResouriceArray[:index], occupiedResouriceArray[index+1:]...)
			} else {
				occupiedResouriceArray = occupiedResouriceArray[:index]
			}

			return nil
		}
	}

	return errors.New("asset hasn't aquired by requester")
}
