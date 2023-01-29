package main

import (
	"fmt"
	"github.com/cmdargs/request"
)

func main() {
	r := request.ParseCmd()
	fmt.Println(r.Cookies)
}