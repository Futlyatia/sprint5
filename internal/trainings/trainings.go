package trainings

import (
	"errors"
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
	objects := strings.Split(datastring, ",")
	if len(objects) != 3 {
		return errors.New("Неверный формат строки")
	}
	steps, err := strconv.Atoi(objects[0])
	if err != nil {
		return err
	}
	t.Steps = steps
	t.TrainingType = objects[1]
	duration, err := time.ParseDuration(objects[2])
	if err != nil {
		return err
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	switch t.TrainingType {
	case "Бег":
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`Тип тренировки: %s\n
							Длительность: %.2f ч.\n
							Дистанция: %.2f км.\n
							Скорость: %.2f км/ч\n
							Сожгли калорий: %.2f\n`, t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
	case "Ходьба":
		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`Тип тренировки: %s\n
							Длительность: %.2f ч.\n
							Дистанция: %.2f км.\n
							Скорость: %.2f км/ч\n
							Сожгли калорий: %.2f\n`, t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
	default:
		return "", errors.New("Неизвестный тип тренировки")
	}
}
