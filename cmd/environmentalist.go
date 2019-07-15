package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/j4ng5y/environmentalist/environmentalist"
	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"

	watchPhp      bool
	watchNode     bool
	vaultType     string
	vaultAuthType string

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
	environmentalistCmd.PersistentFlags().StringVarP(&vaultType, "vault-type", "v", "hashicorp-vault", "the vault-type flag tells environmentalist what vault we want to extract secrects from")
	environmentalistCmd.PersistentFlags().StringVarP(&vaultAuthType, "vault-auth-type", "", "approle", "the vault-auth-type flag tells envrionmentalist what authentication type to use to log into vault")
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
	if vaultType == "hashicorp-vault" {
		if vaultAuthType == "approle" {
			r := bufio.NewReader(os.Stdin)
			fmt.Print("Enter your Hashicorp Vault Role ID: ")
			rID, err := r.ReadString('\n')
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			s := bufio.NewReader(os.Stdin)
			fmt.Print("Enter your Hashicorp Vault Secret ID: ")
			sID, err := s.ReadString('\n')
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			H := environmentalist.NewHCV()
			H = H.GetToken(H.AppRoleAuth(rID, sID))
			// TODO: Do more things to actually start the daemon
		}
	}
}

func stopDaemon(ccmd *cobra.Command, args []string) {
	// Logic for stopping the Daemon
}
