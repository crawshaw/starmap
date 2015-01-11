package spk

import (
	"fmt"
	"math"
	"time"
)

// Chebyshev contains position and velocity data encoded as Chebyshev polynomials.
//
// SPK Type 3: ftp://naif.jpl.nasa.gov/pub/naif/toolkit_docs/C/req/spk.html#Type%203:%20Chebyshev%20(position%20and%20velocity)
type Chebyshev struct {
	Start       time.Time
	Duration    time.Duration
	Records     []float64
	RecordCount int
	RecordSize  int
}

func (c *Chebyshev) At(t time.Time) (pos, vel [3]float64, err error) {
	panic("TODO")
}

// Chebyshev contains position-only data encoded as Chebyshev polynomials.
//
// SPK Type 2: ftp://naif.jpl.nasa.gov/pub/naif/toolkit_docs/C/req/spk.html#Type%202:%20Chebyshev%20(position%20only)
type ChebyshevPosOnly struct {
	Chebyshev
}

func (c *ChebyshevPosOnly) At(t time.Time) (pos, vel [3]float64, err error) {
	panic("TODO")
}

func newChebyshev(d []float64) (*Chebyshev, error) {
	var (
		recInit   = d[len(d)-4]
		recIntlen = d[len(d)-3]
		size      = d[len(d)-2]
		count     = d[len(d)-1]
	)
	if _, frac := math.Modf(size); frac != 0 {
		return nil, fmt.Errorf("spk: Chebyshev: invalid record size: %f", size)
	}
	if _, frac := math.Modf(count); frac != 0 {
		return nil, fmt.Errorf("spk: Chebyshev: invalid record count: %f", count)
	}

	c := &Chebyshev{
		Start:       ftime(recInit),
		Duration:    time.Duration(recIntlen) * time.Second,
		Records:     d[:len(d)-4],
		RecordCount: int(count),
		RecordSize:  int(size),
	}

	if c.RecordCount*c.RecordSize != len(c.Records) {
		return nil, fmt.Errorf("spk: Chebyshev: rec count=%d, size=%d does not fit data len %d", c.RecordCount, c.RecordSize, len(c.Records))
	}
	return c, nil
}

func newChebyshevPosOnly(d []float64) (*ChebyshevPosOnly, error) {
	c, err := newChebyshev(d)
	if err != nil {
		return nil, err
	}
	return &ChebyshevPosOnly{*c}, nil
}
