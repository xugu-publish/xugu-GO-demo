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
	"log"
	"time"
)

func main() {

	db, _ := sql.Open("xugusql", "IP=127.0.0.1;"+
		"DB=SYSTEM;User=SYSDBA;PWD=SYSDBA;"+
		"Port=5138;AUTO_COMMIT=on;CHAR_SET=UTF8")

	res, err := db.Exec("insert into go_3th_test values(?,?);", 1, time.Now())
	if err != nil {

		log.Fatal(err)
	}

	effec, _ := res.RowsAffected()
	fmt.Printf("Number of rows affected: %d\n", effec)

	err = db.Close()
	if err != nil {

		log.Fatal(err)
	}

}
