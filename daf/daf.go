// Package daf implements a parser for the DAF file format.
//
// Can parse planetary position informtion from NAIF.
//
// See toolkit/doc/html/req/daf.html.
package daf

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strings"
)

// Reader reads a SPICE Double Precision Array File (DAF).
type Reader struct {
	Header FileRecord

	r     io.ReadSeeker
	next  int64
	order binary.ByteOrder
}

type FileRecord struct {
	LOCIDW string // identification word "DAF/____"
	ND     uint32 // number of float64 components in each summary
	NI     uint32 // number of int32 components in each summary
	LOCIFN string // description of file
	NC     int    // number of bytes in name summary
	SS     int    // size of a single summary
	FWARD  uint32 // number for first summary record
	BWARD  uint32 // number for last summary record
	FREE   uint32 // first free address in file
	FTPSTR []byte // checksum of sorts, see SUBROUTINE ZZFTPSTR
}

// readFileRecord reads the contents of a DAF file record.
//
// See "The File Record" in "DAF Organization".
func (r *Reader) readFileRecord() error {
	var record [1024]uint8
	if _, err := io.ReadFull(r.r, record[:]); err != nil {
		return fmt.Errorf("daf: read file record: %v", err)
	}

	switch locfmt := string(record[88:96]); locfmt {
	case "LTL-IEEE":
		r.order = binary.LittleEndian
	case "BIG-IEEE":
		r.order = binary.BigEndian
	default:
		return fmt.Errorf("daf: bad LOCFMT: %q", locfmt)
	}

	h := FileRecord{
		LOCIDW: string(record[:8]),
		ND:     r.order.Uint32(record[8:]),
		NI:     r.order.Uint32(record[12:]),
		LOCIFN: string(bytes.TrimSpace(record[16:76])), // ASCII
		FWARD:  r.order.Uint32(record[76:]),
		BWARD:  r.order.Uint32(record[80:]),
		FREE:   r.order.Uint32(record[84:]),
		FTPSTR: record[699:727],
	}
	h.SS = int(h.ND + (h.NI+1)/2)
	h.NC = 8 * h.SS
	r.Header = h

	if h.ND > 124 {
		return fmt.Errorf("daf: ND=%d, maximum value is 124", h.ND)
	}
	if h.NI > 250 {
		return fmt.Errorf("daf: NI=%d, maximum value is 250", h.NI)
	}
	if h.NI < 2 {
		return fmt.Errorf("daf: NI=%d, minimum value is 2", h.NI)
	}
	if s := h.ND + (h.NI+1)/2; s > 125 {
		return fmt.Errorf("daf: invalid ND=%d, NI=%d", h.ND, h.NI)
	}
	if !strings.HasPrefix(h.LOCIDW, "DAF/") {
		return fmt.Errorf("daf: LOCIDW %q, missing DAF/ prefix", h.LOCIDW)
	}

	r.next = int64(h.FWARD)
	return nil
}

type Summary struct {
	Name  string
	Float []float64
	Int   []int32
}

type SummaryRecord struct {
	Next    int64
	Prev    int64
	Summary []Summary
}

func (r *Reader) Read(rec int64) (*SummaryRecord, error) {
	_, err := r.r.Seek(1024*(rec-1), 0)
	if err != nil {
		return nil, err
	}
	var buf [1024]byte
	b := buf[:]
	if _, err := io.ReadFull(r.r, b); err != nil {
		return nil, err
	}
	count := int(math.Float64frombits(r.order.Uint64(b[16:])))
	s := &SummaryRecord{
		Next:    int64(math.Float64frombits(r.order.Uint64(b[0:]))),
		Prev:    int64(math.Float64frombits(r.order.Uint64(b[8:]))),
		Summary: make([]Summary, 0, count),
	}
	b = b[24:]
	for i := 0; i < count; i++ {
		sum := Summary{
			Float: make([]float64, r.Header.ND),
			Int:   make([]int32, r.Header.NI),
		}
		for j := range sum.Float {
			sum.Float[j] = math.Float64frombits(r.order.Uint64(b))
			b = b[8:]
		}
		for j := range sum.Int {
			sum.Int[j] = int32(r.order.Uint32(b))
			b = b[4:]
		}
		s.Summary = append(s.Summary, sum)
	}

	b = buf[:]
	if _, err := io.ReadFull(r.r, b); err != nil {
		return nil, err
	}
	for i := range s.Summary {
		s.Summary[i].Name = string(bytes.TrimSpace(b[:r.Header.NC]))
		b = b[r.Header.NC:]
	}

	return s, nil
}

func (r *Reader) Next() (*SummaryRecord, error) {
	if r.next == 0 {
		return nil, io.EOF
	}
	s, err := r.Read(r.next)
	if err == nil {
		r.next = s.Next
	}
	return s, err
}

func NewReader(in io.ReadSeeker) (*Reader, error) {
	r := &Reader{r: in}
	if err := r.readFileRecord(); err != nil {
		return nil, err
	}

	return r, nil
}
