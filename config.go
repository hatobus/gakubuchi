package main

import "github.com/spf13/pflag"

type config struct {
	corner string
	edge string
	spaces int
}

func newConfig(args []string) (*config, error) {
	flags := pflag.NewFlagSet("", pflag.ExitOnError)

	flags.StringP("corner", "c", "+", "corner character in gakubuchi")
	flags.StringP("edge", "e", "-", "edge character in gakubuchi")
	flags.IntP("spaces", "s", 1, "spaces in head/tail")

	if err := flags.Parse(args); err != nil {
		return nil, err
	}

	var cfg config

	cfg.corner, _ = flags.GetString("corner")
	cfg.edge, _ = flags.GetString("edge")
	cfg.spaces, _ = flags.GetInt("spaces")

	return &cfg, nil
}
