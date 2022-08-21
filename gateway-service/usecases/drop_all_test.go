package usecases_test

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ilhammhdd/go-toolkit/errorkit"
)

func TestTruncateAllTable(t *testing.T) {
	type truncateDB struct {
		truncateQueries []string
		dbName          string
	}

	tqRental := []string{
		"use kudaki_rental;",
		"set FOREIGN_KEY_CHECKS=0;",
		"truncate table carts;",
		"truncate table cart_items;",
		"truncate table checkouts;",
		"set FOREIGN_KEY_CHECKS=1;"}

	tqStore := []string{
		"use kudaki_store;",
		"set FOREIGN_KEY_CHECKS=0;",
		"truncate table storefronts;",
		"truncate table items;",
		"set FOREIGN_KEY_CHECKS=1;"}

	// tqUser := []string{
	// 	"use kudaki_user;",
	// 	"set FOREIGN_KEY_CHECKS=0;",
	// 	"truncate table users;",
	// 	"truncate table profiles;",
	// 	"truncate table unverified_users;",
	// 	"truncate table reset_passwords;",
	// 	"set FOREIGN_KEY_CHECKS=1;"}

	truncateDBs := []truncateDB{
		{truncateQueries: tqRental, dbName: "kudaki_rental"},
		{truncateQueries: tqStore, dbName: "kudaki_store"},
		/* {truncateQueries: tqUser, dbName: "kudaki_user"} */}

	for _, tDB := range truncateDBs {
		dbDataSource := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", "root", "mysqlrocks", "tcp(178.62.107.160:3307)", tDB.dbName)
		db, err := sql.Open("mysql", dbDataSource)
		errorkit.ErrorHandled(err)
		defer db.Close()

		for _, query := range tDB.truncateQueries {
			t.Log(query)
			_, err := db.Exec(query)
			errorkit.ErrorHandled(err)
		}
	}
}
