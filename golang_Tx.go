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

	db, _ := sql.Open("xugusql", "IP=127.0.0.1;"+
		"DB=SYSTEM;User=SYSDBA;PWD=SYSDBA;"+
		"Port=5138;AUTO_COMMIT=on;CHAR_SET=UTF8")

	Tx, err := db.Begin()
	if err != nil {

		log.Fatal(err)
	}

	_, err = Tx.Exec("insert into go_1th_test values(9,'9',0.99)")
	if err != nil {

		log.Fatal(err)
	}

	//err = Tx.Commit()
	err = Tx.Rollback()
	if err != nil {

		log.Fatal(err)
	}

	db.Close()

}
