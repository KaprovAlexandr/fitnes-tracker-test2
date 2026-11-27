// Package spentenergy считает дистанцию, скорость и сожженные калории.
package spentenergy

import (
	"errors"
	"time"
)

const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

// WalkingSpentCalories считает сколько калорий ушло на ходьбу, учитывая шаги, вес, рост и время.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0.0, errors.New("некорректные параметры")
	}

	speed := MeanSpeed(steps, height, duration)
	baseCalories := (weight * speed * duration.Minutes()) / float64(minInH)

	return baseCalories * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories считает сколько калорий ушло при беге, учитывая шаги, вес, рост и время.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0.0, errors.New("некорректные параметры")
	}

	speed := MeanSpeed(steps, height, duration)

	return (weight * speed * duration.Minutes()) / float64(minInH), nil
}

// MeanSpeed рассчитывает среднюю скорость движения в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps < 0 {
		return 0.0
	}
	return Distance(steps, height) / duration.Hours()
}

// Distance рассчитывает пройденную дистанцию в километрах на основе количества шагов и роста пользователя.
func Distance(steps int, height float64) float64 {
	if steps < 0 {
		return 0.0
	}

	stepLength := height * stepLengthCoefficient
	totalMeters := float64(steps) * stepLength

	return totalMeters / float64(mInKm)
}
