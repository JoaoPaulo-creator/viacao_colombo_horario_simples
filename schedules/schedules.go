package schedules

import "time"

const (
	SATURDAY string = "SABADO"
	SUNDAY   string = "DOMINGO"
	UTIL     string = "UTIL" // business day
)

type DayOfTheWeek struct {
	Day string
}

// TODO: avaliar necessidade de manter esta função
func ScheduleForTheDay(lineName string) {

}

// Pega o dia da semana
func getDay() DayOfTheWeek {
	var d DayOfTheWeek

	today := time.Now().Weekday()
	switch today {
	case time.Sunday:
		d = dayOfTheWeek(SUNDAY)
	case time.Saturday:
		d = dayOfTheWeek(SATURDAY)
	default:
		d = dayOfTheWeek(UTIL)
	}

	return d
}

func dayOfTheWeek(day string) DayOfTheWeek {
	return DayOfTheWeek{Day: day}
}
