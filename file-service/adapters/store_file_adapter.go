package adapters

import (
	"mime/multipart"
	"net/http"

	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-file-service/entities"
)

type StoreFile struct {
	R *http.Request
}

func (sf *StoreFile) GetFile() *entities.File {
	f := entities.File{
		Name:     sf.R.MultipartForm.File["file"][0].Filename,
		UUIDName: uuid.New().String(),
		Size:     sf.R.MultipartForm.File["file"][0].Size,
		MimeType: sf.R.MultipartForm.File["file"][0].Header.Get("Content-Type")}

	return &f
}

func (sf *StoreFile) GetMultipartFile() *multipart.File {
	mfFile, err := sf.R.MultipartForm.File["file"][0].Open()
	errorkit.ErrorHandled(err)

	return &mfFile
}
