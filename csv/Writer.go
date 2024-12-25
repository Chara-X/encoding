package csv

import (
	"encoding/csv"
	"io"
)

type Writer struct{ w *csv.Writer }

func NewWriter(w io.Writer) *Writer           { return &Writer{w: csv.NewWriter(w)} }
func (w *Writer) Write(record []string) error { return w.w.Write(record) }
func (w *Writer) Flush()                      { w.w.Flush() }
