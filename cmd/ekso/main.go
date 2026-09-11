package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/whatnaut96/ekso/internal/configuration"
	yaml "go.yaml.in/yaml/v4"
)

type EksoConfig struct {
	Hosts []configuration.InventoryItem `yaml:"hosts"`
}

func main() {
	command := &cli.Command{
		UseShortOptionHandling: true,
		Name:                   "ekso",
		Usage:                  "a distributed task runner for any host",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "host-file", Aliases: []string{"H"}, Usage: "path to the config file defining inventory and procedures"},
			&cli.StringSliceFlag{Name: "host", Aliases: []string{"i"}, Usage: "name of the inventory set to run the procedure on"},
			&cli.BoolFlag{Name: "verbose", Aliases: []string{"v"}, Usage: "enable verbose output"},
		},
		OnUsageError: func(ctx context.Context, cmd *cli.Command, err error, isSubcommand bool) error {
			fmt.Fprintf(cmd.ErrWriter, "Error: %v\n\n", err)
			cli.ShowCommandHelp(ctx, cmd, "ekso")
			return err
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			var configRaw []byte
			var err error
			filePath := cmd.String("file")
			if filePath != "" {
				configRaw, err = os.ReadFile(filePath)
				if err != nil {
					return fmt.Errorf("could not read the config file: %v", err)
				}
				var eksoConfig EksoConfig
				err = yaml.Unmarshal(configRaw, &eksoConfig)
				if err != nil {
					return fmt.Errorf("could not deserialize ekso config: %v", err)
				}

				if len(eksoConfig.Hosts) == 0 {
					return fmt.Errorf("config contains no hosts")
				}

				for _, host := range cmd.StringSlice("host") {
					slog.Info("got it", "h", host)
				}

			} else {
				return fmt.Errorf("no configuration file specified")
			}
			return nil
		},
	}

	if err := command.Run(context.Background(), os.Args); err != nil {
		os.Exit(1)
	}
}
