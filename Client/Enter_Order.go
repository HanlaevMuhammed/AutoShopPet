package clientorder

import (
	structur "AutoShop/Structur"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Ordr map[string]structur.Order

func EnterOrder(orders Ordr, availabilitys map[string]structur.AvailabilitySt) {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Доступные категории:\n")
	categories := make(map[string]bool)
	for _, item := range availabilitys {
		categories[item.Category] = true
	}

	for category := range categories {
		fmt.Println("-", category)
	}

	for {
		fmt.Println("Введите категорию товара(exit для выхода): ")
		fmt.Print("> ")
		if !scanner.Scan() {
			fmt.Println("Ошибка чтения")
			return
		}

		category := strings.TrimSpace(scanner.Text())
		if strings.ToLower(category) == "exit" {
			break
		}
		fmt.Printf("Товары в категории '%s':\n", category)
		found := false
		for name, item := range availabilitys {
			if item.Category == category {
				fmt.Printf("- %s: %.2f руб.\n", name, item.Price)
				found = true
			}
		}

		if !found {
			fmt.Println("Категория не найдена или пуста")
			continue
		}
		for {
			fmt.Println("Введите название товара(exit для выхода):")
			fmt.Print("> ")
			if !scanner.Scan() {
				fmt.Println("Ошибка записи")
				return
			}

			name := strings.TrimSpace(scanner.Text())
			if strings.ToLower(name) == "exit" {
				break
			}

			if avail, ok := availabilitys[name]; ok {
				orders[name] = structur.Order{
					Category: avail.Category,
					Name:     avail.Name,
					Price:    avail.Price,
				}
				fmt.Printf("Добавлено в заказ: %s -> %s -> %.2f руб\n", avail.Category, avail.Name, avail.Price)
			} else {
				fmt.Println("Товар не найден в текущей категории")

			}

		}

	}

}
