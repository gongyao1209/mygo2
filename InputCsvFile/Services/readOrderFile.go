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

type jinchangBox struct {
	JingChangDate string // 进场日期
	CarNumber string// 车号
	JingChangZhuangtai string// 进场状态
	TidanCode string// 提单号
	XiangHao string// 箱号
	ChuanMing string// 船名
	HangCi string// 航次
	XiangXin string// 箱型
	ChiCun string// 尺寸
	KongZhong string// 空重

	QianFeng string// 铅封/计划/车号
	JingChangYuanyin string// 进场原因
	JieFei string// 结费
	CheDui string// 车队
	CheDuiDianHua string// 车队电话
	JingChangCaoZuo string// 进场操作员
	BeiZhu string// 备注
	LuoXiangTiao string// 落箱条
	LaiYuan string// 来源
	SiJi string// 司机

	ChuanGongsi string // 船公司
	JingChangXiangKuang string // 进场箱况
	JingChangHao string // 堆场号
	ChangWei string // 场位
	JingChangYanXiangRen string // 进场验箱人
	WeixianPingJiBie string // 危险品级别
	UN string// UN号
	HuoMing string // 货名
	XieYiJiage string// 协议价格
	JinChuChangStatus string // 进出场状态

	ChaoGao string// 超高
	ChaoKuan string// 超宽
	XiangZhu string// 箱主
	JiDan string// 寄单
	ChaDianShijian string// 插电时间
	BaDianShijian string // 拔电时间
	isYunDi string // 是否运抵
	MaTouCode string// 码头代码
}

/**
alter table `box_202102` add `JingChangDate` varchar(64) DEFAULT '' COMMENT '进场日期';
alter table `box_202102` add `CarNumber` varchar(64) DEFAULT '' COMMENT '车号';
alter table `box_202102` add `JingChangZhuangtai` varchar(64) DEFAULT '' COMMENT '进场状态';
alter table `box_202102` add `TidanCode` varchar(64) DEFAULT '' COMMENT '提单号';
alter table `box_202102` add `XiangHao` varchar(64) DEFAULT '' COMMENT '箱号';
alter table `box_202102` add `ChuanMing` varchar(64) DEFAULT '' COMMENT '船名';
alter table `box_202102` add `HangCi` varchar(64) DEFAULT '' COMMENT '航次';
alter table `box_202102` add `XiangXin` varchar(64) DEFAULT '' COMMENT '箱型';
alter table `box_202102` add `ChiCun` varchar(64) DEFAULT '' COMMENT '尺寸';
alter table `box_202102` add `KongZhong` varchar(64) DEFAULT '' COMMENT '空重';
alter table `box_202102` add `QianFeng` varchar(64) DEFAULT '' COMMENT '铅封/计划/车号';

alter table `box_202102` add `JingChangYuanyin` varchar(64) DEFAULT '' COMMENT '进场原因';
alter table `box_202102` add `JieFei` varchar(64) DEFAULT '' COMMENT '结费';
alter table `box_202102` add `CheDui` varchar(64) DEFAULT '' COMMENT '车队';
alter table `box_202102` add `CheDuiDianHua` varchar(64) DEFAULT '' COMMENT '车队电话';
alter table `box_202102` add `JingChangCaoZuo` varchar(64) DEFAULT '' COMMENT '进场操作员';
alter table `box_202102` add `BeiZhu` varchar(64) DEFAULT '' COMMENT '备注';
alter table `box_202102` add `LuoXiangTiao` varchar(64) DEFAULT '' COMMENT '落箱条';
alter table `box_202102` add `LaiYuan` varchar(64) DEFAULT '' COMMENT '来源';
alter table `box_202102` add `SiJi` varchar(64) DEFAULT '' COMMENT '司机';
alter table `box_202102` add `ChuanGongsi` varchar(64) DEFAULT '' COMMENT '船公司';

alter table `box_202102` add `JingChangXiangKuang` varchar(64) DEFAULT '' COMMENT '进场箱况';
alter table `box_202102` add `JingChangHao` varchar(64) DEFAULT '' COMMENT '堆场号';
alter table `box_202102` add `ChangWei` varchar(64) DEFAULT '' COMMENT '场位';
alter table `box_202102` add `JingChangYanXiangRen` varchar(64) DEFAULT '' COMMENT '进场验箱人';
alter table `box_202102` add `WeixianPingJiBie` varchar(64) DEFAULT '' COMMENT '危险品级别';
alter table `box_202102` add `UN` varchar(64) DEFAULT '' COMMENT 'UN号';
alter table `box_202102` add `HuoMing` varchar(64) DEFAULT '' COMMENT '货名';
alter table `box_202102` add `XieYiJiage` varchar(64) DEFAULT '' COMMENT '协议价格';
alter table `box_202102` add `JinChuChangStatus` varchar(64) DEFAULT '' COMMENT '进出场状态';
alter table `box_202102` add `ChaoGao` varchar(64) DEFAULT '' COMMENT '超高';

alter table `box_202102` add `ChaoKuan` varchar(64) DEFAULT '' COMMENT '超宽';
alter table `box_202102` add `XiangZhu` varchar(64) DEFAULT '' COMMENT '箱主';
alter table `box_202102` add `JiDan` varchar(64) DEFAULT '' COMMENT '寄单';
alter table `box_202102` add `ChaDianShijian` varchar(64) DEFAULT '' COMMENT '插电时间';
alter table `box_202102` add `BaDianShijian` varchar(64) DEFAULT '' COMMENT '拔电时间';
alter table `box_202102` add `isYunDi` varchar(64) DEFAULT '' COMMENT '是否运抵';
alter table `box_202102` add `MaTouCode` varchar(64) DEFAULT '' COMMENT '码头代码';
 */


