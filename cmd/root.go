package cmd

import (
	"fmt"
	"os"

	"./config"
	"./request"
	"github.com/spf13/cobra"
)

var gla struct {
	Args, ConfFile, Host, Subj, Token, Type, URI string
	Req                                          request.Conf
}

var rootCmd = &cobra.Command{
	Use:   "glat",
	Short: "CLI tool for Gitlab API",
	Long:  "CLI tool for Gitlab API",
	Run:   func(cmd *cobra.Command, args []string) {},
}

// Execute - is just main function
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConf() {
	conf, err := config.Get(gla.ConfFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	gla.Host = conf.Host
	gla.Token = conf.Token
}

func init() {
	rootCmd.PersistentFlags().StringVar(&gla.Host, "host", gla.Host, "Gitlab host")
	rootCmd.PersistentFlags().StringVarP(&gla.ConfFile, "config", "c", "", "Config file (default $HOME/.glacli.yaml)")
	rootCmd.PersistentFlags().StringVar(&gla.Token, "token", gla.Token, "Secret token")
	initConf()
}
