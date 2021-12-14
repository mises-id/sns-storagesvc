package main

import (
	"context"
	"os"
	"time"

	"github.com/mises-id/sns-storagesvc/cmd/web/rest"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	app := cli.NewApp()
	app.Action = func(c *cli.Context) error {
		return rest.Start(ctx)
	}
	app.Commands = cli.Commands{
		{
			Name:  "",
			Usage: "./sns-storagesvc",
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) error {
				return rest.Start(ctx)
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
