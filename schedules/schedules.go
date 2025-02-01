package schedules

import (
	"horario-simples/config"
	"time"
)

type DayOfTheWeek struct {
	Day string
}

func ScheduleForTheDay() string {
	return config.PerformRequest()

}

// Pega o dia da semana
func getDay() DayOfTheWeek {
	var d DayOfTheWeek

	today := time.Now().Weekday()
	switch today {
	case time.Sunday:
		d = dayOfTheWeek(config.SUNDAY)
	case time.Saturday:
		d = dayOfTheWeek(config.SATURDAY)
	default:
		d = dayOfTheWeek(config.UTIL)
	}

	return d
}

func dayOfTheWeek(day string) DayOfTheWeek {
	return DayOfTheWeek{Day: day}
}
