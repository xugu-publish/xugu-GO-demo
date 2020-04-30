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
	"log"
)

func main() {

	db, err := sql.Open("xugusql", "IP=127.0.0.1;"+
		"DB=SYSTEM;User=SYSDBA;PWD=SYSDBA;"+
		"Port=5138;AUTO_COMMIT=on;CHAR_SET=UTF8")

	// CREATE A TABLE
	_, err = db.Exec("create table go_1th_test(c1 int, c2 varchar, c3 float);")
	if err != nil {

		log.Fatal(err)
	}

	// CREATE A TABLE
	_, err = db.Exec("create table go_2th_test(c1 int, c2 blob, c3 clob);")
	if err != nil {

		log.Fatal(err)
	}

	// CREATE A TABLE
	_, err = db.Exec("create table go_3th_test(c1 int, c2 datetime);")
	if err != nil {

		log.Fatal(err)
	}

	err = db.Close()
	if err != nil {

		log.Fatal(err)
	}

}
