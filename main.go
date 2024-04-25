package hpcrencryptioncertificate

import (
	_ "embed"
)

//go:embed certificate/ibm-hyper-protect-container-runtime-1-0-s390x-15-encrypt.crt
var EncryptionCertificate string
