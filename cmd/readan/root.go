package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFilePath string

var rootCmd = &cobra.Command{
	Use:   "readan",
	Short: "Readan server",
	Args:  cobra.ExactArgs(0),
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&configFilePath, "config", "c", "config.dev.yaml", "config file path")
}

func loadConfig() {
	viper.SetConfigFile(configFilePath)
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("failed to read config file: %v", err)
	}
}
