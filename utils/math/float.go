package mathutils

import (
	"math"
)

// Rounds U p
//
// Exemples:
//
//	12.3416 , 2 -> 12.35
//	12.3456 , 2 -> 12.35
//
// Parameters:
//   - val(float64): the value to round
//   - procision(int): the rouding precision
//
// Return:
//   - the rounded value (float64)
func RoundUp(val float64, precision int) float64 {
	return math.Ceil(val*(math.Pow10(precision))) / math.Pow10(precision)
}

// Rounds Down
//
// Exemples:
//
//	12.3416 , 2 -> 12.34
//	12.3456 , 2 -> 12.34
//
// Parameters:
//   - val(float64): the value to round
//   - procision(int): the rouding precision
//
// Return:
//   - the rounded value (float64)
func RoundDown(val float64, precision int) float64 {
	return math.Floor(val*(math.Pow10(precision))) / math.Pow10(precision)
}

// Rounds to the nearest
// Exemples:
//
//	12.3416 , 2 -> 12.34
//	12.3456 , 2 -> 12.35
//
// Parameters:
//   - val(float64): the value to round
//   - procision(int): the rouding precision
//
// Return:
//   - the rounded value (float64)
func Round(val float64, precision int) float64 {
	return math.Round(val*(math.Pow10(precision))) / math.Pow10(precision)
}
