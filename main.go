package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"gg": 1,
	}
	fmt.Println(m["gg"])
	m["key"] = 5
	fmt.Println(m["key"])
	fmt.Println(m["f"]) // default ของ value คือ 0

	//if นี้สำหรับการ เช็คว่า key นั้นมีจริงๆหรือป่าวววว
	if _, ok := m["f"]; ok {
		fmt.Println("have")
	} else {
		fmt.Println("havn't")
	}
	//------------------------------------------------------------------------------------

}
