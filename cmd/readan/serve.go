package main

import (
	"net/http"

	"github.com/aryahadii/readan/apis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the http server",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		loadConfig()

		if err := serve(); err != nil {
			logrus.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() error {
	router := apis.InitAndGetHTTPRouter()
	serverAddr := viper.GetString("server.addr")

	logrus.Infof("serving on '%s'", serverAddr)
	return http.ListenAndServe(serverAddr, router)
}
