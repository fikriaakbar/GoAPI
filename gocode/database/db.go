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
	db, err := sql.Open("godror", "hasibuan/hasibuan1@xe")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into DATA_BARANG(NAMA_BARANG, JUMLAH_BARANG, TANGGAL_DIMASUKAN, TANGGAL_UPDATE) VALUES(:1, :2, :3, :4)") 
	if err != nil { 
    fmt.Println(err) 
    return 
	}

	res, err := db.Query("AIR JORDAN 4", 10, "NOW()", "NOW()")  
	if err != nil { 
    fmt.Println(".....Error Inserting data") 
    fmt.Println(err) 
    return  
	} 

	rowCnt := res.RowsAffected()
	fmt.Println(rowCnt, "rows inserted.")
}

