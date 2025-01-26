package cmd

import (
	"goweb1/internal/config"
	"goweb1/internal/network"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func RunCmd(filePath string) *Cmd {
	cmd := &Cmd{
		config:  config.NewConfig(filePath),
		network: network.SetupNetwork(),
	}

	cmd.network.RunServer(cmd.config.Server.Port)

	return cmd
}
