package main

import (
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/astrocalc/internal/mjdnCmd"
	"github.com/spiegel-im-spiegel/gofacade"
)

const (
	// Name はアプリケーション名を定義する
	Name string = "astrocalc"
	// Version はアプリケーションのバージョン番号を定義する
	Version string = "0.1.0"
)

func setupFacade(cxt *gofacade.Context) *gofacade.Facade {
	fcd := gofacade.NewFacade(cxt)
	fcd.AddCommand(mjdnCmd.Name, mjdnCmd.Command(cxt, Name))
	return fcd
}

func main() {
	cxt := gofacade.NewContext(os.Stdin, os.Stdout, os.Stderr)
	fcd := setupFacade(cxt)
	rtn, err := fcd.Run(Name, Version, os.Args[1:])
	if err != nil {
		cxt.Error(fmt.Sprintln(err))
	}
	os.Exit(rtn)
}
