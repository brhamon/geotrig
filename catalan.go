package main

import (
	"fmt"
	"math/big"
)

//A048881     a(n) = A000120(n+1) - 1
//    Highest power of 2 dividing n-th Catalan number
func A048881(n int64) int {
	return PopCount(n+1) - 1
}

//Catalan [A000108]     Catalan numbers: C(n) = binomial(2n,n)/(n+1) = (2n)!/(n!(n+1)!)
func Catalan(n int) *big.Int {
	z := new(big.Int)
	z.Binomial(int64(n)*2, int64(n))
	bigN := big.NewInt(int64(n) + 1)
	return z.Div(z, bigN)
}

//A002596   Numerators in expansion of sqrt(1+x)
//  a(n+2) = C(n+1)/2^k(n+1), n >= 0
//    where C(n) = A000108(n), k(n) = A048881(n)
func A002596(n int) *big.Int {
	if n < 2 {
		return big.NewInt(1)
	}
	n--
	c := Catalan(n)
	var denom big.Int
	denom.SetBit(&denom, A048881(int64(n)), 1)
	c.Div(c, &denom)
	if (n % 2) != 0 {
		c.Neg(c)
	}
	return c
}

//A056981     a(n) = A002596(n)^2
func A056981(n int) *big.Int {
	c := A002596(n)
	return c.Mul(c, c)
}

//A005187     a(n) = a(floor(n/2)) + n
func A005187(n int) int {
	if n <= 0 {
		return 0
	}
	return A005187(n/2) + n
}

//A056982 a(n) = 4^A005187(n)
func A056982(n int) *big.Int {
	c := big.NewInt(0)
	return c.SetBit(c, A005187(n)*2, 1)
}

// KummerGaussTerm returns the n-th term of the Kummer-Gauss series
func KummerGaussTerm(n int) *big.Rat {
	r := new(big.Rat)
	return r.SetFrac(A056981(n), A056982(n))
}

// RatKummerGauss provides an estimate of the perimeter of an ellipse with a given h-ratio
// It sums the first 100 terms of the Kummer-Gauss series.
func RatKummerGauss(h *big.Rat) *big.Rat {
	hProd := big.NewRat(1, 1)
	sum := big.NewRat(0, 1)
	for term := 0; term < 100; term++ {
		kg := KummerGaussTerm(term)
		if term != 0 {
			hProd.Mul(hProd, h)
		}
		kg.Mul(kg, hProd)
		sum.Add(sum, kg)
	}
	return sum
}

// EllipsePerimeterPrecise uses rationals to calculate the perimeter of an ellipse
// with radii a and b, then returns the result as a float64.
func EllipsePerimeterPrecise(a float64, b float64) float64 {
	var aRat big.Rat
	var bRat big.Rat

	aRat.SetFloat64(a)
	bRat.SetFloat64(b)
	result := RatEllipsePerimeter(&aRat, &bRat)
	res, _ := result.Float64()
	return res
}

// RatEllipsePerimeter retains the ellipe perimeter as a Rat
func RatEllipsePerimeter(aRat *big.Rat, bRat *big.Rat) *big.Rat {
	var tmp big.Rat
	var diff big.Rat
	var sum big.Rat
	var sum2 big.Rat

	diff.Sub(aRat, bRat)
	diff.Mul(&diff, &diff)
	sum.Add(aRat, bRat)
	sum2.Mul(&sum, &sum)
	// tmp becomes 'h'
	tmp.Quo(&diff, &sum2)
	kg := RatKummerGauss(&tmp)
	kg.Mul(kg, &sum)
	// tmp becomes Pi
	RatPi(&tmp)
	kg.Mul(kg, &tmp)
	fmt.Printf("EllipsePerimeterPrecise: %s\n", kg.FloatString(100))
	return kg
}

// CircumferencePrecise uses rationals to calculate the circumference of
// a circle with radius r, then returns the result as a float64
func CircumferencePrecise(r float64) float64 {
	var rRat big.Rat

	rRat.SetFloat64(r)
	result := RatCircumference(&rRat)
	res, _ := result.Float64()
	return res
}

// RatCircumference calculates the circumference of a circle with
// radius rRat
func RatCircumference(rRat *big.Rat) *big.Rat {
	var piRat big.Rat

	RatPi(&piRat)
	piRat.Mul(&piRat, rRat)
	piRat.Mul(&piRat, big.NewRat(2, 1))
	fmt.Printf("CircumferencePrecise: %s\n", piRat.FloatString(100))
	return &piRat
}

func calculateEarthParameters() {
	aRat := new(big.Rat)
	aRat.SetFloat64(WGS84A)
	ep := RatCircumference(aRat)
	bRat := new(big.Rat)
	bRat.SetFloat64(WGS84B)
	pp := RatEllipsePerimeter(aRat, bRat)
	factor := big.NewRat(1, 4)
	ep.Mul(ep, factor)
	pp.Mul(pp, factor)
	ef := new(big.Rat)
	ef.Quo(pp, ep)
	fmt.Printf("90-deg equatorial arc: %s\n", ep.FloatString(100))
	fmt.Printf("     90-deg polar arc: %s\n", pp.FloatString(100))
	err := new(big.Rat)
	err.Sub(ep, pp)
	fmt.Printf("                error: %s\n", err.FloatString(100))
	fmt.Printf("       scaling factor: %s\n", ef.FloatString(100))
}
