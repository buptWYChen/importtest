package importtest

import "fmt"

// TestPrintln2 打印字符串，前加上"测试输出:"字样
func TestPrintln2(str string) {
	fmt.Println("测试输出2:" + str)
}
