package main

import (
	"fmt"
	"horario-simples/schedules"
	"time"
)

var p = fmt.Println

func main() {
	ticker := time.NewTicker(time.Hour * 2) // checa todos os dias
	defer ticker.Stop()

	for t := range ticker.C {
		if t.Hour() == 15 && t.Minute() == 05 {
			p(schedules.ScheduleForTheDay())
		}
	}
}
