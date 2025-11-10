package clientorder

import (
	structur "AutoShop/Structur"
	"fmt"
)

type Ordr (map[string]structur.Order)

func (a Ordr) Print() {
	if len(a) == 0 {
		fmt.Println("Заказов нет")
		return
	}

	fmt.Println("--- Ваш Заказ ---")
	total := 0.0
	for name, order := range a {
		fmt.Printf("Товар: %s\n", name)
		fmt.Printf("Категория: %s\n", order.Category)
		fmt.Printf("Цена: %.2f руб.\n", order.Price)
		fmt.Println("  ---")
		total += order.Price
	}
	fmt.Print("Итоговая сумма заказа: %.2f руб.\n", total)
}

func (o Ordr) PrintByCategory() {
	if len(o) == 0 {
		fmt.Println("Заказов нет")
		return
	}

	categories := make(map[string][]structur.Order)
	for _, order := range o {
		categories[order.Category] = append(categories[order.Category], order)
	}

	fmt.Println("\n--- Ваш заказ по категориям ---")
	total := 0.0
	for category, orders := range categories {
		fmt.Printf("\n%s:\n", category)
		categoryTotal := 0.0
		for _, order := range orders {
			fmt.Printf("  - %s: %.2f руб.\n", order.Name, order.Price)
			categoryTotal += order.Price
		}
		fmt.Printf("  Всего по категории: %.2f руб.\n", categoryTotal)
		total += categoryTotal
	}
	fmt.Printf("\nИтоговая сумма заказа: %.2f руб.\n", total)
}

func (o Ordr) GetTotal() float64 {
	total := 0.0
	for _, order := range o {
		total += order.Price
	}
	return total
}

func (o Ordr) GetCount() int {
	return len(o)
}

func (o Ordr) RemoveItem(name string) bool {
	if _, exists := o[name]; exists {
		delete(o, name)
		return true
	}
	return false
}

func (o Ordr) Clear() {
	for k := range o {
		delete(o, k)
	}
}
