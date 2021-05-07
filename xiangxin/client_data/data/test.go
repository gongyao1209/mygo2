package data

import (
	"bufio"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mygo2/db"
	"os"
)

var my_maps map[string]int

func GetData() string {

	//return get_enter_name(1, "2")

	my_maps = init_month_client()

	clients := find_clients_by_month("2021-01-01", "2021-02-01")

	fmt.Println("客户名称", " ", "二月份有没有下单")
	for _, client := range clients {
		fmt.Println(client.enter_name, " ", client.is_next_month)
	}

	//创建一个新文件，写入内容 5 句 “http://c.biancheng.net/golang/”
	filePath := "/Users/gongyao/workspace/goproject/src/mygo2/202101月新增下单客户.csv"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("error")
		}
	}(file)
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)

	fmt.Println("客户名称", " ", "二月份有没有下单")

	i, err := write.WriteString(fmt.Sprint("来源", "\t", "客户名称", "\t", "二月份有没有下单"))
	if err != nil {
		return ""
	}
	fmt.Println(i)

	write.WriteString("\n")
	for _, client := range clients {
		var en_type_name string
		if client.enter_type == 1 {
			en_type_name = "箱信中台客户"
		} else {
			en_type_name = "Erp"
		}
		write.WriteString(fmt.Sprint(en_type_name, "\t", client.enter_name, "\t", client.is_next_month))
		write.WriteString("\n")
	}

	//for i := 0; i < 5; i++ {
	//	write.WriteString("http://c.biancheng.net/golang/ \n")
	//}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()

	return ""
}

func get_enter_name(id int, id_type int) string {

	mydb := db.GetDB()

	var sql string
	if id_type == 1 {
		sql = `select e.name from enter e where e.id = ?`
	} else {
		sql = `select e.name from erp_partner e where e.id = ?`
	}

	bind := make([]interface{}, 0)
	bind = append(bind, id)

	var name string
	err := mydb.QueryRow(sql, bind...).Scan(&name)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return name
}

type Enter struct {
	id            int
	enter_key     string
	enter_type    int
	enter_name    string
	is_next_month int
}

func find_clients_by_month(begin, end string) []Enter {

	mydb := db.GetDB()

	sql := `
         select concat('1_', c.enter_id) as enter_key, c.enter_id as id, 1 as type
         from stat_zhongtai_client as c
         where c.min_create_order_time >= ? and c.min_create_order_time < ? and c.order_count > 1
         union all
         select concat('2_', p.parnter_id) as enter_key, p.parnter_id as id, 2 as type
         from stat_erp_parnter as p
         where p.min_create_order_time >= ? and p.min_create_order_time < ? and p.order_count > 1`

	bind := make([]interface{}, 0)
	bind = append(bind, begin)
	bind = append(bind, end)
	bind = append(bind, begin)
	bind = append(bind, end)

	enters := make([]Enter, 0)

	rows, err := mydb.Query(sql, bind...)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		enter := Enter{}
		_ = rows.Scan(&enter.enter_key, &enter.id, &enter.enter_type)

		enter.enter_name = get_enter_name(enter.id, enter.enter_type)

		if _, exists := my_maps[enter.enter_key]; exists {
			enter.is_next_month = 1
		}

		enters = append(enters, enter)
	}

	return enters
}

func init_month_client() map[string]int {
	mydb := db.GetDB()

	sql := `select s.huoyue_json from stat_month_client s where s.month_int = 202102`
	rows, err := mydb.Query(sql)
	if err != nil {
		panic(err.Error())
	}

	huoyue_enters := make(map[string]int)
	for rows.Next() {
		var json_str string
		_ = rows.Scan(&json_str)

		res := make([]string, 0)
		json.Unmarshal([]byte(json_str), &res)
		if len(res) > 0 {
			for _, v := range res {
				huoyue_enters[v] = 1
			}
		}

	}

	return huoyue_enters
}
