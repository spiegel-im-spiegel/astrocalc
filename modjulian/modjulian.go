/**
 * Astronomical calculation for Golang.
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/deed.ja
 */

// modjulian パッケージは
// 修正ユリウス日（Modified Julian Date）の計算を行います。
package modjulian

import "time"

// DayNumber は
// 日付から修正ユリウス通日を取得します。
//
// 時刻（時分秒）は無視します。
// 1970年1月1日以前のグレゴリオ暦では Fliegel の公式を使って計算します。
// 1970年1月1日以降は UNIX Time を用いて通日を取得します。
func DayNumber(t time.Time) int64 {
	if t.Sub(time.Unix(0, 0)) >= 0 {
		return dnUnix(t)
	} else {
		return dnGregorian(t)
	}
}

// dnGregorian は
// Fliegel の公式を使い，日付から修正ユリウス通日を計算します。
//
// 時刻（時分秒）は無視します。
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
	return (1461*y)/4 + y/400 - y/100 + (153*m+2)/5 + d - 678881
}

// dnUnix は
// UNIX Time で1970年1月1日からの通日を取得し，修正ユリウス通日を計算します。
//
// 時刻（時分秒）は無視します。
// 1970年1月1日以前の日付では正しく計算できません。
func dnUnix(t time.Time) int64 {
	const (
		onday   = int64(86400) //seconds
		baseDay = int64(40587) //Modified Julian Date at January 1, 1970
	)
	return (t.Unix())/onday + baseDay
}
