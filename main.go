package main

import (
	"fmt"
	"horario-simples/config"
)

func main() {
	content := "isso eh um teste"
	r := config.PerformRequest(content)
	fmt.Println(r)
}
