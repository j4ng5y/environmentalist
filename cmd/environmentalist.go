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

	hashiVault            bool
	hashiVaultAuthType    string
	awsSSM                bool
	awsSSMCredentialType  string
	awsSSMRegion          string
	awsSSMProfileName     string
	awsSSMAccessKeyID     string
	awsSSMSecretAccessKey string

	environmentalistCmd = &cobra.Command{
		Use:     "environmentalist",
		Version: version,
		Short:   "An appliction to provide a consistent API for managing secrets",
		Long: `Environmentalist is an application that provides a consistent API for using a number of secrets management tools including:
  * Hashicorp Vault
  * AWS SSM
  * Ansible Vault
	etc...
	
The Server runs as both a RESTful service as well as a gRPC service so it should be usable for almost any situation.

A RESTful request to access a secret looks something like this:
  curl -X GET https://environmentalist:5005/hashicorp-vault/get/mySharedSecret

A RESTful request to store a new secret looks something like this:
  curl -X POST -H "Content-Type: application/json" -d '{"mySharedSecret": "thisIsASuperSecretPassword"}' https://environmentalist:5005/hashicorp-vault/new/mySharedSecret

A RESTful request to delete a secret looks something like this:
  curl -X DELETE https://environmentalist:5005/hashicorp-vault/delete/mySharedSecret

A RESTful request to modify a secret looks something like this:
  curl -X PUT -H "Content-Type: application/json" -d '{"mySharedSecret": "thisIsANewSuperSecretPassword"}' https://environmentalist:5005/hashicorp-vault/update/mySharedSecret

Please see https://github.com/j4ng5y/envrionmentalist for a full API breakdown.`,
		Run: func(ccmd *cobra.Command, args []string) {},
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
	environmentalistCmd.PersistentFlags().BoolVarP(&hashiVault, "hashicorp-vault", "v", false, "the hashicorp-vault flag tells environmentalist that we want to use the hashicorp vault")
	environmentalistCmd.PersistentFlags().StringVarP(&hashiVaultAuthType, "hashicorp-vault-auth-type", "", "approle", "the hashicorp-vault-auth-type flag tells envrionmentalist what authentication type to use to log into the hashi-corp vault")
	environmentalistCmd.PersistentFlags().BoolVarP(&awsSSM, "aws-ssm", "s", false, "the aws-ssm flag tells envrionmentalist that we want to use aws ssm")
	environmentalistCmd.PersistentFlags().StringVarP(&awsSSMCredentialType, "aws-ssm-credential-type", "", "profile", "the aws-ssm-credential-type flag tells environmentalist what type of credentials to look for to access AWS. (Options:\"profile\", \"manual\", \"role\"")
	environmentalistCmd.PersistentFlags().StringVarP(&awsSSMRegion, "aws-ssm-region", "", "us-east-1", "the aws-ssm-region flag tells envrionmentalist what AWS region to connect to")
	environmentalistCmd.PersistentFlags().StringVarP(&awsSSMProfileName, "aws-ssm-profile-name", "", "Default", "the aws-ssm-profile-name flag tells envrionmentalist what AWS profile to connect to AWS with")
	environmentalistCmd.PersistentFlags().StringVarP(&awsSSMAccessKeyID, "aws-ssm-access-key-id", "", "", "the aws-ssm-access-key-id flag tells envrionmentalist what Access Key to use to connect to AWS with")
	environmentalistCmd.PersistentFlags().StringVarP(&awsSSMSecretAccessKey, "aws-ssm-secret-access-key", "", "", "the aws-ssm-secret-access-key flag tells envrionmentalist what Secret Key to use to connect to AWS with")
	if hashiVault {
		environmentalistCmd.MarkFlagRequired("hashicorp-vault-auth-type")
	}
	if awsSSM {
		environmentalistCmd.MarkFlagRequired("aws-ssm-credential-type")
	}
	switch awsSSMCredentialType {
	case "profile":
		environmentalistCmd.MarkFlagRequired("aws-ssm-region")
		environmentalistCmd.MarkFlagRequired("aws-ssm-profile-name")
	case "manual":
		environmentalistCmd.MarkFlagRequired("aws-ssm-region")
		environmentalistCmd.MarkFlagRequired("aws-ssm-access-key-id")
		environmentalistCmd.MarkFlagRequired("aws-ssm-secret-access-key")
	case "role":
		log.Print("not implimented yet")
	default:
		log.Print("Invalid aws-ssm-credential-type. Must be one of type: \"profile\", \"manual\", or \"role\"")
	}
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
	if hashiVault {
		if hashiVaultAuthType == "approle" {
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
