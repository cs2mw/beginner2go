package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/world")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT Name,CountryCode,District,Population FROM city WHERE CountryCode = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var name, countryCode, district string
	var population int // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow("AFG").Scan(&name, &countryCode, &district, &population) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The first city which country code is AFG is: %s,%s,%s,%d \n", name, countryCode, district, population)

	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow("NLD").Scan(&name, &countryCode, &district, &population) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The other city which country code is NLD is: %s,%s,%s,%d \n", name, countryCode, district, population)

	// Query all results
	rows, err := stmtOut.Query("DZA")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&name, &countryCode, &district, &population)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("The city which country code is DZA contains: %s,%s,%s,%d \n", name, countryCode, district, population)
	}
}
