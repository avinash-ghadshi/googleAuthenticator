/*
Copyright © 2024 AVINASH GHADSHI <avinash.ghadshi@netcotecloud.com>
*/
package generate

import (
	"log"
	"github.com/spf13/cobra"
	"github.com/pquerna/otp/totp"
)

const ISSUER string = "NETCORE CLOUD PVT. LTD."
var email   string

// generateCmd represents the generate command
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate google authenticator secrete code.",
	Long: `This command generate google authenticator secrete Code.

You need to install Google Authenticator app in your mobile device.
Add the generated code in your app.
To validate verify, you can run gauth verify <OTP> command.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateCode(email)
	},
}

func init() {
	GenerateCmd.Flags().StringVarP(&email, "email", "e", "", "Account Email ID")
        if err := GenerateCmd.MarkFlagRequired("email"); err != nil {
                log.Println(err.Error())
        }
}

func generateCode(email string) {
	// Generate a new TOTP secret key
        secret, err := totp.Generate(totp.GenerateOpts{
                Issuer:      ISSUER,
                AccountName: email,
        })
        if err != nil {
                log.Fatal(err)
        }
	log.Println("Use below secrete key in your google authenticator app.")
	log.Println("Secret Key: ", secret.Secret())
}
