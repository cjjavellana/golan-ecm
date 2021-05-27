package cli

import "flag"

type CmdLineArgs struct {
	ConfigFile string
}

func ParseCli() *CmdLineArgs {
	var cfgFile string

	flag.StringVar(&cfgFile, "config", "config.yaml", "The configuration file to use")
	flag.Parse()

	return &CmdLineArgs{
		cfgFile,
	}
}
