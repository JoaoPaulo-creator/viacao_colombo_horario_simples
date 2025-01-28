package schedules

type Schedules struct {
	Day  string
	Time string
}

func Get() *Schedules {
	return &Schedules{
		"Domingo",
		"15:25",
	}
}
