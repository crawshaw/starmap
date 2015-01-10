package daf

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestReader(t *testing.T) {
	// TODO: find/create small testdata and commit it.
	f, err := os.Open("../testdata/plu043.bsp")
	if err != nil {
		t.Fatal(err)
	}
	r, err := NewReader(f)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for {
		s, err := r.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("count=%d, err: %v", count, err)
		}
		fmt.Printf("summary: %+v\n", s)
	}
}
