package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().Int32P("log-level", "l", 4, "set the logging level")
	RootCmd.PersistentFlags().StringP("config", "c", "", "set the configuration file")

	if err := viper.BindPFlags(RootCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	viper.SetEnvPrefix("achievements")
	viper.AutomaticEnv()

	level := viper.GetInt("log-level")
	log.SetLevel(log.AllLevels[4])
	if level < len(log.AllLevels) && level >= 0 {
		log.SetLevel(log.AllLevels[level])
	}

	// Set config search path
	viper.AddConfigPath("/etc/achievements/")
	viper.AddConfigPath("$HOME/.achievements")
	viper.AddConfigPath(".")

	// Load config
	viper.SetConfigName("config")
	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Debug("No config file found")
		} else {
			log.Panicf("Fatal error in config file: %v \n", err)
		}
	}

	// Load user defined config
	cfg := viper.GetString("config")
	if cfg != "" {
		viper.SetConfigFile(cfg)
		err := viper.ReadInConfig()
		if err != nil {
			log.Panicf("Fatal error in config file: %v \n", err)
		}
	}
}

// RootCmd is the cli root command
var RootCmd = &cobra.Command{
	Use:   "achievements",
	Short: "Launch the startup achievements back-end",
	Run:   func(cmd *cobra.Command, arguments []string) {},
}
