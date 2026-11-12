package tle

import (
	//"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"math"
	"github.com/TylerWerne40/sat-track/internal/routing"
	"github.com/ianlopshire/go-fixedwidth"
)

type line1 struct {
	SatelliteNum   int     `fixed:"3,7"`
	Classification string  `fixed:"8,8"`
	EpochYear      int     `fixed:"19,20"`
	EpochDays      float64 `fixed:"21,32"`
	BStar          string  `fixed:"54,61"`
}

type line2 struct {
	SatelliteNum   int     `fixed:"3,7"`
	Inclination    float64 `fixed:"9,16"`
	RightAscension float64 `fixed:"18,25"`
	Eccentricity   string  `fixed:"27,33"`
	ArgPerigee     float64 `fixed:"35,42"`
	MeanAnomaly    float64 `fixed:"44,51"`
	MeanMotion     float64 `fixed:"53,63"`
}

	
func parseBStar(s string) float64 {
	s = strings.TrimSpace(s)
	if len(s) < 3 {
		return 0
	}

	// Exponent is the last two characters: sign and digit ("-3", "+0").
	expStr := s[len(s)-2:]
	mantStr := s[:len(s)-2]

	sign := 1.0
	if strings.HasPrefix(mantStr, "-") {
		sign = -1
		mantStr = mantStr[1:]
	} else if strings.HasPrefix(mantStr, "+") {
		mantStr = mantStr[1:]
	}

	// Implied decimal after the first digit: "10196" → 0.10196
	if mantStr == "" {
		return 0
	}
	f, err := strconv.ParseFloat("0."+mantStr, 64)
	if err != nil {
		return 0
	}

	e, err := strconv.Atoi(expStr)
	if err != nil {
		return 0
	}

	return sign * f * math.Pow(10, float64(e))
}
 
func copyTo(src, dst interface{}) error {
	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}
	dstVal := reflect.ValueOf(dst)
	if dstVal.Kind() != reflect.Ptr {
		return fmt.Errorf("dst must be a pointer")
	}
	dstVal = dstVal.Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		sf := srcVal.Type().Field(i)
		df := dstVal.FieldByName(sf.Name)
		if !df.IsValid() || !df.CanSet() {
			continue
		}

		// Special-case: BStar string → float64
		if sf.Name == "BStar" {
			df.Set(reflect.ValueOf(parseBStar(srcVal.Field(i).String())))
			continue
		}

		if sf.Type.AssignableTo(df.Type()) {
			df.Set(srcVal.Field(i))
		} else if sf.Type.ConvertibleTo(df.Type()) {
			df.Set(srcVal.Field(i).Convert(df.Type()))
		}
	}
	return nil
}   

func checksumOK(line string) bool {
    if len(line) < 69 {
        return false
    }
    sum := 0
    for _, c := range line[:68] {
        switch {
        case c >= '0' && c <= '9':
            sum += int(c - '0')
        case c == '-':
            sum += 1
        }
    }
    return int(line[68]-'0') == sum%10
}

func ParseTLE(lines []string) (*routing.TLE, error) {
  if len(lines) != 2 {
	return nil, fmt.Errorf("expected exactly 2 lines, got %d", len(lines))
  }
  for i, l := range lines {
  	if len(l) < 69 {
		return nil, fmt.Errorf("invalid TLE line %d length: %d", i+1, len(l))
  	}
	if !checksumOK(l) {
		return nil, fmt.Errorf("invalid TLE line %d checksum.", i+1)
  	}
  }
  t := &routing.TLE{}

  l1 := line1{}
  if err := fixedwidth.Unmarshal([]byte(lines[0]), &l1); err != nil {
	return nil, fmt.Errorf("unmarshal line 1: %w", err)
  }

  l2 := line2{}
  if err := fixedwidth.Unmarshal([]byte(lines[1]), &l2); err != nil {
	return nil, fmt.Errorf("unmarshal line 2: %w", err)
  } 
  if l1.SatelliteNum != l2.SatelliteNum {
	return nil, fmt.Errorf("catalog number mismatch: line1=%d line2=%d", l1.SatelliteNum, l2.SatelliteNum)
  }
  if err := copyTo(l1, t); err != nil {
	return nil, err
  }
  if err := copyTo(l2, t); err != nil {
	return nil, err
  } 
  // Eccentricity in TLE is stored without a leading "0.", e.g. "0008217" → 0.0008217
  if ec, err := strconv.ParseFloat(l2.Eccentricity, 64); err == nil {
	t.Eccentricity = ec / 1e7
  }
  return t, nil
}
