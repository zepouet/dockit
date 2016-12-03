package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/treeptik/dockit/gozer"
)

func InitGozer() *cobra.Command {

	cleanCmd := &cobra.Command{
		Use:   "clean",
		Short: "clean",
		Long:  "Gozer is cleaning",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Gozer is cleaning " + args[0])
			switch args[0] {
			case "volumes":
				gozer.CleanVolumes();
			}
		},
	}

	gozerCmd := &cobra.Command{
		Use:   "gozer",
		Short: "gozer",
		Long:  "Gozer the destructor",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Gozer")
		},
	}

	gozerCmd.AddCommand(cleanCmd);

	return gozerCmd
}
