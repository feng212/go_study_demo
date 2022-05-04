package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	num1 := 100
	num2 := 200
	newInt := new(int)
	fmt.Printf("num1的地址%p,值%v\n", &num1, num1)
	fmt.Printf("num2的地址%p,值%v\n", &num2, num2)
	//newInt本身的值保存的是引用的地址
	fmt.Printf("原始newInt:地址%p,值%v,引用的地址%p,引用的值%v\n", &newInt, newInt, &(*newInt), *newInt)
	//修改newInt的值
	//1.修改其引用的值
	*newInt = num1
	fmt.Printf("改变引用值后newInt:地址%p,值%v,引用的地址%p,引用的值%v\n", &newInt, newInt, &(*newInt), *newInt)
	//2.修改其地址，改变其引用
	newInt = &num2
	fmt.Printf("修改地址后newInt:地址%p,值%v,引用的地址%p,引用的值%v\n", &newInt, newInt, &(*newInt), *newInt)

	slice1 := make([]string, 5)
	fmt.Printf("直接makeslices1地址%p,值%v,长度%d,容量%d\n", &slice1, slice1, len(slice1), cap(slice1))
	var slice2 []string
	fmt.Printf("make前slices2地址%p,值%v,长度%d,容量%d\n", &slice2, slice2, len(slice2), cap(slice2))
	slice2 = make([]string, 10, 15)
	fmt.Printf("make后slices2地址%p,值%v,长度%d,容量%d\n", &slice2, slice2, len(slice2), cap(slice2))

	student1 := make([]Student, 0)
	student2 := make([]*Student, 0)
	student1 = []Student{{"a", 5}}
	//指针切片
	student2 = []*Student{{"b", 10}}
	fmt.Println(student1)
	for i := 0; i < len(student2); i++ {
		fmt.Println(student2[i])
	}
	student1[0].Age = 110
	student2[0].Age = 220
	fmt.Println(student1)
	for i := 0; i < len(student2); i++ {
		fmt.Println(*student2[i])
	}

	//currentTime := time.Now()
	t1 := time.Now().Unix()
	t2 := time.Now().Second() //1651649451 51
	fmt.Println(t1, t2)
	a := "dfdfdf"
	fmt.Printf("%s\n", a)
	fmt.Println(t1 - 1651649451)

}
