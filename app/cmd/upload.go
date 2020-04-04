package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	gdrivemodule "github.com/weisurya/go-openapi-converter/modules/gdrive"
)

var (
	uploadCmd = &cobra.Command{
		Use:     "upload",
		Short:   "Upload file to Google Drive",
		Long:    fmt.Sprintf("Upload file to Google Drive.\nIn order to use this feature, you need to have OAuth credential that connected to your Google Cloud Platform. Please follow the tutorial on README.md first."),
		Run:     uploadToGoogleDrive,
		Example: "./go-openapi-converter upload -c credentials.sample.json -d directory -e foo@gmail.com -f result.sample.docx",
	}
)

const (
	defaultCredentialPath = "./credentials.sample.json"
	defaultEmpty          = ""
)

var (
	credentialPath string
	directoryName  string
	directoryID    string
	email          string
	file           string
)

func init() {
	uploadCmd.Flags().StringVarP(&credentialPath, "credential", "c", defaultCredentialPath, "OAuth Credential path in JSON format")
	uploadCmd.MarkFlagRequired("credential")

	uploadCmd.Flags().StringVarP(&directoryName, "dirname", "d", defaultEmpty, "Directory name in Google Drive. It will create a new folder with this name. This flag is mandatory IF you do not provide directory ID")

	uploadCmd.Flags().StringVarP(&directoryID, "id", "i", defaultEmpty, "Directory ID in Google Drive. It will upload the file into this folder. (https://drive.google.com/drive/folders/<directory_id>)")

	uploadCmd.Flags().StringVarP(&email, "email", "e", defaultCredentialPath, "Google email address")
	uploadCmd.MarkFlagRequired("email")

	uploadCmd.Flags().StringVarP(&file, "file", "f", defaultEmpty, "File to be uploaded")
	uploadCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(uploadCmd)
}

func uploadToGoogleDrive(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	if directoryID == "" && directoryName == "" {
		fmt.Println("Please provide directory name OR directory ID")
		os.Exit(1)
	}

	driveService, err := gdrivemodule.NewGDriveHandler(ctx, credentialPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var id string
	if directoryID == "" {
		newDir, err := driveService.CreateNewFolder(directoryName, "root")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		id = newDir.Id

		_, err = driveService.SetWriterPermission(id, email)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		id = directoryID
	}

	_, filename := filepath.Split(file)

	newFile, err := driveService.UploadDocxFile(file, filename, id)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = driveService.SetWriterPermission(newFile.Id, email)

	return
}
