package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию

	if len(datastring) == 0 {
		return fmt.Errorf("неверные параметры")
	}
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("неверные параметры")
	}
	ds.Steps, err = strconv.Atoi(parts[0])
	if ds.Steps <= 0 {
		return fmt.Errorf("0 шагов")
	}
	if err != nil {
		return fmt.Errorf("неверные параметры")
	}
	ds.Duration, err = time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("неверные параметры")
	}
	if ds.Duration.Minutes() <= 0 {
		return fmt.Errorf("0 минут")
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	kkal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	str := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, kkal)
	return str, nil
}
