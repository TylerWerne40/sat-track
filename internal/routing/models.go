package routing




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
