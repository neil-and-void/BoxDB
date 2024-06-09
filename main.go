package main

import (
	"fmt"
	"github.com/neil-and-void/boxdb/src/memtable"
)

func main() {
	fmt.Println("hello world!")

	var maxHeight uint8 = 3
	skipList := memtable.New(maxHeight)

	for i := 0; i < 10; i++ {
		skipList.Put(fmt.Sprintf("%d", i), fmt.Sprintf("%d-val", i))
	}

	skipList.Print()

	// val := skipList.Get("5")
	// fmt.Println(val)
}
