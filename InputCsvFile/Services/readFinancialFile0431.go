package Services

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type fromFinancial struct { // 箱号,箱型,尺寸,车号,进场日期,出场日期,封号,码头代码,进箱原因,金额
	XiangHao string //箱号,
	XiangXin string // 箱型,
	ChiCun string // 尺寸,
	CheHao string // 车号,
	JingChangDate string // 进场日期,
	ChuChangDate string // 出场日期,
	QianFenHao string // 封号,
	MaTouCode string // 码头代码,
	JinChangYuanyin string // 进箱原因,
	Price string // 金额,
}
/**
alter table `financial_202102` add `XiangHao` varchar(64) DEFAULT '' COMMENT '箱号';
alter table `financial_202102` add `XiangXin` varchar(64) DEFAULT '' COMMENT '箱型';
alter table `financial_202102` add `ChiCun` varchar(64) DEFAULT '' COMMENT '尺寸';
alter table `financial_202102` add `CheHao` varchar(64) DEFAULT '' COMMENT '车号';
alter table `financial_202102` add `JingChangDate` varchar(64) DEFAULT '' COMMENT '进场日期';
alter table `financial_202102` add `ChuChangDate` varchar(64) DEFAULT '' COMMENT '出场日期';
alter table `financial_202102` add `QianFenHao` varchar(64) DEFAULT '' COMMENT '封号';
alter table `financial_202102` add `MaTouCode` varchar(64) DEFAULT '' COMMENT '码头代码';
alter table `financial_202102` add `JinChangYuanyin` varchar(64) DEFAULT '' COMMENT '进箱原因';
alter table `financial_202102` add `Price` varchar(64) DEFAULT '' COMMENT '金额';
alter table `financial_202102` add `BuZhu` varchar(64) DEFAULT '' COMMENT '补助';
alter table `financial_202102` add `RealPrice` varchar(64) DEFAULT '' COMMENT '实际支付';

 */

var from_chan = make(chan fromFinancial, 1)

func ReadFile3() string {

	go insert_from_financial(from_chan)

	str := read_from_financial()

	time.Sleep(5 * time.Second)

	return str
}


func get_from_financial_dir() string {

	cDir, _ := os.Getwd()

	str := cDir + "/File/FinancialFile/"
	fmt.Println(str)

	return str
}

func read_from_financial() string {
	
	k := 0

	err := filepath.Walk(get_from_financial_dir(), func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if fi.IsDir() { // 忽略目录
			return nil
		}

		k++
		fmt.Println("filename:",filename)

		openCsvFile3(filename)

		return nil
	})

	if err == nil {
		return ""
	}

	return err.Error()
}

func openCsvFile3(fileName string)  {

	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var hang = 0

	for true {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return
		}

		number += 1
		hang += 1
		if hang == 1 { // 第一行的数据删除
			continue
		}

		// 将数据变为对象
		myFinancial := fromFinancial{
			strings.TrimSpace(record[0]),
			strings.TrimSpace(record[1]),
			strings.TrimSpace(record[2]),
			strings.TrimSpace(record[3]),
			strings.TrimSpace(record[4]),
			strings.TrimSpace(record[5]),
			strings.TrimSpace(record[6]),
			strings.TrimSpace(record[7]),
			strings.TrimSpace(record[8]),
			strings.TrimSpace(record[9]),
		}
		if len(myFinancial.XiangHao) == 0{
			continue
		}

		from_chan <- myFinancial
	}
}


func insert_from_financial(ch chan fromFinancial)  {

pool:
	for true {
		select {
		case fin :=<- ch :
			// 插入数据库数据
			insert_db_from_financial(fin)

			//fmt.Println(fin)
			break
		case <- time.After(2 * time.Second):
			fmt.Println("financial exit")
			close(ch)
			break pool
		}
	}

	fmt.Println("financial success")

	return
}

func insert_db_from_financial(item fromFinancial)  {
	// 获取数据库
	db := GetDB()

	// 插入数据
	sql := `INSERT INTO from_financial_202102 (XiangHao, XiangXin, ChiCun, CheHao, JingChangDate, ChuChangDate, QianFenHao, MaTouCode, JinChangYuanyin, Price) VALUES (?,?,?,?,?,?,?,?,?,?)`

	bind := make([]interface{}, 0)

	bind = append(bind, item.XiangHao)
	bind = append(bind, item.XiangXin)
	bind = append(bind, item.ChiCun)
	bind = append(bind, item.CheHao)
	bind = append(bind, item.JingChangDate)
	bind = append(bind, item.ChuChangDate)
	bind = append(bind, item.QianFenHao)
	bind = append(bind, item.MaTouCode)
	bind = append(bind, item.JinChangYuanyin)
	bind = append(bind, item.Price)

	db.Exec(sql, bind...);
}