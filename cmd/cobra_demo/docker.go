package main

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
)

func main() {
	var cmdPull = &cobra.Command{
		Use:   "pull [OPTIONS] NAME[:TAG|@DIGEST]",
		Short: "Pull an image or a repository from a registry",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Pull: " + strings.Join(args, " "))
		},
	}

	var rootCmd = &cobra.Command{Use: "docker"}
	rootCmd.AddCommand(cmdPull)
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	rootCmd.Execute()
}