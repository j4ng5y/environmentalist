package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"

	watchPhp  bool
	watchNode bool
	vaultType string

	environmentalistCmd = &cobra.Command{
		Use:     "environmentalist",
		Version: version,
		Short:   "",
		Long:    "",
		Run:     func(ccmd *cobra.Command, args []string) {},
	}

	runCmd = &cobra.Command{
		Use:   "run",
		Short: "run the environmentalist daemon",
		Long:  "",
		Run:   runDaemon,
	}

	stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "stop the envrionmentalist daemon",
		Long:  "",
		Run:   stopDaemon,
	}
)

func init() {
	environmentalistCmd.AddCommand(runCmd)
	environmentalistCmd.AddCommand(stopCmd)
	environmentalistCmd.PersistentFlags().BoolVarP(&watchPhp, "php", "p", false, "the php flag tells environmentalist that we want to watch files associated with PHP")
	environmentalistCmd.PersistentFlags().BoolVarP(&watchNode, "node", "n", false, "the node flag tells environmentalist that we want to watch files associated with NodeJS")
	environmentalistCmd.PersistentFlags().StringVarP(&vaultType, "vault-type", "v", "hashicorp-vault", "the vault flag tells environmentalist what vault we want to extract secrects from")
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
	// Logic for running the Daemon
}

func stopDaemon(ccmd *cobra.Command, args []string) {
	// Logic for stopping the Daemon
}
