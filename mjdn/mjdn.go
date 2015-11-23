/**
 * Astronomical calculation for Golang.
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/deed.ja
 */

// Package mjdn は修正ユリウス日（Modified Julian Day）の計算を行います。
package mjdn

import (
	"fmt"
	"time"
)

// StartGregorian はグレゴリオ暦開始年を指定します。
// 既定は1582年10月15日です。
var StartGregorian = time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC)

// MJDN は修正ユリウス日（Modified Julian Day）の型を定義します。
type MJDN int64

// DayNumber は日付から修正ユリウス通日を取得します。
//
// 時分秒は無視します。
// グレゴリオ暦開始年より前のユリウス暦では dnJulian() を使って計算します。
// 1970年1月1日より前のグレゴリオ暦では dnGregorian() を使って計算します。
// 1970年1月1日以降は UNIX Time を用いて通日を取得します。
// 紀元前1年は0年として計算します（BC 1 year => 0 year）。
func DayNumber(t time.Time) MJDN {
	if t.Sub(time.Unix(0, 0)) >= 0 {
		return dnUnix(t)
	} else if t.Sub(StartGregorian) >= 0 {
		return dnGregorian(t)
	}
	return dnJulian(t)
}

// ToTime は修正ユリウス通日から日付（time.Time 形式）を取得します。
//
// 時刻（時分秒）は 00:00:00  でセットします。
// グレゴリオ暦開始年より前はユリウス暦とみなして計算します。
func (mjdn MJDN) ToTime(loc *time.Location) time.Time {
	if mjdn < dnGregorian(StartGregorian) {
		return mjdn.toJulian(loc)
	}
	return mjdn.toGregorian(loc)
}

// String は MJDN を文字列に展開します。
func (mjdn MJDN) String() string {
	return fmt.Sprintf("%v (%s)", int64(mjdn), mjdn.ToTime(time.UTC).Format("2006-01-02"))
}

// dnGregorian は Fliegel の公式を使い，グレゴリオ暦日から修正ユリウス通日を計算します。
//
// 時分秒は無視します。
// 計算を端折っているため紀元前の日付では正しく計算できません。
func dnGregorian(t time.Time) MJDN {
	y := int64(t.Year())
	m := int64(t.Month())
	if m < 3 {
		y--
		m += 9
	} else {
		m -= 3
	}
	d := int64(t.Day()) - 1
	return MJDN((1461*y)/4 + y/400 - y/100 + (153*m+2)/5 + d - 678881)
}

// toGregorian は修正ユリウス通日からグレゴリオ暦日を計算します。
//
// 時刻（時分秒）は 00:00:00  でセットします。
func (mjdn MJDN) toGregorian(loc *time.Location) time.Time {
	n := int64(mjdn) + 678881
	nn := fracByFloor(4*(n+1), 146097)
	a := 4*n + 3 + 4*fracByFloor(3*(nn+1), 4)
	b := 5*fracByFloor(modInt(a, 1461), 4) + 2
	y := fracByFloor(a, 1461)
	m := fracByFloor(b, 153) + 3
	if m > 12 {
		y++
		m -= 12
	}
	d := fracByFloor(modInt(b, 153), 5)
	return time.Date(int(y), time.Month(m), int(d+1), 0, 0, 0, 0, loc)
}

// dnJulian は Fliegel の公式を使い，ユリウス暦日から修正ユリウス通日を計算します。
//
// 時分秒は無視します。
// 紀元前1年は0年として計算します（BC 1 year => 0 year）。
func dnJulian(t time.Time) MJDN {
	y := int64(t.Year())
	m := int64(t.Month())
	if m < 3 {
		y--
		m += 9
	} else {
		m -= 3
	}
	d := int64(t.Day()) - 1
	return MJDN(fracByFloor(1461*y, 4) + (153*m+2)/5 + d - 678883)
}

// toJulian は修正ユリウス通日からユリウス暦日を計算します。
//
// 時刻（時分秒）は 00:00:00  でセットします。
func (mjdn MJDN) toJulian(loc *time.Location) time.Time {
	a := 4*(int64(mjdn)+678883) + 3
	b := 5*fracByFloor(modInt(a, 1461), 4) + 2
	y := fracByFloor(a, 1461)
	m := fracByFloor(b, 153) + 3
	if m > 12 {
		y++
		m -= 12
	}
	d := fracByFloor(modInt(b, 153), 5)
	return time.Date(int(y), time.Month(m), int(d+1), 0, 0, 0, 0, loc)
}

// dnUnix は UNIX Time で1970年1月1日からの通日を取得し，修正ユリウス通日を計算します。
//
// 時分秒は無視します。
// 計算を端折っているため1970年1月1日以前の日付では正しく計算できません。
func dnUnix(t time.Time) MJDN {
	const (
		onday   = int64(86400) //seconds
		baseDay = int64(40587) //Modified Julian Date at January 1, 1970
	)
	return MJDN((t.Unix())/onday + baseDay)
}

// fracByFloor は整数同士の除算結果を床関数で返します。
func fracByFloor(child, mother int64) int64 {
	delta := int64(0)
	if child < 0 && mother > 0 {
		delta = -1
	} else if child > 0 && mother < 0 {
		delta = -1
	}
	return child/mother + delta
}

// modInt は整数同士の剰余（modulo）を返します。

// 分母は絶対値をとります。かつ余りは必ず正の値になるようにします。
func modInt(child, mother int64) int64 {
	if mother < 0 {
		mother = 0 - mother
	}
	mod := child % mother
	if mod < 0 {
		return mother + mod
	}
	return mod
}
