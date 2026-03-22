package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	name := flag.String("name", "Devops 學習", "輸入你的名字")
	flag.Parse()
	fmt.Printf("Hello %s!, 這是我的 DevOps 學習旅程~\n", *name)
	fmt.Printf("今天是 %s，我成功用 ENTRYPOINT 跑起來了", time.Now().Format("2006-01-02"))
}
