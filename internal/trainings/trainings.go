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
	errTrainingType = errors.New("неизвестный тип тренеровки slice[1] != 'Ходьба'|| slice[1] != 'Бег' ")
	//errRunningTrainingType = errors.New("неизвестный тип тренеровки slice[1] != 'Бег'")
	//errWalkingTrainingType = errors.New("неизвестный тип тренеровки slice[1] != 'Ходьба'")
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
		return fmt.Errorf("ошибка Parse в пакете trainings: %w", errLennSlice)
	}
	for _, value := range slice {
		if value == "" {
			return fmt.Errorf("ошибка Parse в пакете trainings: %w", errValueMissing)
		}
	}
	t.Steps, err = strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("ошибка Parse в пакете trainings: %w", errParsInt)
	}
	if slice[1] == "Бег" {
		t.TrainingType = slice[1]
	} else if slice[1] == "Ходьба" {
		t.TrainingType = slice[1]
	} else {
		return fmt.Errorf("ошибка Parse в пакете trainings: %w", errTrainingType)
	}

	t.Duration, err = time.ParseDuration(slice[2])
	if err != nil {
		return fmt.Errorf("ошибка Parse в пакете trainings: %w", errParsDuration)
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
		calors, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
	} else {
		calors, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	}
	if err != nil {
		fmt.Println(err)
	}
	output := fmt.Sprintf(`
	Тип тренировки: %s
	Длительность: %0.2f ч.
	Дистанция: %0.2f км.
	Скорость: %0.2f км/ч
	Сожгли калорий: %0.2f
	`, t.TrainingType, t.Duration.Hours(), distans, speed, calors)
	return output, nil
}
