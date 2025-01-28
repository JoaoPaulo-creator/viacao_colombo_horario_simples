package schedules

import (
	"fmt"
	"strings"
	"time"
)

const (
	SATURDAY string = "SABADO"
	SUNDAY   string = "DOMINGO"

	UTIL string = "UTIL"

	// Linhas de onibus
	SAOSEBASTIAO = "São Sebastião"
	SAOGABRIEL   = "São Gabriel"
	OSASCO       = "Osasco"
)

var LineMap = map[string]string{
	"saoSebastiao": SAOSEBASTIAO,
	"saoGabriel":   SAOGABRIEL,
	"osasco":       OSASCO,
}

type Schedules struct {
	Day  string
	Time []string
}

type DayOfTheWeek struct {
	Day string
}

type Lines struct {
	Name     string
	Schedule []string
}

// TODO: colocar um nome melhor nessa função
func Get() map[string][]Schedules {
	schedules := map[string][]Schedules{
		"São Sebastião": []Schedules{
			{
				Day:  UTIL,
				Time: []string{"14:00", "14:45"},
			},
			{
				Day:  SATURDAY,
				Time: []string{"11:00", "12:30"},
			},
			{
				Day:  SUNDAY,
				Time: []string{"12:00", "16:00"},
			},
		},
		"São Gabriel": []Schedules{
			{
				Day:  UTIL,
				Time: []string{"14:00", "14:45"},
			},
			{
				Day:  SATURDAY,
				Time: []string{"14:00", "14:45"},
			},
			{
				Day:  SUNDAY,
				Time: []string{"14:00", "14:45"},
			},
		},
		"Osasco": []Schedules{
			{
				Day:  UTIL,
				Time: []string{"14:00", "14:45"},
			},
			{
				Day:  SATURDAY,
				Time: []string{"14:00", "14:45"},
			},
			{
				Day:  SUNDAY,
				Time: []string{"14:00", "14:45"},
			},
		},
	}

	return schedules
}

// FIX: arrumar isso quando estiver com saco
func ScheduleForTheDay(line string) Lines {
	// Seta os horários a serem enviados se baseando no dia da semana
	gd := getDay()
	schedules := Get()

	time := make([]string, 0)
	for _, day := range schedules[line] {
		if gd.Day == UTIL {
			time = append(time, day.Time...)
			fmt.Printf("Horários osasco: Dia %s | Horarios: %s\n", day.Day, strings.Join(time, " | "))
			clear(time)
		}
	}

	l := Lines{
		Name:     line,
		Schedule: time,
	}

	return l
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
