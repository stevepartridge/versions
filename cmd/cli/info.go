package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print the info of the service",
	Long:  `The info of the current configured service`,
	Run:   info,
}

func info(cmd *cobra.Command, args []string) {

	serviceClient.Check()

	info, err := serviceClient.Info()
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	// fmt.Println("info:", info)
	fmt.Printf("\nService Info:\n  Version = %s (%s)\n  Built   = %s\n\n",
		info.Version,
		info.Build,
		info.BuiltAt,
	)

}
