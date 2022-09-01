package main

import (
	"fmt"
	mw "mware"
)

func main() {
	dbs := mw.DbStore{}
	dbs.New()

	statusCode := dbs.Query("INSERT INTO sr VALUES ( 2, 'asset_pn' )")
	if statusCode == 0 {
		fmt.Printf("Success in Query")
	} else {
		fmt.Printf("Error in Query")
	}
	dbs.Close()
}
