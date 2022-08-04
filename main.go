package main

import (
	"fmt"
	"go-config/config"
)

func main() {
	fmt.Println(config.GetDatabasePassword())
}
