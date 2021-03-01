package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	///*
	//	fmt.Println("Hello, World!")
	// */
	//fmt.Println("Hello, World!")
	//var x float64 = 20.2
	//y := 42
	//y = 222
	//z := 111
	//fmt.Println(float64(x))
	//fmt.Println(y)
	//fmt.Printf("x is of type %T\n", x)
	//fmt.Printf("y is of type %T\n", y)
	////for y > 100{
	////	if y % 2 == 0 {
	////		fmt.Println(y)
	////	}
	////	y--
	////}
	//fmt.Println(max(z, y, x))
	//
	//country := [3]string{"nigeria", "egypt", "sweden"}
	//fmt.Println("Before function call: ", country[0])
	//// [nigeria egypt germany]
	//fmt.Println("After function call: ", country[1])
	//// [nigeria egypt sweden]
	//
	////myProfile := profile{
	////	Age:  15,
	////	Name: "Adeshina",
	////	Salary: 300,
	////	TechInterest: false,
	////}
	//myProfile := profile{
	//	name: "hi",
	//	age : 12,
	//}
	//fmt.Println(getValue(&myProfile))
	//
	//numbers := []int {0, 0, 0, 0, 0}
	//numbers = make([]int, 5, 63)
	//printSlice(numbers)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
//
//type profile struct {
//	name string
//	age  int
//}
////func modifyStruct(p profile) profile {
////	p.Age = 85
////	p.Name = "Balqees"
////	p.Salary = 500.45
////	p.TechInterest = true
////
////	return p
////}
//
//func getValue(profile *profile) int {
//	fmt.Println(*profile)
//	return 1
//}
//
///* function returning the max between two numbers */
//func max(num1, num2 int, num3 float64) int {
//	/* local variable declaration */
//	var result int
//
//	if (num1 > num2) {
//		result = num1
//	} else {
//		result = num2
//	}
//	return result
//}
//
//func printSlice(x []int) {
//	fmt.Println("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
//	fmt.Println("", len(x), cap(x), x)
//}
