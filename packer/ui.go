package packer

import (
	"fmt"
	"io"
)

// The Ui interface handles all communication for Packer with the outside
// world. This sort of control allows us to strictly control how output
// is formatted and various levels of output.
type Ui interface {
	Say(format string, a ...interface{})
	Error(format string, a ...interface{})
}

// The ReaderWriterUi is a UI that writes and reads from standard Go
// io.Reader and io.Writer.
type ReaderWriterUi struct {
	Reader io.Reader
	Writer io.Writer
}

func (rw *ReaderWriterUi) Say(format string, a ...interface{}) {
	_, err := fmt.Fprintf(rw.Writer, format+"\n", a...)
	if err != nil {
		panic(err)
	}
}

func (rw *ReaderWriterUi) Error(format string, a ...interface{}) {
	_, err := fmt.Fprintf(rw.Writer, format+"\n", a...)
	if err != nil {
		panic(err)
	}
}
