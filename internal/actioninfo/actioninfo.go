// actioninfo реализует вывод общей информации обо всех тренировках и прогулках
package actioninfo

import "fmt"

// интерфейс DataParser
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, str := range dataset {
		err := dp.Parse(str)
		if err != nil {
			fmt.Printf("ошибка парсинга строк: %v\n", err)
			continue
		}
		output, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("ошибка ActionInfo: %v\n", err)
		}
		fmt.Println(output)
	}
}
