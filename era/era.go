/**
 * Astronomical calculation for Golang.
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/deed.ja
 */

// Package era は年号に関する処理を定義します。
package era

import (
	"fmt"
	"time"

	"github.com/spiegel-im-spiegel/astrocalc/mjdn"
)

// Era は年号に関する情報を定義します。
type Era struct {
	// 年号
	name, cname string

	// 年号の開始時点を起点とした年数（1, 2, 3, ...）
	year int64

	// 元の日付
	org *time.Time

	// 年号の開始時点（修正ユリウス通日）
	start mjdn.MJDN
}

// New は Era インスタンスを作成します。
func New(name string, cname string, start mjdn.MJDN) *Era {
	return &Era{name: name, cname: cname, start: start, year: 1}
}

// Get は年号＋年の情報を取得します。
func Get(era *Era, year int64) *Era {
	copy := era.Copy()
	copy.org = nil
	copy.year = year
	return copy
}

// Copy は Era インスタンスのコピーを作成します。
func (era *Era) Copy() *Era {
	return &Era{name: era.name, cname: era.cname, start: era.start, year: era.year, org: era.org}
}

// Check は指定した日付が年号の開始時点以降かどうかチェックします。
//
// 開始時点以降なら true
// 開始時点より前なら false
func (era *Era) Check(t time.Time) bool {
	return mjdn.DayNumber(t) >= era.start
}

// ToTime は Era インスタンスを time.Time に変換します。
func (era *Era) ToTime() time.Time {
	if era.org != nil {
		return *era.org
	}
	start := era.start.ToTime(time.UTC)
	if era.Year() == 1 {
		return start
	}
	return time.Date(start.Year()+int(era.Year())-1, 1, 1, 0, 0, 0, 0, time.UTC)
}

// ToEra は指定した日付を年号に基づく値に変換します。
func (era *Era) ToEra(t time.Time) *Era {
	copy := era.Copy()
	copy.org = &t
	start := copy.start.ToTime(t.Location())
	copy.year = int64(t.Year()) - int64(start.Year()) + 1
	return copy
}

// String は年号を出力します。
func (era *Era) String() string {
	if era.cname == "" {
		return fmt.Sprintf("%s %d", era.Name(), era.Year())
	}
	return fmt.Sprintf("%s (%s) %d", era.CName(), era.Name(), era.Year())
}

// Name は年号を出力します。
func (era *Era) Name() string {
	return era.name
}

// CName は年号を和名で出力します。
func (era *Era) CName() string {
	return era.cname
}

// Year は号の開始時点を起点とした年数を出力します。
func (era *Era) Year() int64 {
	return era.year
}
