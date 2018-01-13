package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/FlorentinDUBOIS/achievements/middlewares"
	"github.com/FlorentinDUBOIS/achievements/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/willf/pad"
)

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().Int32P("log-level", "", 4, "set the logging level")
	RootCmd.PersistentFlags().StringP("config", "c", "", "set the configuration file")

	RootCmd.Flags().Int32P("listen", "l", 9300, "set the port to listen")

	if err := viper.BindPFlags(RootCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}

	if err := viper.BindPFlags(RootCmd.Flags()); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	viper.SetEnvPrefix("achievements")
	viper.AutomaticEnv()

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

	level := viper.GetInt("log-level")
	log.SetLevel(log.AllLevels[4])
	if level < len(log.AllLevels) && level >= 0 {
		log.SetLevel(log.AllLevels[level])
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
	Run: func(cmd *cobra.Command, arguments []string) {
		r := echo.New()
		listen := viper.GetInt("listen")

		r.Logger.SetOutput(ioutil.Discard)

		r.Use(middleware.Gzip())
		r.Use(middleware.CORS())
		r.Use(middleware.Recover())
		r.Use(middleware.MethodOverride())

		r.Use(middlewares.Logger())

		router.APIv0.Register(r.Group("/api/v0"))

		for _, route := range r.Routes() {
			log.Debugf("%s %s ---> %s", pad.Right(route.Method, 7, " "), pad.Right(route.Path, 40, " "), route.Name)
		}

		log.Infof("Listen at :%d", listen)
		if err := r.Start(fmt.Sprintf(":%d", listen)); err != nil {
			log.Fatal(err)
		}
	},
}
