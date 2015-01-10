// Package naif names celestial bodies following the NAIF ID code system.
//
// The code system is described in the document "NAIF integer ID codes"
// that is distributed as part of the JPL SPICE toolkit. Anversion of
// the document can be viewed directly at:
//
//	ftp://naif.jpl.nasa.gov/pub/naif/toolkit_docs/FORTRAN/req/naif_ids.html
//
// The full set of codes for asteroids and comets is kept in the JPL
// Small-Body Database, available online:
//
//	http://ssd.jpl.nasa.gov/sbdb.cgi
package naif

import "fmt"

// Code is a Navigation and Ancillary Information Facility (NAIF) ID code.
type Code int32

func (c Code) String() string {
	n := names[c]
	if n == "" {
		return fmt.Sprintf("naif.Code(%d):UNKNOWN", c)
	}
	return n
}

// Register registers a NAIF ID code.
func Register(c Code, name string) {
	if n, ok := names[c]; ok {
		panic(fmt.Sprintf("naif.Register(%d, %q) already registered as %q", c, name, n))
	}
	names[c] = name
	Codes[name] = c
}
