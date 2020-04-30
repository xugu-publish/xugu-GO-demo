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

var Ques = []string{
	"insert into go_1th_test values(1, 'AJX', 0.99);",
	"insert into go_1th_test values(1, 'BJX', 0.99);",
	"insert into go_1th_test values(1, 'CJX', 0.99);",

	"update go_1th_test set c1=1 where c1=1;",
}

func main() {

	db, _ := sql.Open("xugusql", "IP=127.0.0.1;"+
		"DB=SYSTEM;User=SYSDBA;PWD=SYSDBA;"+
		"Port=5138;AUTO_COMMIT=on;CHAR_SET=UTF8")

	for _, sql := range Ques {

		res, err := db.Exec(sql)
		if err != nil {

			log.Fatal(err)
			continue
		}
		fmt.Printf("%s ... ok\n", sql)

		effec, _ := res.RowsAffected()
		fmt.Printf("Number of rows affected: %d\n", effec)

	}

	err := db.Close()
	if err != nil {

		log.Fatal(err)
	}

}
