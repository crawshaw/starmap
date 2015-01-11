// Package spk implements a parser for the binary DAF/SPK SPICE format.
package spk

import (
	"fmt"
	"io"
	"time"

	"github.com/crawshaw/starmap/daf"
	"github.com/crawshaw/starmap/naif"
)

type Reader struct {
	r *daf.Reader
	s []Segment

	// array start/end addresses for each segment
	start []int32
	end   []int32
}

func (r *Reader) Next() (*Segment, error) {
	if len(r.s) == 0 {
		return nil, io.EOF
	}
	s := r.s[0]
	// TODO: We could move this work into the DAF reader. Is that a good idea?
	// DAF specifies that "The initial and final addresses of an array are
	// always the values of the final two integer components of the summary
	// for the array" so it could be done there without specific knowledge of SPK.
	buf, err := r.r.ReadArray(r.start[0], r.end[0])
	if err != nil {
		return nil, err
	}

	r.s = r.s[1:]
	r.start = r.start[1:]
	r.end = r.end[1:]

	var data Data
	switch s.Type {
	case 2:
		data, err = newChebyshevPosOnly(buf)
	case 3:
		data, err = newChebyshev(buf)
	}
	if err != nil {
		return nil, err
	}
	s.Data = data
	return &s, nil
}

type Segment struct {
	Start    time.Time
	End      time.Time
	Target   naif.Code
	Center   naif.Code
	RefFrame naif.Code
	Type     Type
	Data     Data
	//Data     []float64
}

type Data interface {
	// At computes position and velocity at time t.
	//
	// Position is measured in kilometers from the barycenter.
	// Velocity is measured in kilometers/second.
	At(t time.Time) (pos, vel [3]float64, err error)
}

var magic = "DAF/SPK "

func NewReader(in io.ReadSeeker) (*Reader, error) {
	dr, err := daf.NewReader(in)
	if err != nil {
		return nil, err
	}
	r := &Reader{r: dr}

	if got := r.r.Header.LOCIDW; got != magic {
		return nil, fmt.Errorf("spk: header LOCIDW=%q, want %q", got, magic)
	}
	if r.r.Header.NI != 6 {
		return nil, fmt.Errorf("spk: NI=%d, SPK requires NI=6", r.r.Header.NI)
	}
	if r.r.Header.ND != 2 {
		return nil, fmt.Errorf("spk: ND=%d, SPK requires ND=2", r.r.Header.ND)
	}

	for {
		sr, err := r.r.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		for _, s := range sr.Summary {
			r.s = append(r.s, Segment{
				Start:    ftime(s.Float[0]),
				End:      ftime(s.Float[1]),
				Target:   naif.Code(s.Int[0]),
				Center:   naif.Code(s.Int[1]),
				RefFrame: naif.Code(s.Int[2]),
				Type:     Type(s.Int[3]),
			})
			r.start = append(r.start, s.Int[4])
			r.end = append(r.end, s.Int[5])
		}
	}
	return r, nil
}

var epoch = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)

func ftime(s float64) time.Time { return epoch.Add(time.Duration(s * 1e9)) }
