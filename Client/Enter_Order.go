package clientorder

import (
	structur "AutoShop/Structur"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ordr map[string]structur.Order

func EnterOrder(orders Ordr, availabilitys map[string]structur.AvailabilitySt) {
	scanner := bufio.NewScanner(os.Stdin)

	categories := make(map[string]bool)
	for _, item := range availabilitys {
		categories[item.Category] = true
	}

	for {
		fmt.Println("\nДоступные категории:")
		for category := range categories {
			fmt.Println("-", category)
		}
		fmt.Println("\nВведите категорию товара (order для отображения заказа, exit для выхода):")
		fmt.Print("> ")
		if !scanner.Scan() {
			fmt.Println("Ошибка чтения")
			return
		}

		category := strings.TrimSpace(scanner.Text())
		if strings.ToLower(category) == "exit" {
			break
		}

		if strings.ToLower(category) == "order" {
			fmt.Println("\nТекущий заказ:")
			orders.Print()
			continue
		}

		categoryExists := false
		for _, item := range availabilitys {
			if item.Category == category {
				categoryExists = true
				break
			}
		}

		if !categoryExists {
			fmt.Println("Категория не найдена")
			continue
		}

		fmt.Printf("Товары в категории '%s':\n", category)
		found := false
		for name, item := range availabilitys {
			if item.Category == category {
				status := "✓ В наличии"
				if item.Quantity == 0 {
					status = "✗ Нет в наличии"
				} else if item.Quantity < 5 {
					status = "! Осталось менее 5 штук"
				}
				fmt.Printf("- %s: %.2f руб. [%d шт.] %s\n",
					name, item.Price, item.Quantity, status)
				found = true
			}
		}

		if !found {
			fmt.Println("В категории нет товаров")
			continue
		}

		fmt.Println("Введите название товара:")
		fmt.Print("> ")
		if !scanner.Scan() {
			fmt.Println("Ошибка записи")
			return
		}

		name := strings.TrimSpace(scanner.Text())
		if strings.ToLower(name) == "exit" {
			break
		}

		if strings.ToLower(name) == "order" {
			fmt.Println("\nТекущий заказ:")
			orders.Print()
			continue
		}

		avail, exists := availabilitys[name]
		if !exists || avail.Category != category {
			fmt.Println("Товар не найден в указанной категории")
			continue
		}

		if avail.Quantity == 0 {
			fmt.Printf("Товар '%s' временно отсутствует на складе\n", name)
			continue
		}

		fmt.Printf("Введите количество (доступно: %d шт.):\n", avail.Quantity)
		fmt.Print("> ")
		if !scanner.Scan() {
			fmt.Println("Ошибка записи")
			return
		}

		quantityStr := strings.TrimSpace(scanner.Text())
		quantity, err := strconv.Atoi(quantityStr)
		if err != nil || quantity <= 0 {
			fmt.Println("Некорректное количество")
			continue
		}

		if quantity > avail.Quantity {
			fmt.Printf("Недостаточно товара на складе. Доступно только %d шт.\n", avail.Quantity)
			break
		}

		availabilitys[name] = structur.AvailabilitySt{
			Category: avail.Category,
			Name:     avail.Name,
			Price:    avail.Price,
			Quantity: avail.Quantity - quantity,
		}

		if existingOrder, exists := orders[name]; exists {
			existingOrder.Quantity += quantity
			orders[name] = existingOrder
		} else {
			orders[name] = structur.Order{
				Category: category,
				Name:     name,
				Price:    avail.Price,
				Quantity: quantity,
			}
		}

		fmt.Printf("Добавлено в заказ: %s -> %s -> %.2f руб. -> %d шт.\n",
			category, name, avail.Price, quantity)
		fmt.Printf("Остаток на складе: %d шт.\n", availabilitys[name].Quantity)

	}

	if len(orders) > 0 {
		fmt.Println("\n=== ИТОГОВЫЙ ЗАКАЗ ===")
		orders.Print()
	} else {
		fmt.Println("Заказ пуст")
	}
}
