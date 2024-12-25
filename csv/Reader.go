package csv

import (
	"encoding/csv"
	"io"
)

type Reader struct{ r *csv.Reader }

func NewReader(r io.Reader) *Reader                  { return &Reader{r: csv.NewReader(r)} }
func (r *Reader) Read() (record []string, err error) { return r.r.Read() }
