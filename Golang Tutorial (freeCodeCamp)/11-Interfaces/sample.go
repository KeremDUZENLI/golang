package Interfaces

import (
	"fmt"
)

func Sample() {
	var wc WriterCloser = myWriterCloser{}
	fmt.Println(wc)
}

type Writer3 interface {
	Write([]byte) (int, error)
}

type Closer3 interface {
	Close() error
}

type myWriterCloser struct{}

func (mwc myWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

func (mwc myWriterCloser) Close() error {
	return nil
}
