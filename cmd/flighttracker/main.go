package main

import (
	"github.com/tmeshorer/volume/pkg/server"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	// Create the CLI application
	app := cli.App{
		Name:  "flighttracker",
		Usage: "the flight tracker cli",
		Flags: []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name:     "serve",
				Usage:    "serve the flight tracker api",
				Category: "server",
				Action:   serve,
				Flags:    []cli.Flag{},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// serve the tracker
func serve(c *cli.Context) (err error) {
	var srv *server.Server
	if srv, err = server.New(); err != nil {
		return cli.Exit(err, 1)
	}

	if err = srv.Serve(); err != nil {
		return cli.Exit(err, 1)
	}
	return nil
}
