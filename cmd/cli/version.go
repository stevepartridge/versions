package main

import (
	"fmt"

	"github.com/spf13/cobra"
	// "golang.org/x/net/context"
	// "github.com/stevepartridge/versions"
	// pb "github.com/stevepartridge/versions/protos"
)

func versionCommands(cmd *cobra.Command) {

	cmd.PersistentFlags().IntP("id", "i", 0, "ID to get")

	// Create
	createCmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c", "add"},
		Short:   "Create version",
		Run:     createVersion,
	}

	createCmd.PersistentFlags().StringP("version", "v", "", "Version ex: 1.2.3")
	createCmd.PersistentFlags().StringP("name", "n", "", "Version friendly name")
	createCmd.PersistentFlags().StringP("app-id", "", "", "Application ID")

	cmd.AddCommand(createCmd)

	// Get
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get version",
		Run:   getVersion,
	}

	cmd.AddCommand(getCmd)

	// Update
	updateCmd := &cobra.Command{
		Use:     "update",
		Aliases: []string{"u"},
		Short:   "Update version",
		Run:     updateVersion,
	}

	updateCmd.PersistentFlags().StringP("version", "v", "", "Version ex: 1.2.3")
	updateCmd.PersistentFlags().StringP("name", "n", "", "Version friendly name")
	updateCmd.PersistentFlags().StringP("app-id", "", "", "Application ID")
	updateCmd.PersistentFlags().StringP("stable", "s", "", "Stable flag for version")

	cmd.AddCommand(updateCmd)

	// Delete
	deleteCmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"u"},
		Short:   "Delete version",
		Run:     deleteVersion,
	}

	cmd.AddCommand(deleteCmd)

}

func createVersion(cmd *cobra.Command, args []string) {
	serviceClient.Check()

	var (
		version, _ = cmd.Flags().GetString("version")
		name, _    = cmd.Flags().GetString("name")
		appId, _   = cmd.Flags().GetString("app-id")
	)

	fmt.Println("create version", version, name, appId)

	resp, err := serviceClient.CreateVersion(version, name, appId)

	if err != nil {
		fmt.Println("Error", err.Error())
	}

	fmt.Println("version: %+v", resp)
}

func updateVersion(cmd *cobra.Command, args []string) {
	serviceClient.Check()

	var (
		versionId, _ = cmd.Flags().GetInt("id")
		version, _   = cmd.Flags().GetString("version")
		name, _      = cmd.Flags().GetString("name")
		appId, _     = cmd.Flags().GetString("app-id")
		stable, _    = cmd.Flags().GetBool("stable")
	)

	fmt.Println("Update Version", versionId)

	resp, err := serviceClient.UpdateVersion(versionId, version, name, appId, stable)

	if err != nil {
		fmt.Println("Error", err.Error())
	}

	fmt.Println("version: %+v", resp)
}

func getVersion(cmd *cobra.Command, args []string) {
	serviceClient.Check()

	versionId, _ := cmd.Flags().GetInt("id")

	fmt.Println("Get Version", versionId)

	resp, err := serviceClient.GetVersion(versionId)

	if err != nil {
		fmt.Println("Error", err.Error())
	}

	fmt.Println("version: %+v", resp)
}

func deleteVersion(cmd *cobra.Command, args []string) {
	serviceClient.Check()

	versionId, _ := cmd.Flags().GetInt("id")

	fmt.Println("Delete Version", versionId)

	resp, err := serviceClient.DeleteVersion(versionId)

	if err != nil {
		fmt.Println("Error", err.Error())
	}

	fmt.Println("version: %+v", resp)
}
