/**
* @Author: zhangyujin
* @Description:
* @Flie:demoCopyChannelVal
* @Date: 2023/12/10 14:39
 */
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	ch := make(chan Person, 1)
	person := Person{
		Age:  24,
		Name: "zhangyujin",
	}
	ch <- person
	person.Age = 25
	receivedPerson := <-ch
	//Age依然是 24 说明通道传递的是副本 不是指向原始值的指针
	fmt.Println(receivedPerson.Name, receivedPerson.Age)

}
