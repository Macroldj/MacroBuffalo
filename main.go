package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func main() {
	fmt.Println("starting......")
	orm.RunCommand()
	fmt.Println("ended......")
}
