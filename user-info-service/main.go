package main

import (
	"os"
	"strings"

	"github.com/ilhammhdd/kudaki-user-info-service/externals"

	"github.com/ilhammhdd/go-toolkit/safekit"

	"github.com/ilhammhdd/kudaki-user-info-service/externals/mysql"
)

func init() {
	if len(os.Args) > 1 {
		for _, val := range os.Args[1:] {
			f := strings.Split(val, " ")
			os.Setenv(string(f[1]), f[2])
		}
	}

	mysql.CommandDB = mysql.OpenDB(os.Getenv("DB_PATH"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	mysql.QueryDB = mysql.OpenDB(os.Getenv("QUERY_DB_PATH"), os.Getenv("QUERY_DB_USERNAME"), os.Getenv("QUERY_DB_PASSWORD"), os.Getenv("QUERY_DB_NAME"))
}

func main() {
	wp := safekit.NewWorkerPool()

	wp.Worker <- new(externals.AddAddress)
	wp.Worker <- new(externals.UpdateAddress)
	wp.Worker <- new(externals.UpdateProfile)
	wp.Worker <- new(externals.RetrieveAddresses)
	wp.Worker <- new(externals.RetrieveProfile)

	wp.PoolWG.Wait()
}
