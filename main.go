package main

import (
	"errors"
	"fmt"
	"github.com/lvdlvd/go-geo-wgs84"
	"math"
	"strings"
)

const (
	dallas_lat      float64 = 32.776664
	dallas_lon      float64 = -96.796988
	houston_lat     float64 = 29.760427
	houston_lon     float64 = -95.369803
	san_antonio_lat float64 = 29.424122
	san_antonio_lon float64 = -98.493628

	// Result from computation in ellipse.go: eliminates overshoot.
	compensationFactor float64 = 0.9983242984277140011446146599937898342419651738278666798238725123410681830530906421805472817806713887
)

var (
	MaxIters int = 0
)

func Rad(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}

func Deg(rad float64) float64 {
	return rad * (180.0 / math.Pi)
}

type Coord struct {
	Lat  float64
	Lon  float64
	Name string
}

// Tests if an azimuth is the range [-Pi, Pi)
func IsAziNormal(azi float64) bool {
	return azi >= -math.Pi && azi <= math.Pi
}

// Tests if the Coord is normalized.
// A normal coordinate's lat is in the range [-90, 90],
// and its lon is in the range [-180, 180]
func (c *Coord) IsNormal() bool {
	return c.Lat <= 90.0 && c.Lat >= -90.0 && c.Lon <= 180.0 && c.Lon >= -180.0
}

// Tests if the Coord is on either pole
func (c *Coord) IsPolar() bool {
	return math.Abs(c.Lat) == 90.0
}

// Tests if this Coord is identical to a second
// All pairs of non-Equal, non-antipodal points have exactly one geodesic
func (lhs *Coord) Equal(rhs *Coord) bool {
	return lhs.Lat == rhs.Lat && lhs.Lon == rhs.Lon
}

// Tests if this Coord is antipodal with a second Coord.
// All pairs of non-Equal, non-antipodal points have exactly one geodesic
func (lhs *Coord) IsAntipodal(rhs *Coord) bool {
	if lhs.Lat+rhs.Lat == 0.0 {
		if math.Abs(lhs.Lat) == 90.0 {
			return true
		}
		return math.Abs(lhs.Lon-rhs.Lon) == 180.0
	}
	return false
}

func (this *Coord) Antipode() (antipode *Coord) {
	antipode.Lat = -this.Lat
	if this.Lon < 0.0 {
		antipode.Lon = this.Lon + 180.0
	} else {
		antipode.Lon = this.Lon - 180.0
	}
	return
}

// Rotate an angle 'azi' by 'angle' radians.
// azi and angle are in the range [-Pi, +Pi)
func Rotate(azi float64, angle float64) float64 {
	if !IsAziNormal(azi) {
		panic(fmt.Sprintf("abnormal azimuth %f", azi))
	}
	if !IsAziNormal(angle) {
		panic(fmt.Sprintf("abnormal angle %f", angle))
	}
	res := azi + angle
	if res >= math.Pi {
		res -= 2 * math.Pi
	} else if res < -math.Pi {
		res += 2 * math.Pi
	}
	return res
}

