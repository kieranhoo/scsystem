// Copyright 2023 Duc Hung Ho. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"scsystem/cmd/app"
	"sort"

	"github.com/urfave/cli/v2"

	_ "scsystem/docs"
)

func NewClient() *cli.App {
	_app := &cli.App{
		Name:        "checkin",
		Usage:       "hpcc checkin",
		Version:     "0.0.1",
		Description: "API server",
		Commands:    app.Command,
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
