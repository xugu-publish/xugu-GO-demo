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
	"time"
)

func main() {

	db, _ := sql.Open("xugusql", "IP=127.0.0.1;"+
		"DB=SYSTEM;User=SYSDBA;PWD=SYSDBA;"+
		"Port=5138;AUTO_COMMIT=on;CHAR_SET=UTF8")

	// PREPARE STMT
	stmt, err := db.Prepare("insert into go_3th_test values(?,?);")
	if err != nil {

		log.Fatal(err)
	}

	pos := 1
	for pos <= 10 {

		_, err = stmt.Exec(pos, time.Now())
		if err != nil {

			log.Fatal(err)
		}

		pos++
	}

	stmt.Close()

	err = db.Close()
	if err != nil {

		log.Fatal(err)
	}

}
