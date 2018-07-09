package insecure

import (
	"crypto/tls"
)

func GetCertificate() (tls.Certificate, error) {
	return tls.X509KeyPair([]byte(Cert), []byte(Key))
}
