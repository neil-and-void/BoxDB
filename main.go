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
		skipList.Put(fmt.Sprintf("%d", i), "d_val")
	}

	skipList.Print()
}
