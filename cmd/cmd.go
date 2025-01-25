package cmd

import (
	"goweb1/internal/config"
)

type Cmd struct {
	config *config.Config
}

func RunCmd(filePath string) *Cmd {
	cmd := &Cmd{
		config: config.NewConfig(filePath),
	}

	return cmd
}
