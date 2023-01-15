package main

import (
	"fmt"
	"go-config/config"
)

func main() {
	client, err := config.NewClient("config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(client.GetDatabaseUsername())
	fmt.Println(client.GetDatabasePassword())
}
