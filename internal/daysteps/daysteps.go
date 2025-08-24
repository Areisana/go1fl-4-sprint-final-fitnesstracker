package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
	// Ошибка
)

var (
	ErrorWrongFormat = errors.New("Неверный формат данных")
	ErrorWrongSteps  = errors.New("Неверное количество шагов")
	ErrorWrongTime   = errors.New("Неверный формат времени")
)

func parsePackage(data string) (int, time.Duration, error) {
	// преобразование входящей строки в слайс
	dataSlice := strings.Split(data, ",")
	//
	if len(dataSlice) != 2 {
		return 0, 0, ErrorWrongFormat
	}
	steps, err := strconv.Atoi(dataSlice[0])
	// Проверка "была ли ошибка"
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, ErrorWrongSteps
	}
	t, err := time.ParseDuration(dataSlice[1])

	if err != nil {
		return 0, 0, err
	}

	if t <= 0 {
		return 0, 0, ErrorWrongTime
	}

	return steps, t, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}
	distation := float64(steps) * stepLength / mInKm
	totalCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("Количество шагов: %b.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distation, totalCalories)

}
