package personaldata

import "fmt"

// Personal структура описывает данные пользователя
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print() выводит имя, вес, рост.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %f\nРост: %f", p.Name, p.Weight, p.Height)
}
