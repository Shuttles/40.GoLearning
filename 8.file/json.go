package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var syncFromTo = make(map[string]string)
	srcPath1 := "/data/0/"
	srcPath2 := "/data/1/"
	destPath1 := "/data/0"
	destPath2 := "/data/1"
	syncFromTo[srcPath1] = destPath1
	syncFromTo[srcPath2] = destPath2
	syncFromToData, err := json.Marshal(syncFromTo)
	if err != nil {
		fmt.Println("Json.Marshal Error:", err)
		return
	}
	fmt.Println(syncFromToData)

	/*
		fmt.Println("==============After unMarshal===============")
		var syncFromToAfterUnmarshal = make(map[string]string)
		err = json.Unmarshal(syncFromToData, &syncFromToAfterUnmarshal)
		if err != nil {
			fmt.Println("Json Error:", err)
			return
		}
		fmt.Println(syncFromToAfterUnmarshal)
	*/

	// set env
	fmt.Printf("string(syncFromToData) = %v\n", string(syncFromToData))
	if err = os.Setenv("SyncFromTo", string(syncFromToData)); err != nil {
		fmt.Println("Set env error : ", err)
		return
	}
	fromToPathJSON := os.Getenv("SyncFromTo")
	fmt.Println("fromToPathJSON = ", fromToPathJSON)

	for {

	}
	return
}
