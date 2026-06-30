package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	//"3456,Ходьба,3h00m"
	if len(datastring) == 0 {
		return fmt.Errorf("invalid parameters")
	}
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("invalid parameters")
	}
	t.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid parameters")
	}
	if t.Steps < 1 {
		return fmt.Errorf("invalid parameters")
	}
	t.TrainingType = parts[1]
	t.Duration, err = time.ParseDuration(parts[2])
	if t.Duration <= 0 || t.Duration > time.Hour*24 {
		return fmt.Errorf("invalid parameters")
	}
	if err != nil {
		return fmt.Errorf("invalid parameters")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	steps := t.Steps
	dist := spentenergy.Distance(t.Steps, t.Height)
	avarageSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var kkal float64
	var err error
	switch t.TrainingType {
	case "Ходьба":
		kkal, err = spentenergy.WalkingSpentCalories(steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		kkal, err = spentenergy.RunningSpentCalories(steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}
	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %w", err)
	}
	str := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, avarageSpeed, kkal)
	return str, nil
}
