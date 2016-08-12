package main 
var a = "hello world" 	//根据内容自行判断类型
var b string = "this is golang"
var c bool //没有初始化，有默认值
//d:="love12"  //编译出错，只可以用于函数体内

var x, y int = 1 ,2
var u,v=123,"hello"
var ( //这种写法一般只用于声明全局变量
	z int 
	w bool
 )

func main() {
	d := "love" //一般用于局部变量 声明后在相同的代码块内必须被使用 全局变量则可以不使用
	println(a,b,c,d)
	println(x,y,u,v,z,w)
	println(&a)
	

}