// Package daysteps считает шаги и калории за прогулки в течение дня.
package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps хранит информацию о шагах и времени прогулки.
type DaySteps struct {
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

// Parse достает количество шагов и время из сырой текстовой строки.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверное количество параметров в строке")
	}
	if strings.HasPrefix(parts[0], " ") || strings.HasSuffix(parts[0], " ") {
		return errors.New("строка содержит лишние пробелы")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}
	ds.Steps = steps
	ds.Duration = duration
	return nil
}

// ActionInfo собирает красивый текст с итогами прогулки с расстоянием и калориями.
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 || ds.Duration <= 0 || ds.Weight <= 0 || ds.Height <= 0 {
		return "", errors.New("некорректные параметры активности")
	}
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, dist, cal), nil
}
