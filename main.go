package main

import (
	"fmt"
	"github.com/neil-and-void/boxdb/src/memtable"
)

func main() {
	fmt.Println("hello world!")

	maxHeight := 3
	skipList := memtable.New(uint8(maxHeight))

	for i := 0; i < 10; i++ {
		skipList.Put(fmt.Sprintf("%d", i), "d_val")
	}

	skipList.Print()
}
