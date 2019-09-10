package main

import (
	"fmt"
	db2 "gocode/mygo2/db"
)

func main()  {

	db := db2.GetDB()

	sql := `SELECT * FROM student AS s WHERE TRUE`

	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	test := make([]student, 0, 10)
	for rows.Next() {

		var s student
		err = rows.Scan(&s.Sno, &s.Sname, &s.Ssex, &s.Sage, &s.Sdept)
		if err == nil {
			test = append(test, s)
		}
	}

	fmt.Println(test)
}

type student struct {
	Sno int
	Sname string
	Ssex string
	Sage int
	Sdept string
}