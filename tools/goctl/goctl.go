package main

import (
	"github.com/5a6e/newzero/core/load"
	"github.com/5a6e/newzero/core/logx"
	"github.com/5a6e/newzero/tools/goctl/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
