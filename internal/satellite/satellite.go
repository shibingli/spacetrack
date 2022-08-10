package satellite

import (
	goSatellite "github.com/joshuaferrara/go-satellite"
	"time"
)

func ParseTLE(line1, line2 string) (sat goSatellite.Satellite) {
	return goSatellite.TLEToSat(line1, line2, goSatellite.GravityWGS84)
}

func Propagate(sat goSatellite.Satellite, t time.Time) (position, velocity goSatellite.Vector3) {
	return goSatellite.Propagate(sat, t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func TLEPropagate(line1, line2 string, t time.Time) (position, velocity goSatellite.Vector3) {
	tle := ParseTLE(line1, line2)
	return Propagate(tle, t)
}

func ECIToECEF(eciCoords goSatellite.Vector3, t time.Time) (ecfCoords goSatellite.Vector3) {
	return goSatellite.ECIToECEF(eciCoords, GSTimeFromDate(t))
}

func ECIToLLA(eciCoords goSatellite.Vector3, t time.Time) (altitude, velocity float64, ret goSatellite.LatLong) {
	return goSatellite.ECIToLLA(eciCoords, GSTimeFromDate(t))
}

func LatLongDeg(rat goSatellite.LatLong) (deg goSatellite.LatLong) {
	return goSatellite.LatLongDeg(rat)
}

type LLA struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

func LatLonAlt(line1, line2 string, t time.Time) (deg *LLA) {

	pos, _ := TLEPropagate(line1, line2, t)

	alt, _, ret := ECIToLLA(pos, t)

	latLong := LatLongDeg(ret)

	lla := &LLA{
		Latitude:  latLong.Latitude,
		Longitude: latLong.Longitude,
		Altitude:  alt,
	}

	return lla
}

func GSTimeFromDate(t time.Time) float64 {
	return goSatellite.GSTimeFromDate(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
}
