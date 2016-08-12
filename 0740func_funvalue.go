package main

import (
   "fmt"
   "math"
)

func main(){
   /* 声明函数变量 */
	//可以 用一个变量来保存函数
   getSquareRoot := func(x float64) float64 {
      return math.Sqrt(x)
   }

   /* 使用函数 */
   fmt.Println(getSquareRoot(9))

}