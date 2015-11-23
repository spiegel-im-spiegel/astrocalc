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

	"github.com/spiegel-im-spiegel/astrocalc/modjulian"
)

// 年号のインスタンス
var (
	AnnoDomini = New("A.D.", "西暦", -678577, false) // 0001-01-01
)

// Era は年号に関する情報を定義します。
type Era struct {
	// 年号
	name, wamei string

	// 年号の開始時点を起点とした年数（1, 2, 3, ...）
	year int64

	// 元の日付
	org *time.Time

	// 年号の開始時点（修正ユリウス通日）
	// 開始時点のチェックを行わない場合は checkBounds を false にセット
	start       int64
	checkBounds bool
}

// New は Era インスタンスを作成します。
func New(name string, wamei string, start int64, check bool) *Era {
	return &Era{name: name, wamei: wamei, start: start, checkBounds: check, year: 1}
}

// Copy は Era インスタンスのコピーを作成します。
func Copy(era *Era) *Era {
	return &Era{name: era.name, wamei: era.wamei, start: era.start, checkBounds: era.checkBounds, year: era.year, org: era.org}
}

// Check は指定した日付が年号の開始時点以降かどうかチェックします。
//
// 開始時点以降なら true
// 開始時点より前なら false
// 開始時点のチェックが不要の場合は常に true
func (era *Era) Check(t time.Time) bool {
	if !era.checkBounds {
		return true
	}
	return modjulian.DayNumber(t) >= era.start
}

// Import は指定した日付をインポートします。
//
// 開始時点以降なら true
// 開始時点より前なら false
// 開始時点が定義されていない場合は常に true
func (era *Era) Import(t time.Time) bool {
	if !era.Check(t) {
		return false
	}
	era.org = &t
	start := modjulian.ToTime(era.start, t.Location())
	era.year = int64(t.Year()) - int64(start.Year()) + 1
	return true
}

// String は年号を出力します。
func (era *Era) String() string {
	if era.wamei == "" {
		return fmt.Sprintf("%s %d", era.Name(), era.Year())
	}
	return fmt.Sprintf("%s (%s) %d", era.Wamei(), era.Name(), era.Year())
}

// Name は年号を出力します。
func (era *Era) Name() string {
	return era.name
}

// Wamei は年号を和名で出力します。
func (era *Era) Wamei() string {
	return era.wamei
}

// Year は号の開始時点を起点とした年数を出力します。
func (era *Era) Year() int64 {
	return era.year
}
