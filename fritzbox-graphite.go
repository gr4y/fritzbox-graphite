package main

import (
	"github.com/codegangsta/cli"
	"github.com/gr4y/fritzbox-graphite/cmd"
	"os"
	// "time"
)

func main() {
	a := cli.NewApp()
	a.Name = "fritzbox-graphite"
	a.Usage = "Sends Fritz!Box Traffic Data to Graphite"
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
