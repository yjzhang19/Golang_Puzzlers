/**
* @Author: zhangyujin
* @Description:
* @Flie:demoSelect
* @Date: 2023/12/10 16:24
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		close(ch1)
	}()

	go func() {
		time.Sleep(3 * time.Second)
		close(ch2)
	}()
	for {
		select {
		case value, ok := <-ch1:
			if !ok {
				fmt.Println("ch1 closed")
				return
			}
			fmt.Println("Received from ch1:", value)

		case value, ok := <-ch2:
			if !ok {
				fmt.Println("ch2 closed")
				return
			}
			fmt.Println("Received from ch2:", value)
		default:
			fmt.Println("No channel Ready!")
			time.Sleep(time.Millisecond * 5)
		}
	}
}
