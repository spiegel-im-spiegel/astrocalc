package modjulian

import "time"

func DayNumber(t time.Time) int64 {
	if t.Sub(time.Unix(0, 0)) >= 0 {
		return dnUnix(t)
	} else {
		return dnGregorian(t)
	}
}

func dnGregorian(t time.Time) int64 {
	y := int64(t.Year())
	m := int64(t.Month())
	if m < 3 {
		y -= 1
		m += 9
	} else {
		m -= 3
	}
	d := int64(t.Day()) - 1
	return (1461*y)/4 + y/400 - y/100 + (153*m+2)/5 + d - 678882
}

func dnUnix(t time.Time) int64 {
	const (
		onday   = int64(86400) //seconds
		baseDay = int64(40587) //Modified Julian Date at January 1, 1970
	)
	return (t.Unix())/onday + baseDay
}
