package availability

import (
	structur "AutoShop/Structur"
	"fmt"
)

type Availab map[string]structur.AvailabilitySt

func (a Availab) PrintAll() {
	fmt.Println("--- Доступные товары ---")
	for _, item := range a {
		fmt.Printf("Категория: %s, Название: %s, Цена: %.2f\n", item.Category, item.Name, item.Price)
	}
}
