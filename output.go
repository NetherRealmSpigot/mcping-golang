package main

import "os"

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
