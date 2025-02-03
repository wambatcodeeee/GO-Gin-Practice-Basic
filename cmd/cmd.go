package cmd

import (
	"goweb1/internal/config"
	"goweb1/internal/network"
	"goweb1/internal/repository"
	"goweb1/internal/service"
)

type Cmd struct {
	config     *config.Config
	network    *network.Network
	service    *service.Service
	repository *repository.Repository
}

func RunCmd(filePath string) *Cmd {
	cmd := &Cmd{
		config: config.NewConfig(filePath),
	}

	cmd.repository = repository.NewRepository()
	cmd.service = service.NewService(cmd.repository)
	cmd.network = network.SetupNetwork(cmd.service)

	cmd.network.RunServer(cmd.config.Server.Port)

	return cmd
}
