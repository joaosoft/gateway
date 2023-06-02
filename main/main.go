package main

import (
	"github.com/joaosoft/gateway"
)

func main() {
	m, err := gateway.NewGateway()
	if err != nil {
		panic(err)
	}

	if err := m.Start(); err != nil {
		panic(err)
	}
}
