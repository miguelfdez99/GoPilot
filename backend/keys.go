package backend

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var acceptedKeyTypes = []string{"rsa", "dsa", "ecdsa", "ed25519"}

func (b *Backend) GenerateKeys(keyType, keyName, outputPath string, overwrite bool) (privateKeyPath, publicKeyPath string, err error) {

	isAccepted := false
	for _, acceptedType := range acceptedKeyTypes {
		if strings.ToLower(keyType) == acceptedType {
			isAccepted = true
			break
		}
	}
	if !isAccepted {
		return "", "", fmt.Errorf("provided key type '%s' is not accepted", keyType)
	}

	privateKeyPath = fmt.Sprintf("%s/%s", outputPath, keyName)
	publicKeyPath = fmt.Sprintf("%s/%s.pub", outputPath, keyName)

	if !overwrite {
		if _, err := os.Stat(privateKeyPath); err == nil {
			return "", "", fmt.Errorf("private key file already exists: %s", privateKeyPath)
		}
		if _, err := os.Stat(publicKeyPath); err == nil {
			return "", "", fmt.Errorf("public key file already exists: %s", publicKeyPath)
		}
	}

	_, err = exec.Command("ssh-keygen", "-t", keyType, "-f", privateKeyPath, "-N", "").Output()
	if err != nil {
		return "", "", fmt.Errorf("failed to generate keys: %w", err)
	}
	return
}
