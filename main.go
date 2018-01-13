package main

import (
	"runtime"

	"github.com/tosone/worklyzer/cmd"
	"github.com/tosone/worklyzer/cmd/version"
	"github.com/tosone/worklyzer/logging"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Version version
var Version = "no provided"

// BuildStamp BuildStamp
var BuildStamp = "no provided"

// GitHash GitHash
var GitHash = "no provided"

func main() {
	if runtime.GOOS == "windows" {
		logging.Panic("Worklyzer not support windows just linux.")
	}

	version.Setting(Version, BuildStamp, GitHash)

	if err := cmd.RootCmd.Execute(); err != nil {
		logging.Panic(err.Error())
	}
}
