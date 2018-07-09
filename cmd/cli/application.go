package main

import (
	"fmt"

	"github.com/spf13/cobra"
	// "golang.org/x/net/context"
	// "github.com/stevepartridge/versions"
	// pb "github.com/stevepartridge/versions/protos"
)

func applicationCommands(cmd *cobra.Command) {

	appCommands := &cobra.Command{
		Use:     "application",
		Short:   "Manage applications",
		Aliases: []string{"app"},
	}
	appCommands.PersistentFlags().StringP("id", "i", "", "Application ID")

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get application",
		Long:  `The application to get`,
		Run:   getApplication,
	}

	appCommands.AddCommand(getCmd)

	// Create Flags
	createCmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"add"},
		Short:   "Add new application",
		Long:    `The application to get`,
		Run:     createApplication,
	}

	// createCmd.PersistentFlags().StringP("id", "i", "", "Unique Application ID")
	createCmd.PersistentFlags().StringP("app-name", "n", "", "Application Name")
	// createCmd.PersistentFlags().SetNormalizeFunc(n)
	// createCmd.PersistentFlags().String

	appCommands.AddCommand(createCmd)

	cmd.AddCommand(appCommands)

}

func getApplication(cmd *cobra.Command, args []string) {
	serviceClient.Check()

	appId := cmd.Flag("id").Value.String()

	fmt.Printf("Get Application by ID: %s\n", appId)

	resp, err := serviceClient.GetApplication(appId)

	if err != nil {
		fmt.Println("Error", err.Error())
	}

	fmt.Printf(`  
  Application
    ID:      %s
    Name:    %s
    Updated: %s
    Created: %s

`,
		resp.Application.Id,
		resp.Application.Name,
		resp.Application.UpdatedAt,
		resp.Application.CreatedAt,
	)
}

func createApplication(cmd *cobra.Command, args []string) {

	appId := cmd.Flag("id").Value.String()
	appName, err := cmd.Flags().GetString("app-name")

	// TODO: this is so hacky :(
	//       figure out how to parse string with spaces
	if len(args) > 0 {
		for i := range args {
			appName = fmt.Sprintf("%s %s", appName, args[i])
		}
	}

	serviceClient.Check()

	fmt.Printf("Create Application with ID: %s Name: %s\n", appId, appName)

	resp, err := serviceClient.CreateApplication(appId, appId)

	if isError(err) {
		return
	}

	// fmt.Println("application: %+v", resp)
	fmt.Printf(`  Application
    ID:      %s
    Name:    %s
    Updated: %s
    Created: %s
`,
		resp.Application.Id,
		resp.Application.Name,
		resp.Application.UpdatedAt,
		resp.Application.CreatedAt,
	)
}
