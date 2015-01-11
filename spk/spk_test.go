package spk

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestReader(t *testing.T) {
	// TODO: find/create small testdata and commit it.
	f, err := os.Open("../testdata/plu043.bsp")
	//f, err := os.Open("../testdata/de430.bsp")
	if err != nil {
		t.Fatal(err)
	}
	r, err := NewReader(f)
	if err != nil {
		t.Fatal(err)
	}
	for {
		s, err := r.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal(err)
		}
		_ = s
		fmt.Printf("%s-%s (%s) %s\n", s.Target, s.Center, s.RefFrame, s.Type)
		// ChebyshevPosOnly, and Chebyshev.
	}
}
