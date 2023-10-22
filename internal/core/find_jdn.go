package core

import (
	log "github.com/sirupsen/logrus"
	"math"
)

// FindJDN find the date corresponding to astronomical position of body ipl
// related to delta degrees offset back from date tjdUt.
// Once it finds a date where the astronomical position is close to the target,
// it returns that date in Julian Day Number format.
func FindJDN(tjdUt float64, ipl int, delta float64) float64 {
	xx := make([]float64, 1)

	if delta > 360 || delta < 0 {
		return 0
	}

	err := CalcUt(tjdUt, ipl, xx)
	if err != nil {
		log.Println(err)
	}
	targetDegrees := xx[0] - delta
	if targetDegrees < 0 {
		targetDegrees += 360
	}

	currentJDN := tjdUt - (360/365.25)*delta
	step := delta
	tolerance := 0.000000001

	for math.Abs(step) > tolerance {
		err := CalcUt(currentJDN, ipl, xx)
		if err != nil {
			log.Println(err)
		}
		diff := xx[0] - targetDegrees
		if math.Abs(diff) < tolerance {
			return currentJDN
		}

		if (diff > 180) || (diff < 0 && math.Abs(diff) < 180) {
			currentJDN += step
		} else {
			currentJDN -= step
			step /= 2
		}
	}
	return currentJDN
}
