package main

import (
	"github.com/neil-and-void/boxdb/src"
)

func main() {
	options := src.Options{Path: "./boxdb/data", MaxLSMHeight: uint8(3)}
	boxdb := src.NewBoxDB(options)

	boxdb.Put("uuid1234", "james@hi.com")
}
