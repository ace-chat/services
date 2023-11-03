package pkg

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateCode() string {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "ace",
		AccountName: "ace",
		SecretSize:  11,
		Algorithm:   otp.AlgorithmSHA512,
	})

	if err != nil {
		return ""
	}

	// key.URL()

	return key.Secret()
}

func ValidateCode(passcode string, secret string) bool {
	return totp.Validate(passcode, secret)
}
