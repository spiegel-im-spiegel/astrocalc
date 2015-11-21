/**
 * Astronomical calculation for Golang.
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/deed.ja
 */

// Package modjulian は修正ユリウス日（Modified Julian Date）の計算を行います。
package modjulian

import "time"

// StartGregorian がグレゴリオ暦開始年（既定は1582年10月15日）を指します
var StartGregorian = time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC)

// DayNumber は日付から修正ユリウス通日を取得します。
//
// 時刻（時分秒）は無視します。
// グレゴリオ暦開始年より前のユリウス暦では dnJulian() を使って計算します。
// 1970年1月1日より前のグレゴリオ暦では dnGregorian() を使って計算します。
// 1970年1月1日以降は UNIX Time を用いて通日を取得します。
func DayNumber(t time.Time) int64 {
	if t.Sub(time.Unix(0, 0)) >= 0 {
		return dnUnix(t)
	} else if t.Sub(StartGregorian) >= 0 {
		return dnGregorian(t)
	}
	return dnJulian(t)
}

// dnGregorian は Fliegel の公式を使い，グレゴリオ暦日から修正ユリウス通日を計算します。
//
// 時刻（時分秒）は無視します。
func dnGregorian(t time.Time) int64 {
	y := int64(t.Year())
	m := int64(t.Month())
	if m < 3 {
		y--
		m += 9
	} else {
		m -= 3
	}
	d := int64(t.Day()) - 1
	return (1461*y)/4 + y/400 - y/100 + (153*m+2)/5 + d - 678881
}

// dnJulian は Fliegel の公式を使い，ユリウス暦日から修正ユリウス通日を計算します。
//
// 時刻（時分秒）は無視します。
func dnJulian(t time.Time) int64 {
	y := int64(t.Year())
	m := int64(t.Month())
	if m < 3 {
		y--
		m += 9
	} else {
		m -= 3
	}
	//年がマイナスの場合は更に1を引く（床関数対策）
	dy := int64(0)
	if y < 0 {
		dy--
	}
	d := int64(t.Day()) - 1
	return (1461*y)/4 + dy + (153*m+2)/5 + d - 678883
}

// dnUnix は UNIX Time で1970年1月1日からの通日を取得し，修正ユリウス通日を計算します。
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
