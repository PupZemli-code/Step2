package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// Ошибки для оборачивания
var (
	errLennSlice    = errors.New("len(slice) != 3")
	errValueMissing = errors.New("значение по адресу slice[x] отсутствует")
	errParsInt      = errors.New("ошибка приобразования string > int")
	errParsDuration = errors.New("ошибка приобразования string > time.Duranion")
	errDurationNil  = errors.New("значение Training.Duration <= 0")
	errTrainingType = errors.New("тип тренеровки != 'Бег' || slice[1] != 'Ходьба'")
)

// Training хранит данные о тренеровке
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse() парсит тренеровку
func (t *Training) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 3 {
		return fmt.Errorf("ошибка Parse: %w", errLennSlice)
	}
	for _, value := range slice {
		if value == "" {
			return fmt.Errorf("ошибка Parse: %w", errValueMissing)
		}
	}
	t.Steps, err = strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("ошибка Parse: %w", errParsInt)
	}
	if slice[1] != "Бег" || slice[1] != "Ходьба" {
		return fmt.Errorf("ошибка Parse: %w", errTrainingType)
	}
	t.TrainingType = slice[1]
	t.Duration, err = time.ParseDuration(slice[2])
	if err != nil {
		return fmt.Errorf("ошибка Parse: %w", errParsDuration)
	}
	return nil
}

// ActionInfo() формирует и возвращает строку с данными о тренировке
func (t *Training) ActionInfo() (string, error) {
	distans := spentenergy.Distance(t.Steps)
	if t.Duration <= 0 {
		return "", fmt.Errorf("ошибка ActionInfo: %w", errDurationNil)
	}
	speed := spentenergy.MeanSpeed(int(distans), t.Duration)
	calors := 0.0
	var err error
	if t.TrainingType == "Бег" {
		calors, err = spentenergy.RunningSpentCalories(t.Steps, personaldata.Personal.Weight, t.Duration)
	} else {
		calors, err = spentenergy.WalkingSpentCalories(t.Steps, personaldata.Personal(Weight), t.Duration)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Тип тренировки: %s\nДлительность: %0.2f ч.\nДистанция: %0.2f км.\nСкорость: %0.2f км/ч\nСожгли калорий: %0.2f\n", t.TrainingType, t.Duration, distans, speed, calors)
}
