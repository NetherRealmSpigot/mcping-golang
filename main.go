package main

import (
	"fmt"
	"github.com/NetherRealmSpigot/mcping-golang/protocols"
	"github.com/spf13/cobra"
	"github.com/ztrue/tracerr"
	"os"
)

var version = "unknown"

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

func logToStdout(str string) {
	printToStdout(append([]byte(str), '\n'))
}

func printToStdout(b []byte) {
	os.Stdout.Write(b)
}

func logToStderr(str string) {
	printToStderr(append([]byte(str), '\n'))
}

func printToStderr(b []byte) {
	os.Stderr.Write(b)
}

func doPing(conf mcpingConfig) error {
	host, port, fakeHost, protocol, res, err := protocols.Ping(conf.host, conf.port, conf.fakeHost, conf.protocol, conf.timeout)
	if err != nil {
		err = tracerr.Wrap(err)
		return err
	}
	if conf.verbose {
		logToStdout(fmt.Sprintf("server: %s:%d", host, port))
		if fakeHost != host {
			logToStdout(fmt.Sprintf("fakehost: %s", fakeHost))
		}
		logToStdout(fmt.Sprintf("protocol: %d", protocol))
	}
	printToStdout(res)
	return nil
}

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
	rootCmd.Flags().IntVar(&conf.protocol, flagProtocol, protocols.Minecraft_1_8, "Protocol number")
	rootCmd.Flags().IntVar(&conf.timeout, flagTimeout, defaultTimeout, "Timeout in seconds")
	rootCmd.Flags().BoolVar(&conf.verbose, flagVerbose, false, "Verbose output")
	rootCmd.Execute()
}

func main() {
	runCMD()
}
