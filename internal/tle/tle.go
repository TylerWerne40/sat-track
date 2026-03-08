package tle

import (
	"errors"
	"fmt"
	"github.com/TylerWerne40/sat-track/internal/routing"
	"github.com/ianlopshire/go-fixedwidth"
)


func ParseLine(t *routing.TLE, line string) (*routing.TLE, error) {
	if len(line) < 69 {
		return nil, errors.New("invalid TLE line length")
	}

  err := fixedwidth.Unmarshal([]byte(line), t)
  if err != nil {
    return nil, err
  }

  return t, nil
	
}

func ParseTLE(lines []string) (*routing.TLE, error) {
  if (len(lines) != 2) {
    return nil, fmt.Errorf("Not two lines")
  }
  t := &routing.TLE{}
  t, err := ParseLine(t, lines[0])
  if err != nil {
    return nil, err
  }
  t, err = ParseLine(t, lines[1])
  if err != nil {
    return nil, err
  }
  return t, nil
}
