package database

import (
	"database/sql"
	"fmt"
	"net/http"

	//github gordor
	_ "github.com/godror/godror"
)

//SelectDbDate func
func SelectDbDate(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("godror", "hasibuan/hasibuan1@xe")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select nama_barang from data_barang where rownum=1")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {

		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)
}

//InsertDb func
func InsertDb(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("godror", "<your username>/<your password>@service_name")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("insert into DATA_BARANG(NAMA_BARANG, JUMLAH_BARANG, TANGGAL_DIMASUKAN, TANGGAL_UPDATE) VALUES('Air Jordan 7', 100, time.Now(), time.Now())")
	if err != nil {
		fmt.Println("Error insert to DB")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {

		rows.Scan(&thedate)
	}
	fmt.Printf("Finish Insert %s", thedate)
}
