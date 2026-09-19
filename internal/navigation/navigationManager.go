package navigation

import (
	"math"
)

const WORLD_DIAMETER = 10_254

type NavManager struct {
	NavPoints      []MappedObject
	Airfields      []MappedObject
	UseHaversine   bool
	SelectedObject int
	SelectedType   ObjectType
}

// Returns x and y distance from coordinates
func (n NavManager) coordToDistanceHaversine(aLat, aLong, bLat, bLong float64, heading float64) (float64, float64, float64, float64) {
	phi1 := DegToRad(aLat)
	phi2 := DegToRad(bLat)
	lambda1 := DegToRad(aLong)
	lambda2 := DegToRad(bLong)
	deltaPhi := phi2 - phi1
	deltaLambda := lambda2 - lambda1

	theta := math.Pi - Archaversine(Haversine(deltaPhi)+math.Cos(phi1)*math.Cos(phi2)*Haversine(deltaLambda))
	// atan2(sin(lon2 - lon1) * cos(lat2), cos(lat1) * sin(lat2) - sin(lat1) * cos(lat2) * cos(lon2 - lon1))
	bearing := math.Atan2(
		math.Sin(deltaLambda)*math.Cos(phi2),
		math.Cos(phi1)*math.Sin(phi2)-math.Sin(phi1)*math.Cos(phi2)*math.Cos(deltaLambda),
	)

	distanceKm := theta * WORLD_DIAMETER

	distanceY := -math.Cos(bearing-DegToRad(heading)) * distanceKm
	distanceX := math.Sin(bearing-DegToRad(heading)) * distanceKm

	return distanceX, distanceY, distanceKm, bearing
}

func (n NavManager) coordToDistanceTrigonometry(aLat, aLong, bLat, bLong float64, heading float64) (float64, float64, float64, float64) {
	phi1 := DegToRad(aLat)
	phi2 := DegToRad(bLat)
	lambda1 := DegToRad(aLong)
	lambda2 := DegToRad(bLong)
	deltaPhi := phi2 - phi1
	deltaLambda := lambda2 - lambda1

	theta := math.Sqrt(math.Pow(deltaPhi, 2) + math.Pow(deltaLambda, 2))
	var bearing float64
	if theta == 0 {
		bearing = 0
	} else if math.Cos(phi1)*math.Sin(phi2)-math.Sin(phi1)*math.Cos(phi2)*math.Cos(deltaLambda) >= 0 {
		bearing = math.Asin(deltaLambda / theta)
	} else {
		bearing = math.Mod(math.Pi-math.Asin(deltaLambda/theta), 2*math.Pi)
	}

	distanceKm := theta * WORLD_DIAMETER

	distanceY := -math.Cos(bearing-DegToRad(heading)) * distanceKm
	distanceX := math.Sin(bearing-DegToRad(heading)) * distanceKm

	return distanceX, distanceY, distanceKm, bearing
}

func (n NavManager) CoordToDistance(aLat, aLong, bLat, bLong float64, heading float64) (float64, float64, float64, float64) {
	var x, y, dist, bearing float64
	if n.UseHaversine {
		x, y, dist, bearing = n.coordToDistanceHaversine(
			float64(aLat),
			float64(aLong),
			float64(bLat),
			float64(bLong),
			float64(heading),
		)
	} else {
		x, y, dist, bearing = n.coordToDistanceTrigonometry(
			float64(aLat),
			float64(aLong),
			float64(bLat),
			float64(bLong),
			float64(heading),
		)
	}
	return x, y, dist, bearing
}

func RadToDeg(a float64) float64 {
	return a * 180 / math.Pi
}

func DegToRad(a float64) float64 {
	return a / 180 * math.Pi
}

func Archaversine(a float64) float64 {
	return 2 * math.Acos(math.Sqrt(a))
}

func Haversine(a float64) float64 {
	return (1 - math.Cos(a)) / 2
}
