package main

import (
	"fmt"
	"printer/myModule"
)

func main() {
	fmt.Println("Module main!")
	fmt.Println(myModule.Version("Alexey"))
}
