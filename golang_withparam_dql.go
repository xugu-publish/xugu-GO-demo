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

	rows, err := db.Query("select * from go_1th_test where c1=?;", 1)
	if err != nil {
		log.Fatal(err)
	}

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

	/*
	   Get result set line by line
	*/
	for rows.Next() {
		err = rows.Scan(pvals...)
		if err != nil {

			log.Fatal(err)
			break
		}

		for _, v := range pvals {

			fmt.Printf("%s\t", string(*(v.(*[]byte))))
		}
		fmt.Printf("\n")
	}

	rows.Close()

	err = db.Close()
	if err != nil {

		log.Fatal(err)
	}

}
