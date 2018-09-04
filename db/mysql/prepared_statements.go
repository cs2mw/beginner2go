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
	stmtCount, err := db.Prepare("SELECT count(Name) FROM city WHERE CountryCode = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtCount.Close()

	stmtSum, err := db.Prepare("select sum(Population) from city where CountryCode=?")
	if err != nil {
		panic(err.Error)
	}
	defer stmtSum.Close()

	stmtOut, err := db.Prepare("SELECT Name,CountryCode,District,Population FROM city WHERE CountryCode = ? order by Name asc,District desc")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var name, countryCode, district string
	var population, rowsCount int // we "scan" the result in here

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
	fmt.Printf("The first city which country code is NLD is: %s,%s,%s,%d \n", name, countryCode, district, population)

	// Query all results
	err = stmtCount.QueryRow("DZA").Scan(&rowsCount)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The city which country code is DZA contains:%d \n", rowsCount)

	rows, err := stmtOut.Query("DZA")
	if err != nil {
		panic(err.Error())
	}
	count := 0
	for rows.Next() {
		err := rows.Scan(&name, &countryCode, &district, &population)
		if err != nil {
			panic(err.Error())
		}
		count++
		fmt.Printf("Row %d:%s,%s,%s,%d \n", count, name, countryCode, district, population)
	}

	var rowsSum int
	err = stmtSum.QueryRow("DZA").Scan(&rowsSum)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The city which country code is DZA contains:%d \n people.", rowsSum)
}
