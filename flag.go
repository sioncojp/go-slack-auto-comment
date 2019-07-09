package main

import (
	"github.com/urfave/cli"
)

// FlagSet ...flagを設定
func FlagSet() *cli.App {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config-dir, c",
			Usage: "Load configuration *.toml in target dir",
		},
		cli.StringFlag{
			Name:  "region, r",
			Value: "ap-northeast-1",
			Usage: "Setting AWS region for tomlssm",
		},
	}
	return app
}
