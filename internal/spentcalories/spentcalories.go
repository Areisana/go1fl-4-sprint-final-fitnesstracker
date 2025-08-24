package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

var (
	ErrorWrongFormat   = errors.New("неверный формат данных")
	ErrorWrongSteps    = errors.New("неверное количество шагов")
	ErrorWrongTime     = errors.New("неверный формат времени")
	ErrorWrongWeight   = errors.New("неверный вес")
	ErrorWrongHeight   = errors.New("неверный рост")
	ErrorWrongActivity = errors.New("неизвестный тип тренировки")
)

func parseTraining(data string) (int, string, time.Duration, error) {
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) != 3 {
		return 0, "", 0, ErrorWrongFormat
	}
	steps, err := strconv.Atoi(dataSlice[0])

	if err != nil {
		return 0, "", 0, err
	}

	if steps <= 0 {
		return 0, "", 0, ErrorWrongSteps
	}
	t, err := time.ParseDuration(dataSlice[2])

	if err != nil {
		return 0, "", 0, err
	}

	if t <= 0 {
		return 0, "", 0, ErrorWrongTime
	}
	activity := dataSlice[1]
	return steps, activity, t, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distance := float64(steps) * stepLength / mInKm
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distance := distance(steps, height)
	meanSpeed := distance / duration.Hours()
	return meanSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}
	if duration <= 0 {
		log.Println(ErrorWrongTime)
		return "", ErrorWrongTime
	}
	if steps <= 0 {
		log.Println(ErrorWrongTime)
		return "", ErrorWrongTime
	}
	switch activity {
	case "Ходьба":
		walkingSpentCalories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		distance := distance(steps, height)
		meanSpeed := meanSpeed(steps, height, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activity, duration.Hours(), distance, meanSpeed, walkingSpentCalories), nil
	case "Бег":
		runningSpentCalories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		distance := distance(steps, height)
		meanSpeed := meanSpeed(steps, height, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activity, duration.Hours(), distance, meanSpeed, runningSpentCalories), nil
	default:
		return "", ErrorWrongActivity
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, ErrorWrongSteps
	}
	if weight <= 0 {
		return 0, ErrorWrongWeight
	}
	if duration <= 0 {
		return 0, ErrorWrongTime
	}
	meanSpeed := meanSpeed(steps, height, duration)
	runningSpentCalories := (weight * meanSpeed * duration.Minutes()) / minInH
	return runningSpentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, ErrorWrongSteps
	}
	if height <= 0 {
		return 0, ErrorWrongHeight
	}
	if weight <= 0 {
		return 0, ErrorWrongWeight
	}
	if duration <= 0 {
		return 0, ErrorWrongTime
	}
	meanSpeed := meanSpeed(steps, height, duration)
	runningSpentCalories := (weight * meanSpeed * duration.Minutes()) / minInH
	walkingSpentCalories := runningSpentCalories * walkingCaloriesCoefficient
	return walkingSpentCalories, nil
}
