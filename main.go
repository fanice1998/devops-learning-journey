package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "Devops learner", "Enter your name")
	flag.Parse()
	fmt.Printf("Hello %s!, This my DevOps learning journey~\n", *name)
}
