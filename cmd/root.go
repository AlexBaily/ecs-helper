package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(descClusterCmd)
}

var rootCmd = &cobra.Command{
	Use:   "ecs-helper",
	Short: "ecs-helper is a command line utility for ECS maintenance",
	Long: `ecs-helper is a command line utility for ECS maintenance, 
troubleshooting and other activies.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of ecs-helper",
	Long:  `Print the version of ecs-helper`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ecs-helper v0.0.1")
	},
}

var descClusterCmd = &cobra.Command{
	Use:   "describe-clusters",
	Short: "Prints information on ECS clusters",
	Long: `Prints information on ECS clusters
includes information on ecs services and host instances.`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO ecs.DescribeClusters()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
