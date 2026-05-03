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

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	objects := strings.Split(datastring, ",")
	if len(objects) != 2 {
		return errors.New("Неверный формат строки")
	}
	steps, err := strconv.Atoi(objects[0])
	if err != nil {
		return err
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(objects[1])
	if err != nil {
		return err
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(`Количество шагов: %d.
						Дистанция составила %.2f км.
						Вы сожгли %.2f ккал.`, ds.Steps, distance, calories), nil
}
