package rest

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
)

type StoreFile struct{}

func (sf *StoreFile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if validateErrs, ok := sf.validate(r); !ok {
		resBody := adapters.ResponseBody{Errs: &validateErrs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	// --------------------------------------------------------------------------------------

	rFile, err := r.MultipartForm.File["file"][0].Open()
	errorkit.ErrorHandled(err)
	defer errorkit.ErrorHandled(rFile.Close())

	var body bytes.Buffer
	mpWriter := multipart.NewWriter(&body)
	mpWriterW, err := mpWriter.CreateFormFile("file", r.MultipartForm.File["file"][0].Filename)
	errorkit.ErrorHandled(err)

	_, err = io.Copy(mpWriterW, rFile)
	errorkit.ErrorHandled(err)

	errorkit.ErrorHandled(mpWriter.Close())

	// --------------------------------------------------------------------------------------

	client := http.Client{}

	url := os.Getenv("FILE_SERVICE_REST_ADDRESS") + "/file"

	request, err := http.NewRequest(http.MethodPost, url, &body)
	errorkit.ErrorHandled(err)

	header := make(http.Header)
	header.Set("Content-Type", mpWriter.FormDataContentType())

	request.Header = header

	response, err := client.Do(request)
	errorkit.ErrorHandled(err)

	// --------------------------------------------------------------------------------------

	resBody, err := ioutil.ReadAll(response.Body)
	errorkit.ErrorHandled(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	w.Write(resBody)

	/* var responseBody []byte
	n, err := response.Body.Read(responseBody)
	errorkit.ErrorHandled(err)
	log.Println("read file service reponse body : ", n)

	adapters.NewResponse(response.StatusCode, &adapters.ResponseBody{Data: responseBody}).WriteResponse(&w) */
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

// --------------------------------------------------------------------------------------

type RetrieveFile struct{}

func (rf *RetrieveFile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	client := http.Client{}

	fullPath := r.URL.Query().Get("full_path")
	reqURLRaw := os.Getenv("FILE_SERVICE_REST_ADDRESS") + "/file?full_path=" + fullPath
	reqURL, err := url.Parse(reqURLRaw)
	errorkit.ErrorHandled(err)

	// --------------------------------------------------------------------------------------

	req, err := http.NewRequest(http.MethodGet, reqURL.String(), r.Body)
	errorkit.ErrorHandled(err)

	res, err := client.Do(req)
	errorkit.ErrorHandled(err)

	_, err = io.Copy(w, res.Body)
	errorkit.ErrorHandled(err)
}
