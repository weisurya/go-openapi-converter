package modulegdrive

import (
	"context"
	"io/ioutil"
	"os"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type gDriveHandler struct {
	service *drive.Service
}

// NewGDriveHandler works as a handler to use Google Drive module
func NewGDriveHandler(ctx context.Context, credentialPath string) (GDriveRepository, error) {
	b, err := ioutil.ReadFile(credentialPath)
	if err != nil {
		return nil, err
	}

	opts := option.WithCredentialsJSON(b)

	driveService, err := drive.NewService(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &gDriveHandler{
		service: driveService,
	}, nil
}

// GDriveRepository list out available function
type GDriveRepository interface {
	CreateNewFolder(folderName, parentID string) (*drive.File, error)
	SetWriterPermission(fileID, email string) (*drive.Permission, error)
	UploadDocxFile(filePath, filename, parentID string) (*drive.File, error)
}

func (s gDriveHandler) CreateNewFolder(folderName, parentID string) (*drive.File, error) {
	newDirectory := &drive.File{
		Name:     folderName,
		MimeType: "application/vnd.google-apps.folder",
		Parents:  []string{parentID},
	}

	newDir, err := s.service.Files.Create(newDirectory).Do()
	if err != nil {
		return nil, err
	}

	return newDir, nil
}

func (s gDriveHandler) SetWriterPermission(fileID, email string) (*drive.Permission, error) {
	newPermission := &drive.Permission{
		Role:         "writer",
		Type:         "user",
		EmailAddress: email,
	}

	updatedFile, err := s.service.Permissions.Create(fileID, newPermission).Do()

	if err != nil {
		return nil, err
	}

	return updatedFile, nil
}

func (s gDriveHandler) UploadDocxFile(filePath, filename, parentID string) (*drive.File, error) {
	newFile := &drive.File{
		Name:     filename,
		MimeType: "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		Parents:  []string{parentID},
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	createdFile, err := s.service.Files.Create(newFile).Media(file).Do()
	if err != nil {
		return nil, err
	}

	return createdFile, nil
}
