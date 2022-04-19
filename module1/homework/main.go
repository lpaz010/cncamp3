package main

import (
	"fmt"
	"time"
)

/**
可以参考的工程：
https://github.com/wqhn2020/geekbang-02.git
*/
func main() {
	// 1.1
	hw1_1()
	// 1.2
	hw1_2()
}

/*
编写一个小程序：
给定一个字符串数组
[“I”,“am”,“stupid”,“and”,“weak”]
用 for 循环遍历该数组并修改为
[“I”,“am”,“smart”,“and”,“strong”]
*/
func hw1_1() {
	fmt.Println("题目1.1")
	source := [5]string{"I", "am", "stupid", "and", "weak"}
	fmt.Printf("before: %s\n", source)

	for i, s := range source {
		if "stupid" == s {
			source[i] = "smart"
		} else if "weak" == s {
			source[i] = "strong"
		}
	}

	fmt.Printf("after: %s\n", source)
}

/*
基于 Channel 编写一个简单的单线程生产者消费者模型：

队列：
队列长度 10，队列元素类型为 int
生产者：
每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
消费者：
每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
*/
func hw1_2() {
	fmt.Println("题目1.2")
	msgQueue := make(chan int, 10)

	go produce(msgQueue)

	consume(msgQueue)
}

func produce(queue chan<- int) {
	for i := 0; i < 10; i++ {
		queue <- i
		time.Sleep(time.Second)
	}
	close(queue)
}

func consume(queue <-chan int) {
	for i := range queue {
		fmt.Printf("Get data: %d\n", i)
	}
}
