package main

import (
	"github.com/urfave/cli"
	stdlog "log"
	"os"
)

// nolint:gochecknoglobals // Used for dynamically adding metadata to binary.
var (
	version = "dev"
)

// nolint:funlen
func main() {
	app := cli.NewApp()
	app.Name = "Drone Stackrox plugin"
	app.Usage = "Drone Stackrox plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		// Config flags
		cli.StringFlag{
			Name:   "action",
			Usage:  "stackrox action to perform (scan, check)",
			Value:  "scan",
			EnvVar: "PLUGIN_ACTION",
		},
		cli.StringFlag{
			Name:   "rox_central_address",
			Usage:  "stackrox central adress url",
			EnvVar: "PLUGIN_ROX_CENTRAL_ADDRESS",
		},
		cli.StringFlag{
			Name:   "image",
			Usage:  "image need to be scanned",
			EnvVar: "PLUGIN_IMAGE",
		},
		cli.StringFlag{
			Name:   "token",
			Usage:  "stackrox token",
			EnvVar: "PLUGIN_TOKEN",
		},
		cli.StringFlag{
			Name:   "output",
			Usage:  "scan result output",
			Value:  "json",
			EnvVar: "PLUGIN_OUTPUT",
		},
		cli.StringFlag{
			Name:   "output-file",
			Usage:  "scan result output",
			Value:  "json",
			EnvVar: "PLUGIN_OUTPUT_FILE_DIR",
		},
		cli.StringFlag{
			Name:   "categories",
			Usage:  "image check categories",
			EnvVar: "PLUGIN_CATEGORIES",
		},
	}

	if err := app.Run(os.Args); err != nil {
		stdlog.Fatalf("%#v", err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Action:     c.String("action"),
		Url:        c.String("rox_central_address"),
		Image:      c.String("image"),
		Token:      c.String("token"),
		Output:     c.String("output"),
		OutputFile: c.String("output-file"),
		Categories: c.String("categories"),
	}
	return plugin.Exec()
}
