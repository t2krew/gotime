package gotime

import (
	"math"
	"time"
)

func DiffDays(target, base time.Time) int {
	return diffDays(target, base, 0)
}

func diffDays(target, base time.Time, days int) int {
	// limits the largest representable duration to approximately 290 years.
	if target.Sub(base).Hours()/24 > 289*365 {
		var (
			subDate     = base.AddDate(289, 0, 0)
			subDays = int(math.Ceil(subDate.Sub(base).Hours() / 24))
		)
		return diffDays(target, subDate, subDays+days)
	} else {
		return int(math.Ceil(target.Sub(base).Hours()/24)) + days
	}
}