// Normalizes a triangle, expressed as an array of exactly three Coords,
// such that they are returned in clockwise order, as viewed from above the
// centroid. Detects and returns an error for degenerate conditions, such as:
// * two antipodal Coords,
// * three Coords on the same geodesic, or
// * two equal Coords.
// The northern-most coord is returned first. If the two northern-most Coords
// have the same latitude, the western-most of the two is returned first.
func normalizeTriangle(in []Coord) (out []Coord, err error) {
	if len(in) != 3 {
		return in, errors.New("not a triangle: need 3 coordinates")
	}
	if !in[0].IsNormal() || !in[1].IsNormal() || !in[2].IsNormal() {
		return in, errors.New("bad input: coordinates not normal")
	}
	if in[0].Equal(&in[1]) || in[0].Equal(&in[2]) || in[1].Equal(&in[2]) {
		return in, errors.New("not a triangle: coordinates equal")
	}
	if in[0].IsAntipodal(&in[1]) || in[0].IsAntipodal(&in[2]) || in[1].IsAntipodal(&in[2]) {
		return in, errors.New("invalid triangle: coordinates antipodal")
	}

	// To test if Coords are non-co-geodesic, pick any non-polar Coord
	// and compare the azimuths to the other two Coords.
	var i int
	for i = 0; i < 3; i++ {
		if in[i].IsPolar() {
			continue
		}
		break
	}
	faz := math.MaxFloat64
	for j := 0; j < 3; j++ {
		if i == j {
			continue
		}
		_, tmpAzi, _ := wgs84.Inverse(Rad(in[i].Lat), Rad(in[i].Lon), Rad(in[j].Lat), Rad(in[j].Lon))
		if tmpAzi < 0 {
			tmpAzi += math.Pi
		}
		if math.Abs(tmpAzi-faz) < 1.0e-15 {
			return in, errors.New("invalid triangle: all points on same geodesic")
		}
		faz = tmpAzi
	}

	// Select the northernmost point and call it 'A'
	max := -math.MaxFloat64
	idxA := -1
	for i = 0; i < 3; i++ {
		if in[i].Lat > max {
			max = in[i].Lat
			idxA = i
		}
	}
	if idxA < 0 {
		return in, errors.New("logic error")
	}
	out = make([]Coord, 0, 3)
	out = append(out, in[idxA])
	idxB := -1
	if in[idxA].Lat == 90.0 {
		// special case: 'A' is North Pole
		eastMost := -math.MaxFloat64
		for i = 0; i < 3; i++ {
			if i == idxA {
				continue
			}
			if in[i].Lon > eastMost {
				eastMost = in[i].Lon
				idxB = i
			}
		}
	} else {
		// do either of the other two points have the same latitude?
		westMost := in[idxA].Lon
		for i = 0; i < 3; i++ {
			if i == idxA {
				continue
			}
			if in[i].Lat == in[idxA].Lat {
				idxB = i
				if in[i].Lon < westMost {
					// find the western-most
					westMost = in[i].Lon
				}
			}
		}
		if idxB != -1 && westMost != in[idxA].Lon {
			// switch to western-most
			idxA = idxB
			out[0] = in[idxA]
		}
		idxB = -1
		leastCw := math.MaxFloat64
		az := []float64{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64}
		var tfaz float64
		for i = 0; i < 3; i++ {
			if i == idxA {
				continue
			}
			_, tfaz, _ = wgs84.Inverse(Rad(in[idxA].Lat), Rad(in[idxA].Lon), Rad(in[i].Lat), Rad(in[i].Lon))
			az[i] = tfaz
			// change [-Pi, Pi) -> [0, 2*Pi)
			if tfaz < 0 {
				tfaz += 2 * math.Pi
			}
			if tfaz < leastCw {
				leastCw = tfaz
				idxB = i
			}
		}
		idxC := -1
		for i = 0; i < 3; i++ {
			if i == idxA || i == idxB {
				continue
			}
			idxC = i
			break
		}
		if (az[idxB] > 0 && az[idxC] < 0 && az[idxB]-az[idxC] < math.Pi) ||
			(az[idxC] > 0 && az[idxB] < 0 && az[idxC]-az[idxB] < math.Pi) {
			// If the interior angle between the azimuths AB and AC
			// contains the North pole, then flip B and C.
			tmp := idxB
			idxB = idxC
			idxC = tmp
		}
	}
	if idxB < 0 {
		return out, errors.New("logic error")
	}
	out = append(out, in[idxB])
	for i = 0; i < 3; i++ {
		if i == idxA || i == idxB {
			continue
		}
		break
	}
	if i >= 3 {
		return out, errors.New("logic error")
	}
	out = append(out, in[i])
	return out, nil
}

// Convert an azimuth to a compass bearing
// azimuth is in the range [-Pi, Pi)
// compass bearing is in the range [0, 360)
func azimuthToDeg(azi float64) float64 {
	if !IsAziNormal(azi) {
		panic(fmt.Sprintf("abnormal azimuth %f", azi))
	}
	if azi < 0 {
		azi += 2 * math.Pi
	}
	return Deg(azi)
}

