package main

import (
	"os"
	"strings"

	"github.com/ilhammhdd/kudaki-order-service/externals"

	"github.com/ilhammhdd/go-toolkit/safekit"

	"github.com/ilhammhdd/kudaki-order-service/externals/mysql"
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
	wp := safekit.NewWorkerPool()

	wp.Worker <- new(externals.CheckOut)
	wp.Worker <- new(externals.RetrieveOwnersOrderHistories)
	wp.Worker <- new(externals.RetrieveTenantOrderHistories)
	wp.Worker <- new(externals.ApproveOwnerOrder)
	wp.Worker <- new(externals.DisapproveOwnerOrder)
	wp.Worker <- new(externals.OwnerOrderApproved)
	wp.Worker <- new(externals.OwnerOrderDisapproved)
	wp.Worker <- new(externals.OwnerConfirmReturnment)
	wp.Worker <- new(externals.OwnerConfirmedReturnment)
	wp.Worker <- new(externals.TenantReviewsOwnerOrder)
	wp.Worker <- new(externals.OwnerOrderRented)
	wp.Worker <- new(externals.OwnerOrderRentedOut)

	wp.PoolWG.Wait()
}
