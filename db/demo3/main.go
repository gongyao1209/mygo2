package main

import (
	sql2 "database/sql"
	"fmt"
	db2 "gocode/mygo2/db"
	"sync"
)

// 测试事务
// 1、go中事务的使用，包括创建 事务 关闭事务之后事务的使用
// 2、go程 中 的事务。主要是考虑 事务关闭之后，还在用的事务
func main()  {

	var wg sync.WaitGroup

	d := db2.GetDB()

	tx, err := d.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}


	sql := `UPDATE student SET Sdept = ? WHERE Sno = ?`
	var r sql2.Result
	r, err = tx.Exec(sql, "计算机科学与技术11", 5)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		id1, _ := r.LastInsertId() // 最后插入的值
		id2, _ := r.RowsAffected() // 影响了多少行
		fmt.Println("InsertId: ", id1, ", Affected:", id2)
	}

	wg.Add(1)
	go func(tx *sql2.Tx) {
		defer wg.Done()

		//time.Sleep(1 * time.Second)
		var s stu
		// 事务中获取数据，而且还要判断事务 是否存在
		sql = "SELECT * FROM student WHERE Sno = 5"
		err := tx.QueryRow(sql).Scan(&s.Sno, &s.Sname, &s.Ssex, &s.Sage, &s.Sdept)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(s)
		}
	}(tx)
	wg.Wait()

	tx.Commit()
}

type stu struct {
	Sno int //`Sno` int(10) NOT NULL COMMENT '学号',
	Sname string //`Sname` varchar(16) NOT NULL COMMENT '姓名',
	Ssex string //`Ssex` char(2) NOT NULL COMMENT '姓别',
	Sage int //`Sage` tinyint(2) NOT NULL DEFAULT '0' COMMENT '学生年龄',
	Sdept string //`Sdept` varchar(16) DEFAULT NULL COMMENT '学生所在系别',
}