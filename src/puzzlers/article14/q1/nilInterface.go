/**
* @Author: zhangyujin
* @Description:
* @Flie:nilInterface
* @Date: 2023/12/10 17:32
 */
package main

import "fmt"

type MyInterface interface {
	Method()
}
type MyType struct {
}

func (m *MyType) Method() {
	fmt.Println("method called")
}

func main() {
	var myInterfaceVar MyInterface
	// 接口变量没有分配具体值，此时它的值是真正的 nil
	if myInterfaceVar == nil {
		fmt.Println("myInterfaceVar is nil")
	} else {
		fmt.Println("myInterfaceVar is not nil")
	}
	// 分配具体值给接口变量
	myTypeVar := &MyType{}
	myInterfaceVar = myTypeVar

	// 接口变量有分配具体值，此时它的值不是真正的 nil
	if myInterfaceVar == nil {
		fmt.Println("myInterfaceVar is nil")
	} else {
		fmt.Println("myInterfaceVar is not nil")
	}
	//设置接口变量为 nil
	myInterfaceVar = nil
	if myInterfaceVar == nil {
		fmt.Println("myInterfaceVar is nil")
	} else {
		fmt.Println("myInterfaceVar is not nil")
	}
}
