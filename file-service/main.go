package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ilhammhdd/kudaki-file-service/externals"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/ilhammhdd/kudaki-file-service/externals/mysql"
)

func init() {
	if len(os.Args) > 1 {
		for _, val := range os.Args[1:] {
			f := strings.Split(val, " ")
			os.Setenv(f[1], f[2])
		}
	}

	mysql.CommandDB = mysql.OpenDB(os.Getenv("DB_PATH"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	mysql.QueryDB = mysql.OpenDB(os.Getenv("QUERY_DB_PATH"), os.Getenv("QUERY_DB_USERNAME"), os.Getenv("QUERY_DB_PASSWORD"), os.Getenv("QUERY_DB_NAME"))
}

func main() {
	http.Handle("/file", externals.MethodRouting{
		PostHandler: new(externals.StoreFile),
		GetHandler:  new(externals.RetrieveFile)})

	fmt.Println("on")

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("REST_PORT"))}
	defer server.Close()

	errorkit.ErrorHandled(server.ListenAndServe())
}
