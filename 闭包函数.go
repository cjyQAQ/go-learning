package main

import "fmt"
/*

尽管 x 是 Adder 函数中的局部变量，但由于闭包函数引用了它，x 的生命周期被延长到与闭包函数相同。
这就是闭包的强大之处：闭包捕获并保存了它引用的外部变量的环境，因此这些变量在闭包函数的生命周期内始终有效。


*/


func main(){
	f := A()
	fmt.Println(f(1))       //1
	fmt.Println(f(20))      //21
	fmt.Println(f(300))    //321

}


func A() func(int) int{
	var x int
	return func(detal int) int {
		x += detal
		return x 
	}
}