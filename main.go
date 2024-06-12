package main

import (
	"fmt"
	"github.com/neil-and-void/boxdb/src/memtable"
)

func main() {
	var maxHeight uint8 = 3
	skipList := memtable.New(memtable.StringComparer{}, maxHeight)

	for i := 1; i <= 3; i += 1 {
		skipList.Put(fmt.Sprintf("%d", i), fmt.Sprintf("%d-val", i))
	}

	skipList.Put("2", "Updated")

	skipList.Print()

	// fmt.Println("get")
	// val := skipList.Get("87")
	// if val != nil {
	// 	fmt.Println(*val)
	// }
}
