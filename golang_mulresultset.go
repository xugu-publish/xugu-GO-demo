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
)

func main() {

	db, _ := sql.Open("xugusql", "IP=127.0.0.1;"+
		"DB=SYSTEM;User=SYSDBA;PWD=SYSDBA;"+
		"Port=5138;AUTO_COMMIT=on;CHAR_SET=UTF8")

	rows, err := db.Query("select * from go_1th_test;" +
		"select * from go_3th_test;")
	if err != nil {

		log.Fatal(err)
	}

	set := 1
	for true {

		// Before obtaining each result set, you need
		// to first obtain the result set
		// column number information
		var cols []string
		cols, err = rows.Columns()
		if err != nil {

			log.Fatal(err)
		}

		pvals := make([]interface{}, len(cols))
		for key, _ := range pvals {

			dest := make([]byte, 216)
			pvals[key] = &dest
		}

		fmt.Printf("%dth Result Set:\n", set)

		for rows.Next() {
			err = rows.Scan(pvals...)
			if err != nil {

				log.Fatal(err)
			}

			for _, v := range pvals {

				fmt.Printf("%s\t", string(*(v.(*[]byte))))
			}
			fmt.Printf("\n")
		}

		if !rows.NextResultSet() {
			break
		}

		set++
	}

	rows.Close()

	err = db.Close()
	if err != nil {

		log.Fatal(err)
	}

}
