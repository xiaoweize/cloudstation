/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cloudstation",
	Short: "cloudstation",
	Long:  `cloudstation 文件云中转站，支持阿里云/腾讯云/亚马逊云OSS存储`,
	//不带子命令时的处理逻辑
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			fmt.Println("cloudstation v0.0.1")
		}
		//输出帮助
		return cmd.Usage()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//该flag仅仅适用于根命令root
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "cloudstation version")
}
