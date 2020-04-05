package gotime

import (
	"testing"
	"time"
)

func TestDiffDays(t *testing.T) {
	var base = time.Date(2020, 4, 5, 0, 0, 0, 0, time.Local)
	t.Run("after base", func(t *testing.T) {
		var (
			cases = []struct {
				Date time.Time
				Days int
			}{
				{
					Date: time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local),
					Days: 179,
				},
				{
					Date: time.Date(2040, 4, 5, 0, 0, 0, 0, time.Local),
					Days: 7305,
				},
				{
					Date: time.Date(2120, 1, 1, 0, 0, 0, 0, time.Local),
					Days: 36429,
				},
				{
					Date: time.Date(2320, 12, 31, 0, 0, 0, 0, time.Local),
					Days: 109842,
				},
				{
					Date: time.Date(2720, 12, 31, 0, 0, 0, 0, time.Local),
					Days: 255939,
				},
			}
		)
		for _, c := range cases {
			if diff := DiffDays(c.Date, base); diff != c.Days {
				t.Errorf("DiffDays('%s', '%s'), expect %d, but got %d", c.Date.Format("20060102"), base.Format("20060102"), c.Days, diff)
			}
		}
	})

	t.Run("before base", func(t *testing.T) {
		var (
			cases = []struct {
				Date time.Time
				Days int
			}{
				{
					Date: time.Date(2010, 10, 1, 0, 0, 0, 0, time.Local),
					Days: -3474,
				},
				{
					Date: time.Date(1800, 10, 1, 0, 0, 0, 0, time.Local),
					Days: -80175,
				},
			}
		)
		for _, c := range cases {
			if diff := DiffDays(c.Date, base); diff != c.Days {
				t.Errorf("DiffDays('%s', '%s'), expect %d, but got %d", c.Date.Format("20060102"), base.Format("20060102"), c.Days, diff)
			}
		}
	})

	t.Run("equal base", func(t *testing.T) {
		var (
			cases = []struct {
				Date time.Time
				Days int
			}{
				{
					Date: time.Date(2020, 4, 5, 0, 0, 0, 0, time.Local),
					Days: 0,
				},
			}
		)
		for _, c := range cases {
			if diff := DiffDays(c.Date, base); diff != c.Days {
				t.Errorf("DiffDays('%s', '%s'), expect %d, but got %d", c.Date.Format("20060102"), base.Format("20060102"), c.Days, diff)
			}
		}
	})
}
