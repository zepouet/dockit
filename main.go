package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zepouet/dockit/cmd"
	"github.com/zepouet/dockit/docker"
)

func main() {
	var debug bool

	gozerCmd := cmd.InitGozer()
	versionCmd := cmd.InitVersion()

	// create default root command
	var rootCmd = &cobra.Command{
		Use:   "dockit",
		Short: "dockit is a useful toolkit for dev and production analysis",
		Long:  "dockit binany is a useful toolkit for dev and production analysis",

		Run: func(cmd *cobra.Command, args []string) {
			if debug {
				fmt.Println("Mode debug on")
			}
			docker.Connect()
		},
	}
	//rootCmd.Flags().IntVarP(&port, "port", "p", 8080, "docker-viz server port")
	//rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Run dockit in \"debug\" mode")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(gozerCmd)
	rootCmd.Execute()

}
