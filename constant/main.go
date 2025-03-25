// 定数の計算
package main

import (
	"fmt"
	"go/constant"
	"go/token"
)

func main() {
	x := constant.MakeInt64(5)
	y := constant.MakeInt64(3)
	result := constant.BinaryOp(x, token.ADD, y)

	fmt.Println("計算結果:", result)
}

// 定数の型を判定する
// package main

// import (
// 	"fmt"
// 	"go/constant"
// )

// func main() {
// 	c := constant.MakeFloat64(3.14)
// 	fmt.Println("定数", c, "型:", c.Kind())
// }