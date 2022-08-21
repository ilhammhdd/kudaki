package externals

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-file-service/adapters"

	"github.com/ilhammhdd/kudaki-file-service/externals/mysql"

	"github.com/ilhammhdd/kudaki-file-service/usecases"

	"github.com/google/uuid"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-file-service/entities"
)

type StoreFile struct{}

func (sf *StoreFile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if validateErrs, ok := sf.validate(r); !ok {
		resBody := adapters.ResponseBody{Errs: &validateErrs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.StoreFile{R: r}

	f := adapter.GetFile()
	usecase := usecases.StoreFile{File: f, MFFile: adapter.GetMultipartFile()}
	err := usecase.WriteMultipartFileToFolder("files")
	errorkit.ErrorHandled(err)

	sf.insertMetaToDB(f)

	resBody := adapters.ResponseBody{
		Data: map[string]interface{}{
			"file_name": f.Name,
			"full_path": f.FullPath,
			"mime_type": f.MimeType,
			"file_size": f.Size}}
	adapters.NewResponse(http.StatusOK, &resBody).WriteResponse(&w)
}

func (sf *StoreFile) insertMetaToDB(file *entities.File) {
	dbo := mysql.NewDBOperation(mysql.CommandDB)

	_, err := dbo.Command("INSERT INTO kudaki_file.kudaki_files(uuid,uuid_name,original_name,size,mime_type,full_path,created_at) VALUES(?,?,?,?,?,?,UNIX_TIMESTAMP());",
		uuid.New().String(), file.UUIDName, file.Name, file.Size, file.MimeType, file.FullPath)
	errorkit.ErrorHandled(err)
}

func (sf *StoreFile) validate(r *http.Request) ([]string, bool) {
	errorkit.ErrorHandled(r.ParseMultipartForm(50000000))

	ok := true
	var errs []string
	if len(r.MultipartForm.File["file"]) == 0 {
		errs = append(errs, "empty \"file\" in multipart form")
		ok = false
	}

	return errs, ok
}
