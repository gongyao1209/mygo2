package Services

import (
	"fmt"
	"time"
)

func insertFinancial(ch chan jinchangFinancial)  {

	pool:
	for true {
		select {
		case fin :=<- ch :
			// 插入数据库数据
			insert_db_financial(fin)

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

func insert_db_financial(item jinchangFinancial)  {
	// 获取数据库
	db := GetDB()

	// 插入数据
	sql := `INSERT INTO financial_202102 (XiangHao, XiangXin, ChiCun, CheHao, JingChangDate, ChuChangDate, QianFenHao, MaTouCode, JinChangYuanyin, Price, BuZhu, RealPrice) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`

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
	bind = append(bind, item.BuZhu)
	bind = append(bind, item.RealPrice)

	db.Exec(sql, bind...);
}

func insert_db_box(item jinchangBox)  {
	// 获取数据库
	db := GetDB()

	// 插入数据
	sql := `INSERT INTO m202104_box (
                        JingChangDate, CarNumber, JingChangZhuangtai, TidanCode, XiangHao, ChuanMing, HangCi, XiangXin, ChiCun, KongZhong, QianFeng,
                        JingChangYuanyin, JieFei, CheDui, CheDuiDianHua, JingChangCaoZuo, BeiZhu, LuoXiangTiao, LaiYuan, SiJi, ChuanGongsi,
                        JingChangXiangKuang, JingChangHao, ChangWei, JingChangYanXiangRen, WeixianPingJiBie, UN, HuoMing, XieYiJiage, JinChuChangStatus, ChaoGao,
                        ChaoKuan, XiangZhu, JiDan, ChaDianShijian, BaDianShijian, isYunDi, MaTouCode)
                        VALUES
                               (?,?,?,?,?,?,?,?,?,?,?,
                                ?,?,?,?,?,?,?,?,?,?,
                                ?,?,?,?,?,?,?,?,?,?,
                                ?,?,?,?,?,?,?)`

	bind := make([]interface{}, 0)

	bind = append(bind, item.JingChangDate)
	bind = append(bind, item.CarNumber)
	bind = append(bind, item.JingChangZhuangtai)
	bind = append(bind, item.TidanCode)
	bind = append(bind, item.XiangHao)
	bind = append(bind, item.ChuanMing)
	bind = append(bind, item.HangCi)
	bind = append(bind, item.XiangXin)
	bind = append(bind, item.ChiCun)
	bind = append(bind, item.KongZhong)

	bind = append(bind, item.QianFeng)
	bind = append(bind, item.JingChangYuanyin)
	bind = append(bind, item.JieFei)
	bind = append(bind, item.CheDui)
	bind = append(bind, item.CheDuiDianHua)
	bind = append(bind, item.JingChangCaoZuo)
	bind = append(bind, item.BeiZhu)
	bind = append(bind, item.LuoXiangTiao)
	bind = append(bind, item.LaiYuan)
	bind = append(bind, item.SiJi)

	bind = append(bind, item.ChuanGongsi)
	bind = append(bind, item.JingChangXiangKuang)
	bind = append(bind, item.JingChangHao)
	bind = append(bind, item.ChangWei)
	bind = append(bind, item.JingChangYanXiangRen)
	bind = append(bind, item.WeixianPingJiBie)
	bind = append(bind, item.UN)
	bind = append(bind, item.HuoMing)
	bind = append(bind, item.XieYiJiage)
	bind = append(bind, item.JinChuChangStatus)

	bind = append(bind, item.ChaoGao)
	bind = append(bind, item.ChaoKuan)
	bind = append(bind, item.XiangZhu)
	bind = append(bind, item.JiDan)
	bind = append(bind, item.ChaDianShijian)
	bind = append(bind, item.BaDianShijian)
	bind = append(bind, item.isYunDi)
	bind = append(bind, item.MaTouCode)

	db.Exec(sql, bind...);
}