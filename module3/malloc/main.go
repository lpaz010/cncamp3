package main

import "C"
import (
	"fmt"
	"time"
)

/*
从孟老师那里的代码
*/
func main() {
	holder := []*C.char{}

	// only loop 10 times to avoid exhausting the host memory
	loop := 10
	for i := 0; i < loop; i++ {
		fmt.Printf("Allocating %dMb memory, raw memory is %d\n", i*100, i*100*1024*1025)
		holder.append(holder, (*C.char)(C.allocMemory()))
		time.Sleep(time.Minute)
	}

}
