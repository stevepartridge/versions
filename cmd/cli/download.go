package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/cobra"
	// "golang.org/x/net/context"
	// "github.com/stevepartridge/versions"
	"github.com/stevepartridge/go/file"
	pb "github.com/stevepartridge/versions/protos"
)

func downloadCommands(cmd *cobra.Command) {

	dlCommands := &cobra.Command{
		Use:     "download",
		Short:   "Manage downloads",
		Aliases: []string{"app"},
	}
	dlCommands.PersistentFlags().IntP("version-id", "v", 0, "Version ID")

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get download",
		Long:  `The download to get`,
		Run:   getDownload,
	}

	dlCommands.AddCommand(getCmd)

	// Create Flags
	createCmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"add"},
		Short:   "Add new download",
		Long:    `The download to get`,
		Run:     createDownload,
	}

	createCmd.PersistentFlags().StringP("file", "f", "", "File to add as download")
	createCmd.Flags().StringP("storage", "s", "local", "Storage type to use")
	createCmd.Flags().StringP("filename", "", "", "Resulting filename to use when downloaded")
	createCmd.Flags().StringP("os", "", "", "Specific OS download is intended for")
	createCmd.Flags().StringP("arch", "a", "", "Specific Arch download is intended for")

	dlCommands.AddCommand(createCmd)

	getFileCmd := &cobra.Command{
		Use:   "get-file",
		Short: "Get download file",
		Long:  `The download file to get`,
		Run:   getDownloadFile,
	}

	getFileCmd.Flags().StringP("outfile", "o", "", "Override the default file name to save output file as")

	dlCommands.AddCommand(getFileCmd)

	cmd.AddCommand(dlCommands)

}

func getDownload(cmd *cobra.Command, args []string) {
	serviceClient.Check()

	downloadId, _ := cmd.Flags().GetInt("id")

	fmt.Printf("Get Download by ID: %d\n", downloadId)

	resp, err := serviceClient.GetDownload(downloadId)

	if err != nil {
		fmt.Println("Error", err.Error())
	}

	fmt.Printf(`  Download
    ID:       %d
    Version:  %d
    Filename: %s
    Updated:  %s
    Created:  %s
`,
		resp.Download.Id,
		resp.Download.VersionId,
		resp.Download.Filename,
		resp.Download.UpdatedAt,
		resp.Download.CreatedAt,
	)
}

func getDownloadFile(cmd *cobra.Command, args []string) {
	serviceClient.Check()

	downloadId, _ := cmd.Flags().GetInt("id")
	outfile, _ := cmd.Flags().GetString("outfile")

	fmt.Printf("Get File by Download ID: %d\n", downloadId)

	resp, err := serviceClient.GetDownloadFile(downloadId)

	if err != nil {
		fmt.Println("Error", err.Error())
	}

	if outfile == "" {
		outfile = resp.Filename
	}

	err = ioutil.WriteFile(outfile, resp.Data, 0644)
	if isError(err) {
		return
	}

	// 	fmt.Printf(`  Download
	//     ID:       %d
	//     Version:  %d
	//     Filename: %s
	//     Updated:  %s
	//     Created:  %s
	// `,
	// 		resp.Download.Id,
	// 		resp.Download.VersionId,
	// 		resp.Download.Filename,
	// 		resp.Download.UpdatedAt,
	// 		resp.Download.CreatedAt,
	// 	)
}

func createDownload(cmd *cobra.Command, args []string) {

	downloadId, err := cmd.Flags().GetInt("id")
	if isError(err) {
		return
	}
	versionId, err := cmd.Flags().GetInt("version-id")
	if isError(err) {
		return
	}

	f, err := cmd.Flags().GetString("file")
	if isError(err) {
		return
	}

	if !file.Exists(f) {
		isError(ErrDownloadCreateFileNotFound)
		return
	}

	storage, _ := cmd.Flags().GetString("storage")
	filename, _ := cmd.Flags().GetString("filename")
	os, _ := cmd.Flags().GetString("os")
	arch, _ := cmd.Flags().GetString("arch")

	if filename == "" {
		filename = filepath.Base(f)
	}

	data, err := ioutil.ReadFile(f)
	if isError(err) {
		return
	}

	serviceClient.Check()

	download := &pb.Download{
		VersionId:   int32(versionId),
		StorageType: storage,
		Filename:    filename,
		Os:          os,
		Arch:        arch,
		Data:        data,
	}

	fmt.Printf("Create Download for version: %d File: %s\n", downloadId, filename)

	resp, err := serviceClient.CreateDownload(versionId, download)

	if isError(err) {
		return
	}

	// fmt.Println("application: %+v", resp)
	fmt.Printf(`  Download
    ID:       %d
    Version:  %d
    Filename: %s
    Updated:  %s
    Created:  %s
`,
		resp.Download.Id,
		resp.Download.VersionId,
		resp.Download.Filename,
		resp.Download.UpdatedAt,
		resp.Download.CreatedAt,
	)
}
