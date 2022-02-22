package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type Terminal struct {
	CLI    *cli.App
	Logger *log.Logger
	Client *Client
}

func NewTerminal(progName, progUsage string, config *Client) *Terminal {
	termCLI := cli.NewApp()
	termCLI.Name = progName
	termCLI.Usage = progUsage

	logger := log.New()
	logger.SetLevel(log.InfoLevel)

	return &Terminal{
		CLI:    termCLI,
		Logger: logger,
		Client: config,
	}
}

func (t *Terminal) AddFlag(flag cli.StringFlag) {
	t.CLI.Flags = append(t.CLI.Flags, flag)
}

func (t *Terminal) AddCommand(command cli.Command) {
	t.CLI.Commands = append(t.CLI.Commands, command)
}
