package main

import (
	"fmt"
	"runtime"
)

const version = 29

var cmdVersion = &Command{
	Usage: "version",
	Short: "show version info",
	Long: `

Displays the version of godep as well as the target OS, architecture and go runtime version.
`,
	Run: runVersion,
}

func runVersion(cmd *Command, args []string) {
	fmt.Printf("godep v%d (%s/%s/%s)\n", version, runtime.GOOS, runtime.GOARCH, runtime.Version())
}
