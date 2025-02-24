package main

import (
	"github.com/spf13/cobra"
	"github.com/ztrue/tracerr"
	"os"
)

type mcpingConfig struct {
	host     string
	port     uint16
	fakeHost string
	protocol int
	timeout  int
	verbose  bool
}

const flagHost = "host"
const flagPort = "port"
const flagFakeHost = "fakehost"
const flagProtocol = "protocol"
const flagTimeout = "timeout"
const flagVerbose = "verbose"

const defaultTimeout = 5

func runCMD() {
	conf := mcpingConfig{}
	rootCmd := &cobra.Command{
		Use:     "mcping",
		Short:   "Minecraft Server List Ping tool",
		Version: version,
		Args:    cobra.MaximumNArgs(0),
		Run: func(_ *cobra.Command, _ []string) {
			if err := doPing(conf); err != nil {
				logToStderr(tracerr.SprintSourceColor(err))
				os.Exit(1)
			}
		},
	}
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	rootCmd.Flags().StringVar(&conf.host, flagHost, "127.0.0.1", "Server host")

	rootCmd.Flags().Uint16Var(&conf.port, flagPort, 0, "Server port")

	rootCmd.Flags().StringVar(&conf.fakeHost, flagFakeHost, "", "")

	rootCmd.Flags().IntVar(&conf.protocol, flagProtocol, Minecraft_1_8, "Protocol number")

	rootCmd.Flags().IntVar(&conf.timeout, flagTimeout, 0, "Timeout in seconds")

	rootCmd.Flags().BoolVar(&conf.verbose, flagVerbose, false, "Verbose output")

	rootCmd.Execute()
}
