package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// DaySteps содержит все необходимые данные о дневных прогулках
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

var (
	errLennSlice    = errors.New("len(slice) != 2")
	errValueMissing = errors.New("значение по адресу slice[x] отсутствует")
	errParsInt      = errors.New("ошибка приобразования string > int")
	errParsDuration = errors.New("ошибка приобразования string > time.Duranion")
	errDurationNil  = errors.New("значение Training.Duration <= 0")
)

// Parse() парсит строку с данными формата "678,0h50m" и записывает данные в соответствующие поля структуры DaySteps.
func (ds *DaySteps) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return fmt.Errorf("ошибка Parse в пакете daysteps: %w", errLennSlice)
	}
	for _, value := range slice {
		if value == "" {
			return fmt.Errorf("ошибка Parse в пакете daysteps: %w", errValueMissing)
		}
	}
	ds.Steps, err = strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("ошибка Parse в пакете daysteps: %w", errParsInt)
	}
	ds.Duration, err = time.ParseDuration(slice[1])
	if err != nil {
		return fmt.Errorf("ошибка Parse в пакете daysteps: %w", errParsDuration)
	}
	return nil
}

// ActionInfo() формирует и возвращает строку с данными о прогулке.
//
// Количество шагов: 792.
//
// Дистанция составила 0.51 км.
//
// Вы сожгли 221.33 ккал.
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", fmt.Errorf("ошибка Parse в пакете daysteps: %w", errDurationNil)
	}
	distance := float64(ds.Steps) * StepLength
	calors, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	outout := fmt.Sprintf(`
		Количество шагов: %d
		Дистанция составила %0.2f км.
		Вы сожгли %0.2f ккал
	`, ds.Steps, distance, calors)
	return outout, nil
}
