package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tosone/worklyzer/cmd/service"
	"github.com/tosone/worklyzer/cmd/version"
	"github.com/tosone/worklyzer/logging"
	"github.com/tosone/worklyzer/models"
	"github.com/tosone/worklyzer/scan"
)

var cfgFile string

// RootCmd represents the base command when called without any sub commands
var RootCmd = &cobra.Command{
	Use:   "worklyzer",
	Short: "",
	Long:  ``,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "get version",
	Long:  ``,
	Run: func(_ *cobra.Command, _ []string) {
		version.Initialize()
	},
}

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Fetch information from wakatime",
	Long:  `Fetch information from wakatime`,
	Run: func(_ *cobra.Command, _ []string) {
		service.Initialize()
	},
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", "/etc/worklyzer/config.yaml", "config file")

	cobra.OnInitialize(initConfig)

	RootCmd.AddCommand(serviceCmd)
	RootCmd.AddCommand(versionCmd)
}

func initConfig() {
	defaultConfig()
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile("/etc/worklyzer/config.yaml")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logging.Panic("Cannot find the special config file.")
	}

	logging.Setting()
	models.Connect()
	scan.Run()
}

func defaultConfig() {
	viper.SetDefault("DatabaseEngine", "sqlite3")
	viper.SetDefault("log", "err.log")
}
