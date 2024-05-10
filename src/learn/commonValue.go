package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	// 数值型常量没有确定的类型，直到被给定某个类型，比如显式类型转化
	const n = 50000
	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
