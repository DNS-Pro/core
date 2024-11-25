/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// dumpTokenCmd represents the dumpToken command
var dumpTokenCmd = &cobra.Command{
	Use:   "dumpToken",
	Short: "Dump server config into config token",
	Run: func(cmd *cobra.Command, args []string) {
		srvCfgPath, err := cmd.Flags().GetString("server-config")
		if err != nil {
			logrus.WithError(err).Fatal("can not get \"server-config\" flag")
		}
		srvCfg, err := loadJson[*serverCfg](srvCfgPath)
		if err != nil {
			logrus.WithError(err).Fatal("error loading server config")
		}
		srvCfgStr, err := srvCfg.encodeString()
		if err != nil {
			logrus.WithError(err).Fatal("can not encode server config")
		}
		fmt.Printf("server config token is: %s", srvCfgStr)
	},
}

func init() {
	rootCmd.AddCommand(dumpTokenCmd)
	dumpTokenCmd.PersistentFlags().StringP("server-config", "s", "server.json", "location of server config file")
}
