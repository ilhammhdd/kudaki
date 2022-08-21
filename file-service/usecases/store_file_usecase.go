package usecases

import (
	"io"
	"mime/multipart"
	"net/url"
	"os"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-file-service/entities"
)

type StoreFile struct {
	File   *entities.File
	MFFile *multipart.File
}

func (sf *StoreFile) WriteMultipartFileToFolder(folderPath string) error {
	defer (*sf.MFFile).Close()

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		errorkit.ErrorHandled(os.Mkdir(folderPath, 0755))
	}

	fullPath := folderPath + string(os.PathSeparator) + sf.File.UUIDName
	sf.File.FullPath = fullPath
	defer func() {
		sf.File.FullPath = os.Getenv("GATEWAY_SERVICE_REST_ADDRESS") + "/file?full_path=" + url.PathEscape(sf.File.FullPath)
	}()

	file, err := os.Create(fullPath)
	errorkit.ErrorHandled(err)
	defer file.Close()

	_, err = io.Copy(file, *sf.MFFile)
	errorkit.ErrorHandled(err)

	return file.Sync()
}
