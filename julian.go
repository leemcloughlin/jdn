/*
Copyright (C) 2015 Lee McLoughlin, LMMR Tech Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Convert to and from Julian Day Numbers
//
// Note uses only integer arithmetic
package jdn

import (
	"time"
)

// ToNumber converts a year, month, day into a Julian Day Number.
// Based on http://en.wikipedia.org/wiki/Julian_day#Calculation.
// Only valid for dates after the Julian Day epoch which is January 1, 4713 BCE.
func ToNumber(year int, month time.Month, day int) (julianDay int) {
	if year <= 0 {
		year++
	}
	a := int(14-month) / 12
	y := year + 4800 - a
	m := int(month) + 12*a - 3
	julianDay = int(day) + (153*m+2)/5 + 365*y + y/4
	if year > 1582 || (year == 1582 && (month > time.October || (month == time.October && day >= 15))) {
		return julianDay - y/100 + y/400 - 32045
	} else {
		return julianDay - 32083
	}
}

// FromNumber converts a Julian Day Number to a year, month, day in the
// appropriate calendar (Julian or Gregorian).
func FromNumber(julianDay int) (year int, month time.Month, day int) {
	if julianDay >= 2299161 {
		// Gregorian calendar starting from October 15, 1582
		// This algorithm is from Henry F. Fliegel and Thomas C. Van Flandern
		ell := julianDay + 68569
		n := (4 * ell) / 146097
		ell = ell - (146097*n+3)/4
		i := (4000 * (ell + 1)) / 1461001
		ell = ell - (1461*i)/4 + 31
		j := (80 * ell) / 2447
		day = int(ell - (2447*j)/80)
		ell = j / 11
		month = time.Month(j + 2 - (12 * ell))
		year = int(100*(n-49) + i + ell)
	} else {
		// Julian calendar until October 4, 1582
		// Algorithm from Frequently Asked Questions about Calendars by Claus Toendering
		julianDay += 32082
		dd := (4*julianDay + 3) / 1461
		ee := julianDay - ((1461 * dd) / 4)
		mm := ((5 * ee) + 2) / 153
		day = ee - (153*mm+2)/5 + 1
		month = time.Month(mm + 3 - 12*(mm/10))
		year = dd - 4800 + (mm / 10)
		if year <= 0 {
			year--
		}
	}
	return year, month, day
}
