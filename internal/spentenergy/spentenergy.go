package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("invalid parameters")
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	numberOfCalories := ((weight * averageSpeed * durationMinutes) / float64(minInH)) * walkingCaloriesCoefficient
	return numberOfCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("invalid parameters")
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	numberOfCalories := (weight * averageSpeed * durationMinutes) / float64(minInH)
	return numberOfCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	averageSpeed := dist / duration.Hours()
	return averageSpeed
}

func Distance(steps int, height float64) float64 {
	lengthStep := height * stepLengthCoefficient
	dist := (lengthStep * float64(steps)) / float64(mInKm)
	return dist
}
