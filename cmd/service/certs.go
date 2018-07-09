package main

import (
	"errors"
	"io/ioutil"

	"crypto/tls"

	"github.com/stevepartridge/env"
	"github.com/stevepartridge/go/file"
	"github.com/stevepartridge/versions/insecure"
)

const (
	envTLSCertPath = "TLS_CERT_PATH"
	envTLSKeyPath  = "TLS_KEY_PATH"
)

func getCertificate() (tls.Certificate, error) {

	var (
		publicCert []byte
		privateKey []byte
		err        error
	)

	if env.GetAsBool(envTLSCertPath) && env.GetAsBool(envTLSKeyPath) {

		if file.Exists(env.Get(envTLSCertPath)) {
			publicCert, err = ioutil.ReadFile(env.Get(envTLSCertPath))
			ifError(err)
		}

		if !file.Exists(env.Get(envTLSCertPath)) {
			privateKey, err = ioutil.ReadFile(env.Get(envTLSCertPath))
			ifError(err)
		}

		if publicCert == nil {
			return tls.Certificate{}, errors.New("Unable to find public cert")
		}

		if privateKey == nil {
			return tls.Certificate{}, errors.New("Unable to find private key")
		}

		return tls.X509KeyPair(publicCert, privateKey)

	}

	return insecure.GetCertificate()
}
