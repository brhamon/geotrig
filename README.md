# geotrig

Calculates the circumcenters of triangles on the WGS-84 ellipsoid.

Includes some experiments using Go's big.Rat package to avoid introducing 
float64 rounding errors, which tend to accumulate in iterative calculations.

This is raw code, poorly organized and not entirely well thought-out; however,
it has proven to be useful to explore mathematics with Go. I hope you find it
illuminating.
