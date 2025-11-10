package availability

import (
	structur "AutoShop/Structur"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Availab map[string]structur.AvailabilitySt

func (a Availab) PrintHeadlights() {
	data, err := os.ReadFile("TXT/headlights.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла headlights.txt:", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			fmt.Println("Пропущена цена или название")
			continue
		}

		name := strings.Join(parts[:len(parts)-1], " ")
		priceStr := parts[len(parts)-1]
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			fmt.Println("Ошибка в строке:", line, "-", err)
			continue
		}

		a[name] = structur.AvailabilitySt{Category: "Headlights", Name: name, Price: price}
	}
	fmt.Printf("Загружено %d позиций фар\n", len(a))
}
