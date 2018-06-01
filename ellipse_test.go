package main

import (
	"fmt"
	"testing"
)

func TestKummerGauss(t *testing.T) {
	hVals := []float64{0, 0.1, 0.2, 1, 2, 20}
	for _, h := range hVals {
		fmt.Printf("KummerGauss(%f) = %15.12f\n", h, KummerGauss(h))
	}
}

func TestEllipsePerimeter(t *testing.T) {
	bVals := []float64{10, 5, 3, 1, 0}
	for _, b := range bVals {
		fmt.Printf("EllipsePerimeter(10,%f)=%12.9f\n", b, EllipsePerimeter(10, b))
	}
}
