/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

func getTimeInTimezone(timezone string) (string, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}
	currentTime := time.Now().In(location)
	return currentTime.Format(time.RFC1123), nil
}

var timezoneCmd = &cobra.Command{
	Use:   "timezone",
	Short: "Get the current time in a given timezone",
	Long: `Get the current time in a given timezone.
               This command takes one argument, the timezone you want to get the current time in.
               It returns the current time in RFC1123 format.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		timezone := args[0]
		currentTime, err := getTimeInTimezone(timezone)
		if err != nil {
			log.Fatalln("The timezone string is invalid")
		}
		fmt.Println(currentTime)
	},
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-commands",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(timezoneCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-commands.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
