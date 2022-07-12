package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("打印 :" + os.Args[0])
	s := []string{"pk", "jj", "dd"}
	fmt.Println("打印 :" + strings.Join(s, ","))
}
