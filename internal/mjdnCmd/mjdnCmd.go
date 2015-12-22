package mjdnCmd

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spiegel-im-spiegel/astrocalc/mjdn"
	"github.com/spiegel-im-spiegel/facade"
)

// Name は mjdn コマンド名を定義する
const Name string = "mjdn"

// Context は mjdn コマンドのコンテキストを定義する
type Context struct {
	//Embedded facade.Context
	*facade.Context
	//AppName にはアプリケーション名を格納する
	AppName string
}

// Command は Context のインスタンスを返す
func Command(cxt *facade.Context, appName string) *Context {
	return &Context{Context: cxt, AppName: appName}
}

// Synopsis は mjdn コマンドの概要を返す
func (c Context) Synopsis() string {
	return "Calculation of Modified Julian Day"
}

// Help は mjdn コマンドのヘルプを返す
func (c Context) Help() string {
	helpText := `
Usage: astrocalc mjdn <year> <month> <day>
`
	return fmt.Sprintln(strings.TrimSpace(helpText))
}

// Run は mjdn コマンドを実行する
func (c Context) Run(args []string) int {
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.Usage = func() {
		c.Error(c.Help())
	}
	// Parse commandline flag
	if err := flags.Parse(args); err != nil {
		return facade.ExitCodeError
	}
	if flags.NArg() != 3 {
		c.Error(fmt.Sprintf("年月日を指定してください\n\n%s", c.Help()))
		return facade.ExitCodeError
	}
	argsStr := flags.Args()
	var ymd = make([]int, 3)
	for i, arg := range argsStr {
		num, err := strconv.Atoi(arg)
		if err != nil {
			c.Error(fmt.Sprintln(err))
			return facade.ExitCodeError
		}
		ymd[i] = num
	}
	tm := time.Date(ymd[0], time.Month(ymd[1]), ymd[2], 0, 0, 0, 0, time.UTC)
	c.Output(fmt.Sprint(mjdn.DayNumber(tm)))
	return facade.ExitCodeOK
}
