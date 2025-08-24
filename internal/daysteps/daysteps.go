package daysteps

import (
	"strconv"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	steps, err := strconv.Atoi(data[0])
	// error if steps equals zero
	if steps == 0 {
		return 0, err
	}
	// error if steps not int

	// error if steps have spaces before, after
	t, err := time.ParseDuration(data[1])
	if t == 0 {
		return 0, err
	}
	return steps, t, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	data := parsePackage()
}
