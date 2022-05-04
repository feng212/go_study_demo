package main

import (
	"fmt"
	"go/constant"
)

//var 变量值可改变，在go中每个变量声明必须被使用，否则编译不通过
//const 常量 值不可改变
type StructType struct {
	Name string
	Age  int
}

//1.2全局变量 小写开头不能被外部引用，大写开头可以，结构体定义，体内变量同样
var (
	name string
	age  int
	sex  byte
)

//2.2全局常量
const Title string = "data_type"
const p int = 20

func main() {
	//1.1局部变量，声明后默认初始化，基本数据类型有其对应的默认值
	var intType1 int
	var boolType1 bool
	var stringType1 string
	var floatType1 float64
	var mapsType1 map[string]string
	var structType1 StructType
	var sliceType1 []string
	var channelType1 chan int

	//1.3 语法糖，只对局部变量适用   data_name := value  go会根据其值类型自动适配类型
	//                    相当于   var data_name := value
	intType2 := 5
	boolType2 := true
	sliceType2 := []string{"a", "b", "c"}

	//2 常量 顾名思义必须声明时赋值
	//2.1局部常量
	const stringType2 = "220504"
	const highs int = 10

	//声明后必须被使用
	fmt.Printf("变量类型:    %T,%T,%T ,%T,%T,%T,%T,%T\n",
		intType1, boolType1, stringType1, floatType1, mapsType1, structType1, sliceType1, channelType1)
	fmt.Printf("变量默认值:  %v,%v,%v,%v,%v,%v,%v,%v\n",
		intType1, boolType1, stringType1, floatType1, mapsType1, structType1, sliceType1, channelType1)

	//全局变量默认初始化，声明后可不使用
	fmt.Println(name, age, sex)

	//语法糖模式
	fmt.Printf("变量类型:    %T,%T,%T\n",
		intType2, boolType2, sliceType2)
	fmt.Printf("变量默认值:  %v,%v,%v\n",
		intType2, boolType2, sliceType2)
}
