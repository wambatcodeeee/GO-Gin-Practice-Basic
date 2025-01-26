package main

import (
	"flag"
	"goweb1/cmd"
)

var configFlag = flag.String("config", "./cmd/main/config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.RunCmd(*configFlag)
}
