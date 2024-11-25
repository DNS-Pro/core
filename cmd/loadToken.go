/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// loadTokenCmd represents the loadToken command
var loadTokenCmd = &cobra.Command{
	Use:   "loadToken [flags] [Token]",
	Args:  cobra.ExactArgs(1),
	Short: "Load server config from config token",
	Run: func(cmd *cobra.Command, args []string) {
		srvCfgPath, err := cmd.Flags().GetString("server-config")
		if err != nil {
			logrus.WithError(err).Fatal("can not get \"server-config\" flag")
		}
		// ...
		srvCfgStr := args[0]
		srvCfg, err := decodeServerCfgString(srvCfgStr)
		if err != nil {
			logrus.WithError(err).Fatal("can not decode provided token")
		}
		// ...
		if err := storeJson(&srvCfg, srvCfgPath); err != nil {
			logrus.WithError(err).Fatal("can not store server config")
		}
	},
}

func init() {
	rootCmd.AddCommand(loadTokenCmd)
	loadTokenCmd.PersistentFlags().StringP("server-config", "s", "server.json", "location to store server config")
}
