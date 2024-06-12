package wal

import (
	"fmt"
	"io"
	"os"
)

type WalWriter interface {
	// Write(p []byte) (n int, err error)
	io.Writer
	Sync() error
}

type WalManager struct {
	wal         WalWriter
	blockLength uint
	dir         string
}

func NewWalManager(path string) *WalManager {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)

		if err != nil {
			panic(err)
		}
	}

	walFile, err := os.OpenFile(fmt.Sprintf("%s/0.wal", path), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("could not open wal dir")
	}

	return &WalManager{walFile, uint(8), path}
}

func (wm *WalManager) serialize(k, v string) []byte {
	record := fmt.Sprintf("%s:%s", k, v)
	return []byte(record)
}

func (wm *WalManager) Write(k, v string) (int, error) {
	data := wm.serialize(k, v)

	n, err := wm.wal.Write(data)
	if err != nil {
		return 0, err
	}

	err = wm.wal.Sync()
	if err != nil {
		return 0, err
	}

	return n, err
}
