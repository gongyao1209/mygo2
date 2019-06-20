package package1

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Mobile string `json:"mobile"`
}

//
func NewPerson(id, age int, name, mobile string) *Person {
	return &Person{
		Id:     id,
		Name:   name,
		Age:    age,
		Mobile: mobile,
	}
}

func JsonToPerson(str string) {
	p := Person{}
	json.Unmarshal([]byte(str), &p)
	fmt.Println("name: ", p.Name, ", mobile: ", p.Mobile)
}

func JsonToOther(str string) {
	var i interface{}
	json.Unmarshal([]byte(str), &i)

	c := i.(map[string]interface{})

	for k, v := range c {

		switch v.(type) {
		case float64:
			fmt.Printf("key := %s, value = %.0f\n", k, v.(float64))
		case string:
			fmt.Printf("key := %s, value = %s\n", k, v.(string))
		default:
			fmt.Println("error")
		}
	}
}
