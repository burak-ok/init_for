package zzz

import (
	"fmt"
	"os"
	"time"
)

func init() {
	fmt.Printf("pid: %d\n", os.Getpid())
	fmt.Printf("ppid: %d\n", os.Getppid())

	time.Sleep(20 * time.Second)

	fmt.Printf("sleep done\n")
	// for {
	// 	time.Sleep(10 * time.Second)
	// 	a := fmt.Sprintf("Hello")
	// 	if a == "asd" {
	// 		break
	// 	}
	// }
}
