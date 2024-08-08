package main

import (
	"fmt"
	"os"

	"test/utils"
)

func main() {
	conf, err := utils.LoadConfig()

	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(0)
	}

	port := fmt.Sprintf(":%s", conf.PORT)

	fmt.Printf(port)
}
