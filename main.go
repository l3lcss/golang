package main

import (
	"errors"
	"fmt"
)

type stack []int

func main() {
	var value stack
	value = append(value, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	value = value.push(20)
	fmt.Println(value)
	fmt.Println(value.pop())
	// fmt.Println(value.pop())
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

	// sum, average := Homework(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// fmt.Println("sum = ", sum)
	// fmt.Println("average = ", average)

	//-------------------------day 2------------------------------
	// var slice = []string{"a", "b", "c", "d"}
	// // slice = append(slice, "a")
	// fmt.Println(slice[:2]) // [a b]
	// fmt.Println(slice[2:]) // [c d]

	// slice2 := slice

	// slice3 := append([]string(nil), slice...) // copy

	// index := 1
	// slice = append(slice[:index], slice[index+1:]...)
	// fmt.Println(slice, slice2, slice3)

	// s := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(reverseArray(s))
}

func (s stack) push(a int) stack {
	s = append(s, a)
	return s
}

func (s stack) pop() (int, error) {
	if len(s) == 0 {
		return 0, errors.New("Empty stack")
	}
	a := s[len(s)-1]
	s = s[:len(s)-1]
	return a, nil
}

func reverseArray(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
	}
	return nums
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
