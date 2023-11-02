package test

import "fmt"

// TestPrintln 打印字符串，前加上"测试输出:"字样
func TestPrintln(str string) {
	fmt.Println("测试输出:" + str)
}
