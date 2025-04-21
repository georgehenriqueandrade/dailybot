package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "dailybot",
		Short: "DailyBot is a personal automation CLI tool",
		Long:  "DailyBot helps you automate daily tasks like organizing files, checking websites, sending reminders, and more.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to DailyBot! Use --help to see available commands.")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
