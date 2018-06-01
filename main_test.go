package main

import (
	"fmt"
	"github.com/lvdlvd/go-geo-wgs84"
	"path/filepath"
	"runtime"
	"testing"
)

func testNormalizeTriangle(triangle []Coord, answer []string, t *testing.T) {
	if len(triangle) != 3 {
		t.Errorf("unexpected len %d on triangle input", len(answer))
	}
	if len(answer) != 3 && len(answer) != 0 {
		t.Errorf("unexpected len %d on answer input", len(answer))
	}
	_, file, line, _ := runtime.Caller(1)
	file = filepath.Base(file)
	for permut := 0; permut < 6; permut++ {
		pdebug := fmt.Sprintf("on permut %d from %s:%d", permut, file, line)
		idx1 := permut & 1
		idx2 := idx1 + 1
		tmp := make([]Coord, 0, 3)
		for i := 0; i < 3; i++ {
			if i == idx1 {
				tmp = append(tmp, triangle[idx2])
			} else if i == idx2 {
				tmp = append(tmp, triangle[idx1])
			} else {
				tmp = append(tmp, triangle[i])
			}
		}
		test, err := normalizeTriangle(tmp)
		if err == nil {
			if len(answer) != 0 {
				if len(test) != 3 {
					t.Errorf("unexpected len %d %s", len(test), pdebug)
				}
				found := []string{test[0].name, test[1].name, test[2].name}
				if found[0] != answer[0] || found[1] != answer[1] || found[2] != answer[2] {
					t.Errorf("bad normalize %s\n\tFound: %v\n\tExpected: %v", pdebug, found, answer)
				}
			} else {
				t.Errorf("failure was expected %s", pdebug)
			}
		} else if len(answer) != 0 {
			t.Errorf("err %v %s", err, pdebug)
		}
		triangle = tmp
	}
}

func TestNormalizeTriangle(t *testing.T) {
	testNormalizeTriangle(
		[]Coord{
			{dallas_lat, dallas_lon, "Dallas"},
			{houston_lat, houston_lon, "Houston"},
			{san_antonio_lat, san_antonio_lon, "San Antonio"},
		},
		[]string{
			"Dallas",
			"Houston",
			"San Antonio",
		}, t)
	testNormalizeTriangle(
		[]Coord{
			{48.85814833, 2.29528311, "Eiffel Tower"},
			{40.68934602, -74.0435975, "Statue of Liberty"},
			{37.82446289, -122.40172253, "Golden Gate Bridge"},
		},
		[]string{
			"Eiffel Tower",
			"Statue of Liberty",
			"Golden Gate Bridge",
		}, t)
	testNormalizeTriangle(
		[]Coord{
			{39.75570774, -104.86789455, "Denver"},
			{37.7577, -122.4376, "San Francisco"},
			{47.61302845, -122.3420645, "Seattle"},
		},
		[]string{
			"Seattle",
			"Denver",
			"San Francisco",
		}, t)
	testNormalizeTriangle([]Coord{{0, 0, "O"}, {0, 90, "OE"}, {-90, 0, "S"}}, []string{"O", "OE", "S"}, t)
	testNormalizeTriangle([]Coord{{90, 0, "N"}, {0, 90, "OE"}, {-90, 0, "S"}}, []string{}, t)
	testNormalizeTriangle([]Coord{{0, 0, "O"}, {0, 90, "OE"}, {90, 0, "N"}}, []string{"N", "OE", "O"}, t)
	testNormalizeTriangle([]Coord{{10, 10, "A"}, {10, 10, "A2"}, {20, 20, "B"}}, []string{}, t)
	testNormalizeTriangle([]Coord{{10, 10, "A"}, {10, 10, "A2"}, {20, 20, "B"}}, []string{}, t)
	_, azi, _ := wgs84.Inverse(rad(dallas_lat), rad(dallas_lon), rad(san_antonio_lat), rad(san_antonio_lon))
	colinear := Coord{0, 0, "colinear"}
	colinear.lat, colinear.lon, _ = wgs84.Forward(rad(dallas_lat), rad(dallas_lon), azi, 1000000.0)
	colinear.lat = deg(colinear.lat)
	colinear.lon = deg(colinear.lon)
	testNormalizeTriangle(
		[]Coord{
			{dallas_lat, dallas_lon, "Dallas"},
			{san_antonio_lat, san_antonio_lon, "San Antonio"},
			colinear,
		},
		[]string{}, t)
	testNormalizeTriangle([]Coord{{10, 10, "A"}, {-30, 30, "B"}, {80, -160, "C"}}, []string{"C", "B", "A"}, t)
}
