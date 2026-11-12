package tle

import (
	"math"
	"testing"
)

func almostEqual(a, b, eps float64) bool {
	return math.Abs(a-b) < eps
}

func TestParseBStar(t *testing.T) {
	cases := []struct {
		in   string
		want float64
	}{
		{"10196-3", 0.00010196},
		{" 10196-3", 0.00010196},
		{"-11606-4", -0.000011606},
		{"00000+0", 0},
		{"12345+1", 0.12345 * 10},
	}
	for _, c := range cases {
		got := parseBStar(c.in)
		if !almostEqual(got, c.want, 1e-12) {
			t.Errorf("parseBStar(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestParseTLE_ISS(t *testing.T) {
	line1 := "1 25544U 98067A   26278.49823669  .00005123  00000+0  10196-3 0  9992"
	line2 := "2 25544  51.6314 114.1795 0006877 225.7345 134.3079 15.48742155588859"

	got, err := ParseTLE([]string{line1, line2})
	if err != nil {
		t.Fatalf("ParseTLE returned error: %v", err)
	}

	if got.SatelliteNum != 25544 {
		t.Errorf("SatelliteNum = %d, want 25544", got.SatelliteNum)
	}
	if got.Classification != "U" {
		t.Errorf("Classification = %q, want U", got.Classification)
	}
	if got.EpochYear != 26 {
		t.Errorf("EpochYear = %d, want 26", got.EpochYear)
	}
	if !almostEqual(got.EpochDays, 278.49823669, 1e-8) {
		t.Errorf("EpochDays = %v, want 278.49823669", got.EpochDays)
	}
	if !almostEqual(got.Inclination, 51.6314, 1e-4) {
		t.Errorf("Inclination = %v, want 51.6314", got.Inclination)
	}
	if !almostEqual(got.RightAscension, 114.1795, 1e-4) {
		t.Errorf("RightAscension = %v, want 114.1795", got.RightAscension)
	}
	if !almostEqual(got.Eccentricity, 0.0006877, 1e-7) {
		t.Errorf("Eccentricity = %v, want 0.0006877", got.Eccentricity)
	}
	if !almostEqual(got.ArgPerigee, 225.7345, 1e-4) {
		t.Errorf("ArgPerigee = %v, want 225.7345", got.ArgPerigee)
	}
	if !almostEqual(got.MeanAnomaly, 134.3079, 1e-4) {
		t.Errorf("MeanAnomaly = %v, want 134.3079", got.MeanAnomaly)
	}
	if !almostEqual(got.MeanMotion, 15.48742155, 1e-8) {
		t.Errorf("MeanMotion = %v, want 15.48742155", got.MeanMotion)
	}
	if !almostEqual(got.BStar, 0.00010196, 1e-10) {
		t.Errorf("BStar = %v, want 0.00010196", got.BStar)
	}
}

func TestParseTLE_WrongLineCount(t *testing.T) {
	if _, err := ParseTLE([]string{"only one"}); err == nil {
		t.Error("expected error for one line, got nil")
	}
	if _, err := ParseTLE([]string{"a", "b", "c"}); err == nil {
		t.Error("expected error for three lines, got nil")
	}
}

func TestParseTLE_ShortLine(t *testing.T) {
	short := "1 25544U 98067A"
	line2 := "2 25544  51.6314 114.1795 0006877 225.7345 134.3079 15.48742155588859"
	if _, err := ParseTLE([]string{short, line2}); err == nil {
		t.Error("expected error for short line, got nil")
	}
}

func TestParseTLE_MismatchedCatalog(t *testing.T) {
	line1 := "1 25544U 98067A   26278.49823669  .00005123  00000+0  10196-3 0  9992"
	line2 := "2 99999  51.6314 114.1795 0006877 225.7345 134.3079 15.48742155588859"
	if _, err := ParseTLE([]string{line1, line2}); err == nil {
		t.Error("expected error for mismatched catalog numbers, got nil")
	}
}

func TestParseTLE_BadChecksum(t *testing.T) {
	line1 := "1 25544U 98067A   26278.49823669  .00005123  00000+0  10196-3 0  9990"
	line2 := "2 25544  51.6314 114.1795 0006877 225.7345 134.3079 15.48742155588859"
	if _, err := ParseTLE([]string{line1, line2}); err == nil {
		t.Error("expected error for bad checksum, got nil")
	}
}
