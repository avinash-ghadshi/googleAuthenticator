/*
Copyright © 2024 AVINASH GHADSHI <avinash.ghadshi@netcorecloud.com>

*/
package cmd

import (
	"os"

	"github.com/avinash-ghadshi/gauth/cmd/generate"
	"github.com/avinash-ghadshi/gauth/cmd/verify"
	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gauth",
	Short: "CLI tool for generating and verifying Google Authenticator codes",
	Long: `The gauth is a command-line tool for managing Google Authenticator codes.
    
This tool allows you to:
1. Generate Time-based One-Time Password (TOTP) codes compatible with Google Authenticator.
2. Verify the generated codes for authentication purposes.

Use this tool for adding an extra layer of security to your applications via two-factor authentication (2FA).`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func generateSubCommand() {
        rootCmd.AddCommand(generate.GenerateCmd)
}

func verifySubCommand() {
        rootCmd.AddCommand(verify.VerifyCmd)
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	generateSubCommand()
	verifySubCommand()
}


