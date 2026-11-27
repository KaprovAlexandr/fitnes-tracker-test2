// Package actioninfo собирает и выводит общую информацию обо всех тренировках и прогулках.
package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает методы для парсинга данных активности и вывода информации о ней.
type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

// Info принимает список сырых строк, обрабатывает их и выводит готовые результаты на экран.
func Info(dataset []string, dp DataParser) {
	for _, datastring := range dataset {
		err := dp.Parse(datastring)
		if err != nil {
			log.Printf("ошибка парсинга данных: %v", err)
			continue
		}

		infoString, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка при получении информации об активности: %v", err)
			continue
		}

		fmt.Print(infoString)
	}
}
