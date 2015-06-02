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

// Test that to/from Julian Day Number works

package jdn

import (
	"fmt"
	"testing"
	"time"
)


func ymd(y int, m time.Month, d int) string {
	return fmt.Sprintf("%d-%s-%d", y, m, d)
}

func testJulian(t *testing.T, y int, m time.Month, d int) {
	jd := ToNumber(y, m, d)
	y2, m2, d2 := FromNumber(jd)

	if y != y2 {
		t.Errorf("Year didn't convert in %s out %s", ymd(y, m, d), ymd(y2, m2, d2))
		return
	}
	if m != m2 {
		t.Errorf("Month didn't convert in %s out %s", ymd(y, m, d), ymd(y2, m2, d2))
		return
	}
	if d != d2 {
		t.Errorf("Day didn't convert in %s out %s", ymd(y, m, d), ymd(y2, m2, d2))
		return
	}
	t.Logf("OK %s %d %s", ymd(y, m, d), jd, ymd(y2, m2, d2))
}

func TestJulian001(t *testing.T) {
	testJulian(t, 2012, time.August, 17) // Day I met my girlfriend
}

func TestJulian002(t *testing.T) {
	testJulian(t, 1970, time.January, 1) // Unix epoch
}

func TestJulian003(t *testing.T) {
	testJulian(t, 1, time.January, 1) // Start of the Common Era
}

func TestJulian004(t *testing.T) {
	testJulian(t, -1, time.January, 1) // Before the Common Era
}

func TestJulian005(t *testing.T) {
	testJulian(t, -4000, time.January, 1) // Uruk founded... probably
}

func TestJulian006(t *testing.T) {
	testJulian(t, -4000, time.December, 31) // Uruk 1st countdown to new years... less probably
}
