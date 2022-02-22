package main

import (
	"log"
	"os"

	"github.com/rombintu/square_colony/cli/utils"
	"github.com/urfave/cli"
)

func buildClientCLI(term *utils.Terminal) {
	term.AddCommand(
		cli.Command{
			Name:  "ping",
			Usage: "Ping to server",
			Action: func(c *cli.Context) error {
				ping, err := term.Client.PingServer()
				if err != nil {
					return err
				}
				term.Logger.Info(ping)
				return nil
			},
		},
	)
}

func main() {
	config := utils.GetClientConfig("./config.toml")
	client := utils.NewClient(config)
	terminal := utils.NewTerminal("Golearncli", "Golearn CLI", client)

	buildClientCLI(terminal)

	err := terminal.CLI.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
