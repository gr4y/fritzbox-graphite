package main

import (
	"github.com/codegangsta/cli"
	"github.com/gr4y/fritzbox-graphite/cmd"
	"os"
)

func main() {
	a := cli.NewApp()
	a.Name = "fritzbox-graphite"
	a.Usage = "Sends Fritz!Box Traffic Data to Graphite"
	a.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "router,r",
			Value: "fritz.box",
			Usage: "Fritz!Box Hostname or IP address",
		},
		cli.StringFlag{
			Name:  "carbon-host,ch",
			Value: "localhost",
			Usage: "Carbon Hostname or IP address",
		},
		cli.StringFlag{
			Name:  "carbon-port,cp",
			Value: "2003",
			Usage: "Carbon Port",
		},
		cli.StringFlag{
			Name:  "prefix",
			Value: "metrics.routers.7170",
			Usage: "Prefix",
		},
	}
	a.Action = cmd.CmdFetchData
	a.Run(os.Args)
}
