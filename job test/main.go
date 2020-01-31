package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now(), "task started running...")
	time.Sleep(time.Second * 30)
	fmt.Println(time.Now(), "task completed.")
}
