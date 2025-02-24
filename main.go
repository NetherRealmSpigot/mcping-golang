package main

import (
	"github.com/ztrue/tracerr"
)

var version = "unknown"

func doPing(conf mcpingConfig) error {
	res, err := Ping(conf.host, conf.port, conf.fakeHost, conf.protocol, conf.timeout, conf.verbose)
	if err != nil {
		err = tracerr.Wrap(err)
		return err
	}
	printToStdout(res)
	return nil
}

func main() {
	runCMD()
}
