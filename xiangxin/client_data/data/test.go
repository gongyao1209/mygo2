package data

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mygo2/db"
	"mygo2/xiangxin/helper"
)

var my_maps map[string]int

func GetData() string {

	//return get_enter_name(1, "2")

	my_maps = init_month_client()

	clients := find_clients_by_month("2021-01-01", "2021-02-01")

	_, err := helper.WriteToFile("客户名称", clients)
	if err != nil {
		return err.Error()
	}
	return "处理成功"
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



func find_clients_by_month(begin, end string) []helper.Enter {

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

	enters := make([]helper.Enter, 0)

	rows, err := mydb.Query(sql, bind...)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		enter := helper.Enter{}
		_ = rows.Scan(&enter.Enter_key, &enter.Id, &enter.Enter_type)

		enter.Enter_name = get_enter_name(enter.Id, enter.Enter_type)

		if _, exists := my_maps[enter.Enter_key]; exists {
			enter.Is_next_month = 1
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
