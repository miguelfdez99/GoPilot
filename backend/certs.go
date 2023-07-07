package backend

import (
	"log"
	"os/exec"
)

func (b *Backend) GenerateSelfSignedCertificate(privateKeyPath string) (certificatePath string, err error) {
	certificatePath = privateKeyPath + "-cert.pub"

	_, err = exec.Command("ssh-keygen", "-s", privateKeyPath, "-I", "key_id", "-n", "user", "-V", "+52w", privateKeyPath+".pub").Output()
	if err != nil {
		log.Fatal(err)
	}

	return
}
