package external

import (
	"io"
	"log"
)

type (
	// Writer is an optionally-prefixed writer for
	// 3rd-party logging packages.
	Writer struct {
		log *log.Logger
	}
)

// NewWriter wraps an io.Writer with a *log.Logger which doesn't
// add any prefixing by default.
func NewWriter(out io.Writer) *Writer {
	return &Writer{
		log: log.New(out, "", 0),
	}
}

// SetOutput updates the embedded *log.Logger's output
func (w *Writer) SetOutput(out io.Writer) {
	w.log.SetOutput(out)
}

// Prefix optionally sets a logging prefix for the writer
func (w *Writer) Prefix(prefix string) *Writer {
	w.log.SetPrefix(prefix)
	return w
}

func (w *Writer) Write(data []byte) (int, error) {
	msg := string(data)
	w.log.Output(3, msg)

	return len(data), nil
}