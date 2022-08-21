package externals

import (
	"io"
	"net/http"
	"os"

	"github.com/ilhammhdd/go-toolkit/errorkit"
)

type RetrieveFile struct{}

func (rf *RetrieveFile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fullPath := r.URL.Query().Get("full_path")

	// ----------------------------------------------------------------------------

	file, err := os.Open(fullPath)
	errorkit.ErrorHandled(err)
	defer file.Close()

	_, err = io.Copy(w, file)
	errorkit.ErrorHandled(err)
}
