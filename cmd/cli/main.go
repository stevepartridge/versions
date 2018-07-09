package main

import (
	"fmt"
	"os"

	"crypto/x509"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"

	"github.com/stevepartridge/versions/client"
	"github.com/stevepartridge/versions/insecure"
)

var (
	useInsecure bool

	serviceUri    string
	serviceClient *client.Client

	rootCmd = &cobra.Command{
		Use:   "version",
		Short: "A version management CLI",
	}
)

func main() {

	var err error

	cobra.OnInitialize(initEnv)

	rootCmd.PersistentFlags().BoolVar(&useInsecure, "insecure", true, "Use insecure certs to connect.")
	rootCmd.PersistentFlags().StringVar(&serviceUri, "service", "versions.local:8000", "The full URI for the service.")

	serviceClient, err = client.NewClient(serviceUri)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if useInsecure {

		fmt.Print("\n ==! Using insecure certs !== \n")

		serviceClient.CertPool = x509.NewCertPool()

		if !serviceClient.CertPool.AppendCertsFromPEM([]byte(insecure.RootCA)) {
			fmt.Println("Failed appending Root CA file")
			return
		}
	}

	rootCmd.AddCommand(infoCmd)

	versionCommands(rootCmd)
	applicationCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initEnv() {

	_, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(" ")

}
