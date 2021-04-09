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

type jinchangFinancial struct {
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
	BuZhu string // 补助,
	RealPrice string // 实际支付
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


var financialChan = make(chan jinchangFinancial, 1)

func ReadFile2() string {

	go insertFinancial(financialChan)

	str := readFinancial()

	time.Sleep(5 * time.Second)

	return str
}


func get_financial_dir() string {

	cDir, _ := os.Getwd()

	return cDir + "/File/BoxFile/"
}

func readFinancial() string {
	
	k := 0

	err := filepath.Walk(get_financial_dir(), func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if fi.IsDir() { // 忽略目录
			return nil
		}

		k++
		fmt.Println("filename:",filename)

		openCsvFile2(filename)
		
		//suffix := strings.ToUpper(".dat")
		//dates="20201020"
		//if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {   //判断字符串是不是以.dat结尾
		//	if strings.Contains(fi.Name(),dates){     //判断路径字符串是不是含有dates
		//		fmt.Println("###",fi.Name())
		//		//修改文件权限
		//		if err := os.Chmod(filename, 0777); err != nil {
		//			fmt.Println("11",err)
		//		}

		//		 DATtoCSV(filename)   //获取文件路径之后去处理你的文件
		//
		//	}
		//}
		return nil
	})

	if err == nil {
		return ""
	}

	return err.Error()
}

func openCsvFile2(fileName string)  {

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
		myFinancial := jinchangFinancial{
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
			strings.TrimSpace(record[10]),
			strings.TrimSpace(record[11]),
		}
		if len(myFinancial.XiangHao) == 0{
			continue
		}

		financialChan <- myFinancial
	}
}