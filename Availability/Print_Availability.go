package availability

import (
	"fmt"
)

func (a Availab) PrintAll() {
	fmt.Println("--- Доступные товары ---")
	for _, item := range a {
		fmt.Printf("Категория: %s, Название: %s, Цена: %.2f\n", item.Category, item.Name, item.Price)
	}
}
