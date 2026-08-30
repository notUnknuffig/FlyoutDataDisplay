package stateNavigation

import (
	"math"
)

const WORLD_DIAMETER = 10_254

type NavManager struct {
	NavPoints []MappedObjects
	Airfields []MappedObjects
}

/*
// Returns: Dist X, Dist Y, Dist Diagonal in km
func distance(aLat, aLong, bLat, bLong float32) (float32, float32, float32) {
	distLat := bLat - aLat
	distLong := bLong - aLong
	dist := float32(math.Sqrt(math.Pow(float64(distLat), 2) + math.Pow(float64(distLong), 2)))
	return distLat * 10_254, distLong * 10_254, dist * 10_254
} */

// Returns x and y distance from coordinates
func coordToDistance(aLat, aLong, bLat, bLong float64, heading float64) (float64, float64, float64, float64) {
	phi1 := degToRad(aLat)
	phi2 := degToRad(bLat)
	lambda1 := degToRad(aLong)
	lambda2 := degToRad(bLong)
	deltaPhi := phi2 - phi1
	deltaLambda := lambda2 - lambda1

	theta := math.Pi - Archaversine(Haversine(deltaPhi)+math.Cos(phi1)*math.Cos(phi2)*Haversine(deltaLambda))
	bearing := math.Atan2(math.Sin(deltaLambda)*math.Cos(lambda2), math.Cos(phi1)*math.Sin(phi2)-math.Sin(phi1)*math.Cos(phi2)*math.Cos(deltaLambda))

	distanceKm := theta * WORLD_DIAMETER

	distanceX := math.Sin(bearing+degToRad(heading)) * distanceKm
	distanceY := math.Cos(bearing+degToRad(heading)) * distanceKm

	return distanceX, distanceY, distanceKm, bearing
}

func RadToDeg(a float64) float64 {
	return a * 180 / math.Pi
}

func degToRad(a float64) float64 {
	return a / 180 * math.Pi
}

func Archaversine(a float64) float64 {
	return 2 * math.Acos(math.Sqrt(a))
}

func Haversine(a float64) float64 {
	return (1 - math.Cos(a)) / 2
}
