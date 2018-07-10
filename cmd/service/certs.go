package main

import (
	"io/ioutil"

	"crypto/tls"

	"github.com/stevepartridge/env"
	"github.com/stevepartridge/go/file"
	"github.com/stevepartridge/versions/insecure"
)

const (
	envTLSCertPath   = "TLS_CERT_PATH"
	envTLSKeyPath    = "TLS_KEY_PATH"
	envTLSRootCAPath = "TLS_ROOT_CA"
)

func getCert() ([]byte, error) {

	if env.GetAsBool(envTLSCertPath) {
		if file.Exists(env.Get(envTLSCertPath)) {
			publicCert, err := ioutil.ReadFile(env.Get(envTLSCertPath))
			if err != nil {
				return nil, err
			}
			return []byte(publicCert), nil
		}
	}

	return []byte(insecure.Cert), nil
}

func getKey() ([]byte, error) {

	if env.GetAsBool(envTLSKeyPath) {
		if file.Exists(env.Get(envTLSKeyPath)) {
			privateKey, err := ioutil.ReadFile(env.Get(envTLSKeyPath))
			if err != nil {
				return nil, err
			}
			return []byte(privateKey), nil
		}
	}

	return []byte(insecure.Key), nil
}

func getCertificate() (tls.Certificate, error) {

	publicCert, err := getCert()
	if err != nil {
		return tls.Certificate{}, err
	}

	if publicCert == nil {
		publicCert = []byte(insecure.Cert)
	}

	privateKey, err := getKey()
	if err != nil {
		return tls.Certificate{}, err
	}
	if privateKey == nil {
		privateKey = []byte(insecure.Key)
	}

	return tls.X509KeyPair(publicCert, privateKey)
}

func getRootCA() []byte {

	if env.GetAsBool(envTLSRootCAPath) {
		if file.Exists(env.Get(envTLSRootCAPath)) {
			rootCA, err := ioutil.ReadFile(env.Get(envTLSRootCAPath))
			ifError(err)
			if rootCA != nil {
				return rootCA
			}
		}
	}

	return []byte(insecure.RootCA)
}
