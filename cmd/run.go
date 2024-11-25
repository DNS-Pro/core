/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/DNS-Pro/core/pkg/app"
	"github.com/DNS-Pro/core/pkg/errs"
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run dns pro client",
	Run: func(cmd *cobra.Command, args []string) {
		var srvCfg *serverCfg
		var clnCfg *clientCfg
		// ...
		if tok, err := cmd.Flags().GetString("server-token"); err != nil {
			logrus.WithError(err).Error("can not get \"server-token\" flag")
		} else if tok != "" {
			srvCfg, err = decodeServerCfgString(tok)
			if err != nil {
				logrus.WithError(err).Error("can not decode server-token")
			}
		}
		if srvCfg == nil {
			srvCfgPath, err := cmd.Flags().GetString("server-config")
			if err != nil {
				logrus.WithError(err).Fatal("can not get \"server-config\" flag")
			}
			srvCfg, err = loadJson[*serverCfg](srvCfgPath)
			if err != nil {
				logrus.WithError(err).Fatal("can not load server-config")
			}
		}
		// ...
		clnCfgPath, err := cmd.Flags().GetString("client-config")
		if err != nil {
			logrus.WithError(err).Fatal("can not get \"client-config\" flag")
		}
		clnCfg, err = loadJson[*clientCfg](clnCfgPath)
		if err != nil {
			logrus.WithError(err).Fatal("can not load client-config")
		}
		// ...
		apCfg, err := app.NewAppConfig((*app.ClientConfig)(clnCfg), &srvCfg.DNS, &srvCfg.Authenticator)
		if err != nil {
			logrus.WithError(err).Fatal("can not generate app config")
		}
		ap, err := app.NewApp(apCfg)
		if err != nil {
			logrus.WithError(err).Fatal("can not create app")
		}
		ctx := context.TODO()
		ap.Run(ctx)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().StringP("server-config", "s", "server.json", "location of server config file")
	runCmd.PersistentFlags().StringP("server-token", "t", "", "server config token")
	runCmd.PersistentFlags().StringP("client-config", "c", "client.json", "location of client config file")
}

func loadJson[T interface{}](path string) (T, error) {
	file, err := os.Open(path)
	if err != nil {
		return *new(T), fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)

	var v T
	err = json.Unmarshal(byteValue, &v)
	if err != nil {
		return *new(T), fmt.Errorf("error unmarshalling JSON: %s", err)
	}
	if err := defaults.Set(v); err != nil {
		return *new(T), errs.NewConfigDefaultValueErr(err)
	}
	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(v); err != nil {
		return *new(T), errs.NewConfigValidationErr(err)
	}
	return v, nil
}

func storeJson[T interface{}](val T, path string) error {
	if err := defaults.Set(val); err != nil {
		return errs.NewConfigDefaultValueErr(err)
	}
	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(val); err != nil {
		return errs.NewConfigValidationErr(err)
	}
	// ...
	jsonData, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %s", err)

	}
	// ...
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file: %s", err)
	}
	defer file.Close()
	if _, err = file.Write(jsonData); err != nil {
		return fmt.Errorf("error writing to file: %s", err)
	}
	return nil
}
