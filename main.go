package main

import (
	"fmt"
	"os"

	"github.com/Crodu/goexpert_dollar/client"
	"github.com/Crodu/goexpert_dollar/server"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument.")
		return
	}

	arg := os.Args[1]

	switch arg {
	case "server":
		server.StartServer()
	case "client":
		client.GetExchange()
	default:
		fmt.Println("Invalid argument.")
	}
}