func displayTriangle(t []Coord) {
	fmt.Println(fmt.Sprintf("{%f,%f} %s / {%f,%f} %s / {%f,%f} %s",
		t[0].Lat, t[0].Lon, t[0].Name,
		t[1].Lat, t[1].Lon, t[1].Name,
		t[2].Lat, t[2].Lon, t[2].Name))
}

func displayKmlPath(c []Coord, msg string) {
	for false {
		var b strings.Builder
		b.WriteString(fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
	<kml xmlns="http://www.opengis.net/kml/2.2" xmlns:gx="http://www.google.com/kml/ext/2.2" xmlns:kml="http://www.opengis.net/kml/2.2" xmlns:atom="http://www.w3.org/2005/Atom">
	<Document>
		<name>%s</name>
		<Placemark>
			<name>%s</name>
			<styleUrl>#m_ylw-pushpin</styleUrl>
			<LineString>
				<tessellate>1</tessellate>
				<coordinates>
					`, msg, msg))
		for _, ci := range c {
			b.WriteString(fmt.Sprintf("%f,%f,0 ", ci.Lon, ci.Lat))
		}
		b.WriteString(fmt.Sprintf("%f,%f,0\n", c[0].Lon, c[0].Lat))
		b.WriteString(`			</coordinates>
			</LineString>
		</Placemark>
	</Document>
	</kml>`)
		fmt.Println(b.String())
	}
}

func displayGeodesic(cLat float64, cLon float64, azi float64, msg string) {
	// one-third of Earth's circumference at the Equator
	step := 2 * math.Pi * wgs84.WGS84_a / 3
	aLat, aLon, _ := wgs84.Forward(Rad(cLat), Rad(cLon), azi, step)
	if azi <= 0.0 {
		azi += math.Pi
	} else {
		azi -= math.Pi
	}
	bLat, bLon, _ := wgs84.Forward(Rad(cLat), Rad(cLon), azi, step)
	arr := make([]Coord, 0, 3)
	arr = append(arr, Coord{Deg(bLat), Deg(bLon), ""})
	arr = append(arr, Coord{Deg(cLat), Deg(cLon), ""})
	arr = append(arr, Coord{Deg(aLat), Deg(aLon), ""})
	displayKmlPath(arr, msg)
}

// Estimate the distance to the intersection point X along the adjacent geodesic
// The base points are any disjoint points on non-coincident geodesics.
// The adjacent geodesic passes through the adjacent base point, in the
// direction of the adjacent azimuth.
// The opposite geodesic passes through the opposite base point, in the
// direction of the opposite azimuth.
// The base geodesic passes through both base points.
// 'distance' is the distance between base points in meters
// 'adjacent' is the angle between the base geodesic and the adjacent
//    geodesic, in radians. Input range: [-pi, pi)
// 'opposite' is the angle between the base geodesic and the opposite
//    geodesic, in radians. Input range: [-pi, pi)
// Returns an estimate of the distance along the adjacent geodesic to
//    point X, the nearest point where the adjacent and opposite geodesics
//    intersect, in radians.
// Formula is given by (3) and (4) from Sergio Baselga and Jose Martinez-Llario
// Stud. Geophys. Geod., 62 (2018), DOI: 10.1007/s11200-017-1020-z (in print)
// compensationFactor added to avoid "overshoot" by scaling the spherical distance
// estimate such that in the worst case (equator to pole), overshoot is eliminated.
func estimateDistToX(distance float64, adjacent float64, opposite float64) float64 {
	sVal, cVal := math.Sincos(distance / wgs84.WGS84_a)
	sa, ca := math.Sincos(adjacent)
	return wgs84.WGS84_a * math.Atan2(sVal, (cVal*ca)+((1.0/math.Tan(opposite))*sa)) * compensationFactor
}

func triangleCircumcenter(t []Coord) (coord Coord, err error) {
	accuracy := Rad(0.00001 / 3600)
	tri, err := normalizeTriangle(t)
	if err != nil {
		return
	}
	s, faz, _ := wgs84.Inverse(Rad(tri[0].Lat), Rad(tri[0].Lon), Rad(tri[1].Lat), Rad(tri[1].Lon))
	m01lat, m01lon, m10azi := wgs84.Forward(Rad(tri[0].Lat), Rad(tri[0].Lon), faz, s/2.0)
	// m10azi is azimuth from the midpoint facing back toward 0. Turn left 90 degrees.
	m10azi = Rotate(m10azi, -math.Pi/2.0)
	displayGeodesic(m01lat, m01lon, m10azi, fmt.Sprintf("%s-%s perpendicular midpoint",
		tri[0].Name, tri[1].Name))

	s, faz, _ = wgs84.Inverse(Rad(tri[1].Lat), Rad(tri[1].Lon), Rad(tri[2].Lat), Rad(tri[2].Lon))
	m12lat, m12lon, m21azi := wgs84.Forward(Rad(tri[1].Lat), Rad(tri[1].Lon), faz, s/2.0)
	// find the perpendicular geodesic
	m21azi = Rotate(m21azi, -math.Pi/2.0)
	displayGeodesic(m12lat, m12lon, m21azi, fmt.Sprintf("%s-%s perpendicular midpoint",
		tri[1].Name, tri[2].Name))

	for count := 0; true; count++ {
		if math.Abs(m12lat-m01lat) < accuracy && math.Abs(m12lon-m01lon) < accuracy {
			break
		}
		sAC, azac, azca := wgs84.Inverse(m12lat, m12lon, m01lat, m01lon)
		A := math.Abs(m21azi - azac)
		C := math.Abs(m10azi - azca)
		if count > MaxIters {
			MaxIters = count
		}
		if count > 10 {
			err = errors.New("endless")
			return
		}
		sAX := estimateDistToX(sAC, A, C)
		sCX := estimateDistToX(sAC, C, A)
		m12lat, m12lon, m21azi = wgs84.Forward(m12lat, m12lon, m21azi, sAX)
		m01lat, m01lon, m10azi = wgs84.Forward(m01lat, m01lon, m10azi, sCX)
		// wgs84.Forward returns the arriving azimuth rotated Pi radians
		// such that it faces the original coord. Invert that.
		m21azi = Rotate(m21azi, math.Pi)
		m10azi = Rotate(m10azi, math.Pi)
	}
	coord = Coord{Deg(m12lat), Deg(m12lon), fmt.Sprintf("circumcenter %s-%s-%s", t[0].Name, t[1].Name, t[2].Name)}
	return
}

func doTriangle(triangle []Coord) {
	displayTriangle(triangle)
	coord, err := triangleCircumcenter(triangle)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%s: %f, %f", coord.Name, coord.Lat, coord.Lon))
	displayKmlPath(triangle, fmt.Sprintf("%s-%s-%s", triangle[0].Name, triangle[1].Name, triangle[2].Name))

	for i := 0; i < len(triangle); i++ {
		s, _, _ := wgs84.Inverse(Rad(coord.Lat), Rad(coord.Lon), Rad(triangle[i].Lat), Rad(triangle[i].Lon))
		fmt.Println(fmt.Sprintf("cc-to-%s: %fm", triangle[i].Name, s))
	}
}

func main() {
	calculateEarthParameters()
	doTriangle(
		[]Coord{
			{dallas_lat, dallas_lon, "Dallas"},
			{houston_lat, houston_lon, "Houston"},
			{san_antonio_lat, san_antonio_lon, "San Antonio"},
		})
	doTriangle(
		[]Coord{
			{48.85814833, 2.29528311, "Eiffel Tower"},
			{40.68934602, -74.0435975, "Statue of Liberty"},
			{37.82446289, -122.40172253, "Golden Gate Bridge"},
		})
	doTriangle(
		[]Coord{
			{39.75570774, -104.86789455, "Denver"},
			{37.7577, -122.4376, "San Francisco"},
			{47.61302845, -122.3420645, "Seattle"},
		})
	fmt.Printf("MaxIters = %d\n", MaxIters)
}
