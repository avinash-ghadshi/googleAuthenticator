/*
Copyright © 2024 AVINASH GHADSHI <avinash.ghadshi@netcorecloud.com>

*/
package verify

import (
	"log"
	"time"
	"github.com/spf13/cobra"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const SECRETE_KEY string = "MJJ6V4C4GTRXSOSSPD2NNGHW3MOSE53E"
var (
	code string
	secrete string
)

// verifyCmd represents the verify command
var VerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verifies the google authenticator code",
	Long: `This command verify the google authenticator secrete Code.

You need to install Google Authenticator app in your mobile device.
Generate the secrete code by running "gauth verify -e email"
Add the generated code in your app.
To validate verify, you can run gauth verify -c OTPCODE command.`,
	Run: func(cmd *cobra.Command, args []string) {
		verify(code, secrete)
	},
}

func init() {
	VerifyCmd.Flags().StringVarP(&code, "code", "c", "", "One Time Password (OTP) Code")
        if err := VerifyCmd.MarkFlagRequired("code"); err != nil {
                log.Println(err.Error())
        }
	VerifyCmd.Flags().StringVarP(&secrete, "secrete", "s", "", "Your secrete token")

}

func verify(code, secrete string) {

	if secrete == "" {
		secrete = SECRETE_KEY
	}
	// Validate only the current code (allowing only current time window)
	valid, _ := totp.ValidateCustom(code, secrete, time.Now(), totp.ValidateOpts{
		// Allow only the current code (past and future windows are not valid)
		// Setting this to 0 means no leniency to previous or next time windows
		Period:    30,
		Skew:      0, // Skew controls how many time steps are valid
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
        if valid {
                log.Print("SUCCESS")
        } else {
                log.Print("FAILED")
        }

}
