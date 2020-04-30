package main

/*
   +---------------------------------------+
   |    go-driver-xugusql                  |
   +---------------------------------------+
   |  Date: 2020-04-30 |
   +---------------------------------------+
*/

import (
	_ "./go-driver-xugusql"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {

	db, _ := sql.Open("xugusql", "IP=127.0.0.1;"+
		"DB=SYSTEM;User=SYSDBA;PWD=SYSDBA;"+
		"Port=5138;AUTO_COMMIT=on;CHAR_SET=UTF8")

	fp, err := os.Open("./LOB/CLOB.txt")
	if err != nil {

		log.Fatal(err)
	}

	var contents []byte
	contents, err = ioutil.ReadAll(fp)
	if err != nil {

		log.Fatal(err)
	}
	fp.Close()

	// INSERT LARGE OBJECT DATA INTO THE DATABASE
	_, err = db.Exec("insert into go_2th_test values(1,null,?);", string(contents))
	if err != nil {

		log.Fatal(err)
	}

	// QUERY LARGE OBJECT DATA FROM THE DATABSE
	var rows *sql.Rows
	rows, err = db.Query("select len(c3) from go_2th_test where c1=1;")
	if err != nil {

		log.Fatal(err)
	}

	dest := make([]byte, 256)
	var pval interface{} = &dest

	// Get row result set
	if rows.Next() {

		err = rows.Scan(pval)
		if err != nil {

			log.Fatal(err)
		}
	}

	rows.Close()

	// First get the length of the large object
	rows, err = db.Query("select c3 from go_2th_test where c1=1;")
	if err != nil {

		log.Fatal(err)
	}

	// Manually allocate enough long memory space to store large object data
	length, _ := strconv.Atoi(string(*(pval.(*[]byte))))
	txt := make([]byte, length)
	var lob interface{} = &txt

	// Get row result set
	if rows.Next() {

		err = rows.Scan(lob)
		if err != nil {

			log.Fatal(err)
		}

	}

	fp, err = os.Create("./CLOB.txt")
	if err != nil {

		log.Fatal(err)
	}

	size, err := fp.WriteString(string(*(lob.(*[]byte))))
	if err != nil {

		log.Fatal(err)
	}
	fmt.Printf("Number of bytes written to file: %d\n\n", size)
	fp.Close()

	rows.Close()

	err = db.Close()
	if err != nil {

		log.Fatal(err)
	}

}
