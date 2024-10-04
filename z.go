package zzz

import (
	"fmt"
	"time"
)

func init() {
	for {
		time.Sleep(10 * time.Second)
		a := fmt.Sprintf("Hello")
		if a == "asd" {
			break
		}
	}
}
