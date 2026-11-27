// Package personaldata хранит информацию о пользователе: его имя, вес и рост.
package personaldata

import "fmt"

// Personal собирает в одну структуру профиль пользователя трекера.
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print красиво выводит имя, вес и рост человека на экран.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
