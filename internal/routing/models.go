package routing

import (
  "github.com/TylerWerne40/sat-track/pb/api"
)


type TLE_Package struct {
  Name            string  // Option title line
  Line1           string  // Raw string for line 1
  Line2           string  // Raw string for Line 2
  TwoLines        TLE     // Two Line Element Struct
}

type TLE struct {
    
    // Parsed Fields
  SatelliteNum   int     `fixed:"3,7"`
	Classification string  `fixed:"8,8"`
	EpochYear      int     `fixed:"19,20"`
	EpochDays      float64 `fixed:"21,32"`
	BStar          float64 `fixed:"54,61"` // Note: TLE BStar has an implied decimal

	// Line 2 Fields (Standard TLE columns)
	// If parsing Line 2, these are the positions:
	Inclination    float64 `fixed:"9,16"`
	RightAscension float64 `fixed:"18,25"`
	Eccentricity   float64 `fixed:"27,33"` // Note: TLE Eccentricity has an implied leading decimal
	ArgPerigee     float64 `fixed:"35,42"`
	MeanAnomaly    float64 `fixed:"44,51"`
	MeanMotion     float64 `fixed:"53,63"` 
}

func tlePackageToProto(p TLE_Package) *pb.TLEPackage {
    return &pb.TLEPackage{
        Name:   p.Name,
        Line1:  p.Line1,
        Line2:  p.Line2,
        Parsed: &pb.ParsedTLE{
            SatelliteNum:    int32(p.TwoLines.SatelliteNum),
            Classification:  p.TwoLines.Classification,
            EpochYear:       int32(p.TwoLines.EpochYear),
            EpochDays:       p.TwoLines.EpochDays,
            Bstar:           p.TwoLines.BStar,
            Inclination:     p.TwoLines.Inclination,
            RightAscension:  p.TwoLines.RightAscension,
            Eccentricity:    p.TwoLines.Eccentricity,
            ArgPerigee:      p.TwoLines.ArgPerigee,
            MeanAnomaly:     p.TwoLines.MeanAnomaly,
            MeanMotion:      p.TwoLines.MeanMotion,
        },
    }
}
