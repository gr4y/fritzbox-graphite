package main

import (
	"os"

	"github.com/gr4y/fritzbox-graphite/cmd"
	"github.com/urfave/cli"
	// "time"
)

func main() {
	a := cli.NewApp()
	a.Name = "fritzbox-graphite"
	a.Usage = "Sends Fritz!Box Traffic Data to Graphite"
	a.Version = "0.4.0"
	a.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config,c",
			Value: "/etc/fritzbox-graphite/settings.json",
			Usage: "Path to JSON-Configuration file",
		},
	}
	a.Action = cmd.CmdFetchData
	a.Run(os.Args)
}
