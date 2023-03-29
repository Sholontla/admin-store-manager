package main

import (
	"fmt"

	"service/stores/case1/config"
)

func main() {
	fmt.Println(config.DBConfig())
	fmt.Println(config.GRPCConfig1())
	fmt.Println(config.GRPCConfig2())
	fmt.Println(config.GRPCConfig3())
	fmt.Println(config.GRPCConfig4())
	fmt.Println(config.GRPCConfig5())

}
