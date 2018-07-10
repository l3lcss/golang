package main

import (
	"fmt"
)

func init() {
	// println("init func()")
}

func main() {
	// var name string
	// name := ""
	// fmt.Scanf("%s", &name)
	// fmt.Printf("Hello %s\n", name)

	// for index := 0; index < 10; index++ {
	// 	fmt.Println(index)
	// }

	// var i int64
	// var name string
	// var arr = []string{
	// 	"a",
	// 	"b",
	// }

	// var slice []string ต้อง make เท่านั้นถึงจะ append ได้
	// slice := []string{}
	// slice = append(slice, "aa")
	// fmt.Printf("%+v", slice)
	// fmt.Println()

	// mapVal := map[string]interface{}{
	// 	"key":    "value",
	// 	"number": 5,
	// 	"nestedMap": map[string]interface{}{
	// 		"childKey": "childVal",
	// 	},
	// }
	// fmt.Printf("%+v", mapVal)
	// Hello("l3lcS", 21)
	// s, s2 := Prinln(0)
	// fmt.Println(s, s2)

	sum, average := Homework(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("sum = ", sum)
	fmt.Println("average = ", average)
}

//Hello (name, surname string)
func Hello(name string, age int) {
	fmt.Println("Hello", name, age)
}

//Prinln func
func Prinln(num int, a ...string) (string, string) { //_ คือ ignore ตัวแปรแหละ แต่เอามารับไว้เฉยๆ ไม่ต้องการใช้ เพื่อให้ได้ครบตามที่ มัน return กลับมา
	for _, e := range a { //_ คือ index , e คือ element
		print(e + "")
	}
	return "", "val"
}

//ConditionFunc func
func ConditionFunc(i int) int {
	if i == 0 {
		fmt.Println(false)
		return 0
	}
	return 1
	// if b := i + 5; b < 5 { //ทำ 1 งานก่อน เช็ค , เช็คหลังเครื่องหมาย ;
	// 	fmt.Println(false)
	// }
}

//Homework func
func Homework(a ...int) (int, float32) {
	sum := 0
	count := 0
	for _, e := range a {
		if checkPrimeNumber(e) {
			sum += e
			count++
			// fmt.Println(e)
		}
	}
	// println(count)
	average := float32(sum) / float32(count)
	// println(sum)
	// println(average)
	return sum, average
}

// checkPrimeNumber func
func checkPrimeNumber(n int) bool {
	flag := false
	for index := 2; index <= n/2; index++ {
		if n%index == 0 {
			flag = true
			break
		}
	}
	return !flag
}
