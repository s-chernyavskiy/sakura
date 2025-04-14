package config

import "github.com/jessevdk/go-flags"

type Opts struct {
	Port uint   `short:"p" long:"port" description:"Port for application" default:"6543"`
	Host string `long:"host" description:"Host for application" default:"127.0.0.1"`
}

func NewOpts(args []string) *Opts {
	var options Opts
	args, err := flags.ParseArgs(&options, args)
	if err != nil {
		panic("opts init error")
	}

	return &options
}
