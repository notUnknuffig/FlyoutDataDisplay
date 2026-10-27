package navigation

import (
	"math"
	"strconv"
)

func LoadCoords(degSize float64) []MappedObject {
	navPoints := []MappedObject{}
	for i := 0; i < int(math.Floor(360/degSize))+1; i++ {
		for j := 0; j < int(math.Floor(180/degSize))+1; j++ {
			navPoints = append(navPoints, MappedObject{
				Name:      strconv.FormatFloat(degSize*float64(j), 'f', 0, 64) + " - " + strconv.FormatFloat(degSize*float64(i), 'f', 0, 64),
				Latitude:  float32(degSize)*float32(j) - 90,
				Longitude: float32(degSize)*float32(i) - 180,
				Heading:   0,
				Altitude:  0,
				Allied:    true,
				Type:      NAV_POINT,
			})
		}
	}
	return navPoints
}