var number = 0
var boxChan = make(chan jinchangBox, 1)

func ReadFile() string {

	go insertOrder(boxChan)

	str := read()

	fmt.Println("last ", number)

	time.Sleep(6 * time.Second)

	return str
}


func get_dir() string {

	cDir, _ := os.Getwd()

	return cDir + "/File/202104/order/"
}

func read() string {
	
	k := 0

	err := filepath.Walk(get_dir(), func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if fi.IsDir() { // 忽略目录
			return nil
		}

		k++
		fmt.Println("filename:",filename)

		openCsvFile(filename)
		
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

func openCsvFile(fileName string)  {

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

		hang += 1
		if hang == 1 { // 第一行的数据删除
			continue
		}

		// 将数据变为对象
		myBox := jinchangBox{
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
			strings.TrimSpace(record[12]),
			strings.TrimSpace(record[13]),
			strings.TrimSpace(record[14]),
			strings.TrimSpace(record[15]),
			strings.TrimSpace(record[16]),
			strings.TrimSpace(record[17]),
			strings.TrimSpace(record[18]),
			strings.TrimSpace(record[19]),
			strings.TrimSpace(record[20]),
			strings.TrimSpace(record[21]),
			strings.TrimSpace(record[22]),
			strings.TrimSpace(record[23]),
			strings.TrimSpace(record[24]),
			strings.TrimSpace(record[25]),
			strings.TrimSpace(record[26]),
			strings.TrimSpace(record[27]),
			strings.TrimSpace(record[28]),
			strings.TrimSpace(record[29]),
			strings.TrimSpace(record[30]),
			strings.TrimSpace(record[31]),
			strings.TrimSpace(record[32]),
			strings.TrimSpace(record[33]),
			strings.TrimSpace(record[34]),
			strings.TrimSpace(record[35]),
			strings.TrimSpace(record[36]),
			strings.TrimSpace(record[37]),
		}

		if len(myBox.TidanCode) == 0 {
			continue
		}

		boxChan <- myBox

		//fmt.Println(myBox)
	}
}