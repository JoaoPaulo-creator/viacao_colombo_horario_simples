package main

import (
	"fmt"
	"horario-simples/schedules"
	"log"
	"time"
)

var p = fmt.Println

func main() {
	ticker := time.NewTicker(time.Hour * 2) // checa todos os dias
	now := time.Now().Local()
	defer ticker.Stop()

	for t := range ticker.C {
		log.Printf("Horário de checagem: %v\n", now)
		if t.Hour() == 15 && t.Minute() == 05 {
			p(schedules.ScheduleForTheDay())
		}
	}
}
