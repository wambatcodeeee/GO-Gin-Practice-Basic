package main

import (
	"flag"
	"goweb1/cmd"
)

var configFlag = flag.String("config", "", "config file not found")

func main() {
	flag.Parse()
	cmd.RunCmd(*configFlag)
}
