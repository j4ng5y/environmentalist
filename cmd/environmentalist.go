package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"

	environmentalistCmd = &cobra.Command{
		Use:     "environmentalist",
		Version: version,
		Short:   "",
		Long:    "",
		Run:     func(ccmd *cobra.Command, args []string) {},
	}

	runCmd = &cobra.Command{
		Use:   "run",
		Short: "",
		Long:  "",
		Run:   runDaemon,
	}

	stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "",
		Long:  "",
		Run:   stopDaemon,
	}
)

func init() {
	environmentalistCmd.AddCommand(runCmd)
	environmentalistCmd.AddCommand(stopCmd)
}

// Execute runs the CLI
func Execute() {
	err := environmentalistCmd.Execute()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func runDaemon(ccmd *cobra.Command, args []string) {

}

func stopDaemon(ccmd *cobra.Command, args []string) {

}
