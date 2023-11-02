package importtest

import (
	"fmt"
)

func main() {
	fmt.Println("Start.")
	TestPrintln1("Running...")
	TestPrintln2("Running...")
}

// TestPrintln1 打印字符串，前加上"测试输出:"字样
func TestPrintln1(str string) {
	fmt.Println("测试输出1:" + str)
}
