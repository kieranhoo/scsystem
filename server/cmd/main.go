// Copyright 2023 Duc Hung Ho. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"os/signal"
	"scsystem/api"
	"sort"
	"syscall"

	"github.com/urfave/cli/v2"

	_ "scsystem/docs"
)

var Command = []*cli.Command{
	{
		Name:    "worker",
		Aliases: []string{"w"},
		Usage:   "run worker handle tasks in queue",
		Action: func(c *cli.Context) error {
			w := api.NewWorker(10)
			return w.Run()
		},
	},
	{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "run api server",
		Action: func(c *cli.Context) error {
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
			coreAPI := api.NewServer().Shutdown(sigChan)
			return coreAPI.Run()
		},
	},
}

func NewClient() *cli.App {
	_app := &cli.App{
		Name:        "checkin",
		Usage:       "hpcc checkin",
		Version:     "0.0.1",
		Description: "API server",
		Commands:    Command,
	}

	sort.Sort(cli.FlagsByName(_app.Flags))
	sort.Sort(cli.CommandsByName(_app.Commands))

	return _app
}

func main() {
	client := NewClient()

	if err := client.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
